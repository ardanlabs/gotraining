// Copyright ©2013 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat64

import (
	"math"

	"github.com/gonum/blas"
	"github.com/gonum/blas/blas64"
	"github.com/gonum/lapack/lapack64"
	"github.com/gonum/matrix"
)

// QR is a type for creating and using the QR factorization of a matrix.
type QR struct {
	qr   *Dense
	tau  []float64
	cond float64
}

func (qr *QR) updateCond() {
	// A = QR, where Q is orthonormal. Orthonormal multiplications do not change
	// the condition number. Thus, ||A|| = ||Q|| ||R|| = ||R||.
	n := qr.qr.mat.Cols
	work := make([]float64, 3*n)
	iwork := make([]int, n)
	r := qr.qr.asTriDense(n, blas.NonUnit, blas.Upper)
	v := lapack64.Trcon(matrix.CondNorm, r.mat, work, iwork)
	qr.cond = 1 / v
}

// Factorize computes the QR factorization of an m×n matrix a where m >= n. The QR
// factorization always exists even if A is singular.
//
// The QR decomposition is a factorization of the matrix A such that A = Q * R.
// The matrix Q is an orthonormal m×m matrix, and R is an m×n upper triangular matrix.
// Q and R can be extracted from the QFromQR and RFromQR methods on Dense.
func (qr *QR) Factorize(a Matrix) {
	m, n := a.Dims()
	if m < n {
		panic(matrix.ErrShape)
	}
	k := min(m, n)
	if qr.qr == nil {
		qr.qr = &Dense{}
	}
	qr.qr.Clone(a)
	work := make([]float64, 1)
	qr.tau = make([]float64, k)
	lapack64.Geqrf(qr.qr.mat, qr.tau, work, -1)

	work = make([]float64, int(work[0]))
	lapack64.Geqrf(qr.qr.mat, qr.tau, work, len(work))
	qr.updateCond()
}

// TODO(btracey): Add in the "Reduced" forms for extracting the n×n orthogonal
// and upper triangular matrices.

// RFromQR extracts the m×n upper trapezoidal matrix from a QR decomposition.
func (m *Dense) RFromQR(qr *QR) {
	r, c := qr.qr.Dims()
	m.reuseAs(r, c)

	// Disguise the QR as an upper triangular
	t := &TriDense{
		mat: blas64.Triangular{
			N:      c,
			Stride: qr.qr.mat.Stride,
			Data:   qr.qr.mat.Data,
			Uplo:   blas.Upper,
			Diag:   blas.NonUnit,
		},
		cap: qr.qr.capCols,
	}
	m.Copy(t)

	// Zero below the triangular.
	for i := r; i < c; i++ {
		zero(m.mat.Data[i*m.mat.Stride : i*m.mat.Stride+c])
	}
}

// QFromQR extracts the m×m orthonormal matrix Q from a QR decomposition.
func (m *Dense) QFromQR(qr *QR) {
	r, _ := qr.qr.Dims()
	m.reuseAsZeroed(r, r)

	// Set Q = I.
	for i := 0; i < r*r; i += r + 1 {
		m.mat.Data[i] = 1
	}

	// Construct Q from the elementary reflectors.
	work := make([]float64, 1)
	lapack64.Ormqr(blas.Left, blas.NoTrans, qr.qr.mat, qr.tau, m.mat, work, -1)
	work = make([]float64, int(work[0]))
	lapack64.Ormqr(blas.Left, blas.NoTrans, qr.qr.mat, qr.tau, m.mat, work, len(work))
}

// SolveQR finds a minimum-norm solution to a system of linear equations defined
// by the matrices A and b, where A is an m×n matrix represented in its QR factorized
// form. If A is singular or near-singular a Condition error is returned. Please
// see the documentation for Condition for more information.
//
// The minimization problem solved depends on the input parameters.
//  If trans == false, find X such that ||A*X - b||_2 is minimized.
//  If trans == true, find the minimum norm solution of A^T * X = b.
// The solution matrix, X, is stored in place into the receiver.
func (m *Dense) SolveQR(qr *QR, trans bool, b Matrix) error {
	r, c := qr.qr.Dims()
	br, bc := b.Dims()

	// The QR solve algorithm stores the result in-place into the right hand side.
	// The storage for the answer must be large enough to hold both b and x.
	// However, this method's receiver must be the size of x. Copy b, and then
	// copy the result into m at the end.
	if trans {
		if c != br {
			panic(matrix.ErrShape)
		}
		m.reuseAs(r, bc)
	} else {
		if r != br {
			panic(matrix.ErrShape)
		}
		m.reuseAs(c, bc)
	}
	// Do not need to worry about overlap between m and b because x has its own
	// independent storage.
	x := getWorkspace(max(r, c), bc, false)
	x.Copy(b)
	t := qr.qr.asTriDense(qr.qr.mat.Cols, blas.NonUnit, blas.Upper).mat
	if trans {
		ok := lapack64.Trtrs(blas.Trans, t, x.mat)
		if !ok {
			return matrix.Condition(math.Inf(1))
		}
		for i := c; i < r; i++ {
			zero(x.mat.Data[i*x.mat.Stride : i*x.mat.Stride+bc])
		}
		work := make([]float64, 1)
		lapack64.Ormqr(blas.Left, blas.NoTrans, qr.qr.mat, qr.tau, x.mat, work, -1)
		work = make([]float64, int(work[0]))
		lapack64.Ormqr(blas.Left, blas.NoTrans, qr.qr.mat, qr.tau, x.mat, work, len(work))
	} else {
		work := make([]float64, 1)
		lapack64.Ormqr(blas.Left, blas.Trans, qr.qr.mat, qr.tau, x.mat, work, -1)
		work = make([]float64, int(work[0]))
		lapack64.Ormqr(blas.Left, blas.Trans, qr.qr.mat, qr.tau, x.mat, work, len(work))

		ok := lapack64.Trtrs(blas.NoTrans, t, x.mat)
		if !ok {
			return matrix.Condition(math.Inf(1))
		}
	}
	// M was set above to be the correct size for the result.
	m.Copy(x)
	putWorkspace(x)
	if qr.cond > matrix.ConditionTolerance {
		return matrix.Condition(qr.cond)
	}
	return nil
}

// SolveQRVec finds a minimum-norm solution to a system of linear equations.
// Please see Dense.SolveQR for the full documentation.
func (v *Vector) SolveQRVec(qr *QR, trans bool, b *Vector) error {
	if v != b {
		v.checkOverlap(b.mat)
	}
	r, c := qr.qr.Dims()
	// The Solve implementation is non-trivial, so rather than duplicate the code,
	// instead recast the Vectors as Dense and call the matrix code.
	if trans {
		v.reuseAs(r)
	} else {
		v.reuseAs(c)
	}
	return v.asDense().SolveQR(qr, trans, b.asDense())
}

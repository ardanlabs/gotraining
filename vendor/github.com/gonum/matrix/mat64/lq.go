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

// LQ is a type for creating and using the LQ factorization of a matrix.
type LQ struct {
	lq   *Dense
	tau  []float64
	cond float64
}

func (lq *LQ) updateCond() {
	// A = LQ, where Q is orthonormal. Orthonormal multiplications do not change
	// the condition number. Thus, ||A|| = ||L|| ||Q|| = ||Q||.
	m := lq.lq.mat.Rows
	work := make([]float64, 3*m)
	iwork := make([]int, m)
	l := lq.lq.asTriDense(m, blas.NonUnit, blas.Lower)
	v := lapack64.Trcon(matrix.CondNorm, l.mat, work, iwork)
	lq.cond = 1 / v
}

// Factorize computes the LQ factorization of an m×n matrix a where n <= m. The LQ
// factorization always exists even if A is singular.
//
// The LQ decomposition is a factorization of the matrix A such that A = L * Q.
// The matrix Q is an orthonormal n×n matrix, and L is an m×n upper triangular matrix.
// L and Q can be extracted from the LFromLQ and QFromLQ methods on Dense.
func (lq *LQ) Factorize(a Matrix) {
	m, n := a.Dims()
	if m > n {
		panic(matrix.ErrShape)
	}
	k := min(m, n)
	if lq.lq == nil {
		lq.lq = &Dense{}
	}
	lq.lq.Clone(a)
	work := make([]float64, 1)
	lq.tau = make([]float64, k)
	lapack64.Gelqf(lq.lq.mat, lq.tau, work, -1)
	work = make([]float64, int(work[0]))
	lapack64.Gelqf(lq.lq.mat, lq.tau, work, len(work))
	lq.updateCond()
}

// TODO(btracey): Add in the "Reduced" forms for extracting the m×m orthogonal
// and upper triangular matrices.

// LFromLQ extracts the m×n lower trapezoidal matrix from a LQ decomposition.
func (m *Dense) LFromLQ(lq *LQ) {
	r, c := lq.lq.Dims()
	m.reuseAs(r, c)

	// Disguise the LQ as a lower triangular
	t := &TriDense{
		mat: blas64.Triangular{
			N:      r,
			Stride: lq.lq.mat.Stride,
			Data:   lq.lq.mat.Data,
			Uplo:   blas.Lower,
			Diag:   blas.NonUnit,
		},
		cap: lq.lq.capCols,
	}
	m.Copy(t)

	if r == c {
		return
	}
	// Zero right of the triangular.
	for i := 0; i < r; i++ {
		zero(m.mat.Data[i*m.mat.Stride+r : i*m.mat.Stride+c])
	}
}

// QFromLQ extracts the n×n orthonormal matrix Q from an LQ decomposition.
func (m *Dense) QFromLQ(lq *LQ) {
	r, c := lq.lq.Dims()
	m.reuseAs(c, c)

	// Set Q = I.
	for i := 0; i < c; i++ {
		v := m.mat.Data[i*m.mat.Stride : i*m.mat.Stride+c]
		zero(v)
		v[i] = 1
	}

	// Construct Q from the elementary reflectors.
	h := blas64.General{
		Rows:   c,
		Cols:   c,
		Stride: c,
		Data:   make([]float64, c*c),
	}
	qCopy := getWorkspace(c, c, false)
	v := blas64.Vector{
		Inc:  1,
		Data: make([]float64, c),
	}
	for i := 0; i < r; i++ {
		// Set h = I.
		zero(h.Data)
		for j := 0; j < len(h.Data); j += c + 1 {
			h.Data[j] = 1
		}

		// Set the vector data as the elementary reflector.
		for j := 0; j < i; j++ {
			v.Data[j] = 0
		}
		v.Data[i] = 1
		for j := i + 1; j < c; j++ {
			v.Data[j] = lq.lq.mat.Data[i*lq.lq.mat.Stride+j]
		}

		// Compute the multiplication matrix.
		blas64.Ger(-lq.tau[i], v, v, h)
		qCopy.Copy(m)
		blas64.Gemm(blas.NoTrans, blas.NoTrans,
			1, h, qCopy.mat,
			0, m.mat)
	}
}

// SolveLQ finds a minimum-norm solution to a system of linear equations defined
// by the matrices A and b, where A is an m×n matrix represented in its LQ factorized
// form. If A is singular or near-singular a Condition error is returned. Please
// see the documentation for Condition for more information.
//
// The minimization problem solved depends on the input parameters.
//  If trans == false, find the minimum norm solution of A * X = b.
//  If trans == true, find X such that ||A*X - b||_2 is minimized.
// The solution matrix, X, is stored in place into the receiver.
func (m *Dense) SolveLQ(lq *LQ, trans bool, b Matrix) error {
	r, c := lq.lq.Dims()
	br, bc := b.Dims()

	// The LQ solve algorithm stores the result in-place into the right hand side.
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
	t := lq.lq.asTriDense(lq.lq.mat.Rows, blas.NonUnit, blas.Lower).mat
	if trans {
		work := make([]float64, 1)
		lapack64.Ormlq(blas.Left, blas.NoTrans, lq.lq.mat, lq.tau, x.mat, work, -1)
		work = make([]float64, int(work[0]))
		lapack64.Ormlq(blas.Left, blas.NoTrans, lq.lq.mat, lq.tau, x.mat, work, len(work))

		ok := lapack64.Trtrs(blas.Trans, t, x.mat)
		if !ok {
			return matrix.Condition(math.Inf(1))
		}
	} else {
		ok := lapack64.Trtrs(blas.NoTrans, t, x.mat)
		if !ok {
			return matrix.Condition(math.Inf(1))
		}
		for i := r; i < c; i++ {
			zero(x.mat.Data[i*x.mat.Stride : i*x.mat.Stride+bc])
		}
		work := make([]float64, 1)
		lapack64.Ormlq(blas.Left, blas.Trans, lq.lq.mat, lq.tau, x.mat, work, -1)
		work = make([]float64, int(work[0]))
		lapack64.Ormlq(blas.Left, blas.Trans, lq.lq.mat, lq.tau, x.mat, work, len(work))
	}
	// M was set above to be the correct size for the result.
	m.Copy(x)
	putWorkspace(x)
	if lq.cond > matrix.ConditionTolerance {
		return matrix.Condition(lq.cond)
	}
	return nil
}

// SolveLQVec finds a minimum-norm solution to a system of linear equations.
// Please see Dense.SolveLQ for the full documentation.
func (v *Vector) SolveLQVec(lq *LQ, trans bool, b *Vector) error {
	if v != b {
		v.checkOverlap(b.mat)
	}
	r, c := lq.lq.Dims()
	// The Solve implementation is non-trivial, so rather than duplicate the code,
	// instead recast the Vectors as Dense and call the matrix code.
	if trans {
		v.reuseAs(r)
	} else {
		v.reuseAs(c)
	}
	return v.asDense().SolveLQ(lq, trans, b.asDense())
}

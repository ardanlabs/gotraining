// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat64

import (
	"github.com/gonum/blas"
	"github.com/gonum/blas/blas64"
	"github.com/gonum/lapack/lapack64"
	"github.com/gonum/matrix"
)

// Solve finds a minimum-norm solution to a system of linear equations defined
// by the matrices a and b. If A is singular or near-singular, a Condition error
// is returned. Please see the documentation for Condition for more information.
//
// The minimization problem solved depends on the input parameters:
//  - if m >= n, find X such that ||A*X - B||_2 is minimized,
//  - if m < n, find the minimum norm solution of A * X = B.
// The solution matrix, X, is stored in-place into the receiver.
func (m *Dense) Solve(a, b Matrix) error {
	ar, ac := a.Dims()
	br, bc := b.Dims()
	if ar != br {
		panic(matrix.ErrShape)
	}
	m.reuseAs(ac, bc)

	// TODO(btracey): Add special cases for SymDense, etc.
	aU, aTrans := untranspose(a)
	bU, bTrans := untranspose(b)
	switch rma := aU.(type) {
	case RawTriangular:
		side := blas.Left
		tA := blas.NoTrans
		if aTrans {
			tA = blas.Trans
		}

		switch rm := bU.(type) {
		case RawMatrixer:
			if m != bU || bTrans {
				if m == bU || m.checkOverlap(rm.RawMatrix()) {
					tmp := getWorkspace(br, bc, false)
					tmp.Copy(b)
					m.Copy(tmp)
					putWorkspace(tmp)
					break
				}
				m.Copy(b)
			}
		default:
			if m != bU {
				m.Copy(b)
			} else if bTrans {
				// m and b share data so Copy cannot be used directly.
				tmp := getWorkspace(br, bc, false)
				tmp.Copy(b)
				m.Copy(tmp)
				putWorkspace(tmp)
			}
		}

		rm := rma.RawTriangular()
		blas64.Trsm(side, tA, 1, rm, m.mat)
		work := make([]float64, 3*rm.N)
		iwork := make([]int, rm.N)
		cond := lapack64.Trcon(matrix.CondNorm, rm, work, iwork)
		if cond > matrix.ConditionTolerance {
			return matrix.Condition(cond)
		}
		return nil
	}

	switch {
	case ar == ac:
		if a == b {
			// x = I.
			if ar == 1 {
				m.mat.Data[0] = 1
				return nil
			}
			for i := 0; i < ar; i++ {
				v := m.mat.Data[i*m.mat.Stride : i*m.mat.Stride+ac]
				zero(v)
				v[i] = 1
			}
			return nil
		}
		var lu LU
		lu.Factorize(a)
		return m.SolveLU(&lu, false, b)
	case ar > ac:
		var qr QR
		qr.Factorize(a)
		return m.SolveQR(&qr, false, b)
	default:
		var lq LQ
		lq.Factorize(a)
		return m.SolveLQ(&lq, false, b)
	}
}

// SolveVec finds a minimum-norm solution to a system of linear equations defined
// by the matrix a and the right-hand side vector b. If A is singular or
// near-singular, a Condition error is returned. Please see the documentation for
// Dense.Solve for more information.
func (v *Vector) SolveVec(a Matrix, b *Vector) error {
	if v != b {
		v.checkOverlap(b.mat)
	}
	_, c := a.Dims()
	// The Solve implementation is non-trivial, so rather than duplicate the code,
	// instead recast the Vectors as Dense and call the matrix code.
	v.reuseAs(c)
	m := v.asDense()
	// We conditionally create bm as m when b and v are identical
	// to prevent the overlap detection code from identifying m
	// and bm as overlapping but not identical.
	bm := m
	if v != b {
		bm = b.asDense()
	}
	return m.Solve(a, bm)
}

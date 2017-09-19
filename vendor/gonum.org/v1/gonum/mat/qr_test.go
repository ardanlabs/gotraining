// Copyright ©2013 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat

import (
	"math"
	"math/rand"
	"testing"

	"gonum.org/v1/gonum/blas/blas64"
)

func TestQR(t *testing.T) {
	for _, test := range []struct {
		m, n int
	}{
		{5, 5},
		{10, 5},
	} {
		m := test.m
		n := test.n
		a := NewDense(m, n, nil)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				a.Set(i, j, rand.NormFloat64())
			}
		}
		var want Dense
		want.Clone(a)

		var qr QR
		qr.Factorize(a)
		q := qr.QTo(nil)

		if !isOrthonormal(q, 1e-10) {
			t.Errorf("Q is not orthonormal: m = %v, n = %v", m, n)
		}

		r := qr.RTo(nil)

		var got Dense
		got.Mul(q, r)
		if !EqualApprox(&got, &want, 1e-12) {
			t.Errorf("QR does not equal original matrix. \nWant: %v\nGot: %v", want, got)
		}
	}
}

func isOrthonormal(q *Dense, tol float64) bool {
	m, n := q.Dims()
	if m != n {
		return false
	}
	for i := 0; i < m; i++ {
		for j := i; j < m; j++ {
			dot := blas64.Dot(m,
				blas64.Vector{Inc: 1, Data: q.mat.Data[i*q.mat.Stride:]},
				blas64.Vector{Inc: 1, Data: q.mat.Data[j*q.mat.Stride:]},
			)
			// Dot product should be 1 if i == j and 0 otherwise.
			if i == j && math.Abs(dot-1) > tol {
				return false
			}
			if i != j && math.Abs(dot) > tol {
				return false
			}
		}
	}
	return true
}

func TestSolveQR(t *testing.T) {
	for _, trans := range []bool{false, true} {
		for _, test := range []struct {
			m, n, bc int
		}{
			{5, 5, 1},
			{10, 5, 1},
			{5, 5, 3},
			{10, 5, 3},
		} {
			m := test.m
			n := test.n
			bc := test.bc
			a := NewDense(m, n, nil)
			for i := 0; i < m; i++ {
				for j := 0; j < n; j++ {
					a.Set(i, j, rand.Float64())
				}
			}
			br := m
			if trans {
				br = n
			}
			b := NewDense(br, bc, nil)
			for i := 0; i < br; i++ {
				for j := 0; j < bc; j++ {
					b.Set(i, j, rand.Float64())
				}
			}
			var x Dense
			var qr QR
			qr.Factorize(a)
			qr.Solve(&x, trans, b)

			// Test that the normal equations hold.
			// A^T * A * x = A^T * b if !trans
			// A * A^T * x = A * b if trans
			var lhs Dense
			var rhs Dense
			if trans {
				var tmp Dense
				tmp.Mul(a, a.T())
				lhs.Mul(&tmp, &x)
				rhs.Mul(a, b)
			} else {
				var tmp Dense
				tmp.Mul(a.T(), a)
				lhs.Mul(&tmp, &x)
				rhs.Mul(a.T(), b)
			}
			if !EqualApprox(&lhs, &rhs, 1e-10) {
				t.Errorf("Normal equations do not hold.\nLHS: %v\n, RHS: %v\n", lhs, rhs)
			}
		}
	}
	// TODO(btracey): Add in testOneInput when it exists.
}

func TestSolveQRVec(t *testing.T) {
	for _, trans := range []bool{false, true} {
		for _, test := range []struct {
			m, n int
		}{
			{5, 5},
			{10, 5},
		} {
			m := test.m
			n := test.n
			a := NewDense(m, n, nil)
			for i := 0; i < m; i++ {
				for j := 0; j < n; j++ {
					a.Set(i, j, rand.Float64())
				}
			}
			br := m
			if trans {
				br = n
			}
			b := NewVecDense(br, nil)
			for i := 0; i < br; i++ {
				b.SetVec(i, rand.Float64())
			}
			var x VecDense
			var qr QR
			qr.Factorize(a)
			qr.SolveVec(&x, trans, b)

			// Test that the normal equations hold.
			// A^T * A * x = A^T * b if !trans
			// A * A^T * x = A * b if trans
			var lhs Dense
			var rhs Dense
			if trans {
				var tmp Dense
				tmp.Mul(a, a.T())
				lhs.Mul(&tmp, &x)
				rhs.Mul(a, b)
			} else {
				var tmp Dense
				tmp.Mul(a.T(), a)
				lhs.Mul(&tmp, &x)
				rhs.Mul(a.T(), b)
			}
			if !EqualApprox(&lhs, &rhs, 1e-10) {
				t.Errorf("Normal equations do not hold.\nLHS: %v\n, RHS: %v\n", lhs, rhs)
			}
		}
	}
	// TODO(btracey): Add in testOneInput when it exists.
}

func TestSolveQRCond(t *testing.T) {
	for _, test := range []*Dense{
		NewDense(2, 2, []float64{1, 0, 0, 1e-20}),
		NewDense(3, 2, []float64{1, 0, 0, 1e-20, 0, 0}),
	} {
		m, _ := test.Dims()
		var qr QR
		qr.Factorize(test)
		b := NewDense(m, 2, nil)
		var x Dense
		if err := qr.Solve(&x, false, b); err == nil {
			t.Error("No error for near-singular matrix in matrix solve.")
		}

		bvec := NewVecDense(m, nil)
		var xvec VecDense
		if err := qr.SolveVec(&xvec, false, bvec); err == nil {
			t.Error("No error for near-singular matrix in matrix solve.")
		}
	}
}

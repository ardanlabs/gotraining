// Copyright ©2013 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat64

import (
	"math/rand"
	"testing"
)

func TestLQ(t *testing.T) {
	for _, test := range []struct {
		m, n int
	}{
		{5, 5},
		{5, 10},
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

		lq := &LQ{}
		lq.Factorize(a)
		var l, q Dense
		q.QFromLQ(lq)

		if !isOrthonormal(&q, 1e-10) {
			t.Errorf("Q is not orthonormal: m = %v, n = %v", m, n)
		}

		l.LFromLQ(lq)

		var got Dense
		got.Mul(&l, &q)
		if !EqualApprox(&got, &want, 1e-12) {
			t.Errorf("LQ does not equal original matrix. \nWant: %v\nGot: %v", want, got)
		}
	}
}

func TestSolveLQ(t *testing.T) {
	for _, trans := range []bool{false, true} {
		for _, test := range []struct {
			m, n, bc int
		}{
			{5, 5, 1},
			{5, 10, 1},
			{5, 5, 3},
			{5, 10, 3},
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
			lq := &LQ{}
			lq.Factorize(a)
			x.SolveLQ(lq, trans, b)

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

func TestSolveLQVec(t *testing.T) {
	for _, trans := range []bool{false, true} {
		for _, test := range []struct {
			m, n int
		}{
			{5, 5},
			{5, 10},
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
			b := NewVector(br, nil)
			for i := 0; i < br; i++ {
				b.SetVec(i, rand.Float64())
			}
			var x Vector
			lq := &LQ{}
			lq.Factorize(a)
			x.SolveLQVec(lq, trans, b)

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

func TestSolveLQCond(t *testing.T) {
	for _, test := range []*Dense{
		NewDense(2, 2, []float64{1, 0, 0, 1e-20}),
		NewDense(2, 3, []float64{1, 0, 0, 0, 1e-20, 0}),
	} {
		m, _ := test.Dims()
		var lq LQ
		lq.Factorize(test)
		b := NewDense(m, 2, nil)
		var x Dense
		if err := x.SolveLQ(&lq, false, b); err == nil {
			t.Error("No error for near-singular matrix in matrix solve.")
		}

		bvec := NewVector(m, nil)
		var xvec Vector
		if err := xvec.SolveLQVec(&lq, false, bvec); err == nil {
			t.Error("No error for near-singular matrix in matrix solve.")
		}
	}
}

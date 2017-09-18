// Copyright ©2013 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat

import (
	"math/rand"
	"testing"
)

func TestLUD(t *testing.T) {
	for _, n := range []int{1, 5, 10, 11, 50} {
		a := NewDense(n, n, nil)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				a.Set(i, j, rand.NormFloat64())
			}
		}
		var want Dense
		want.Clone(a)

		var lu LU
		lu.Factorize(a)

		l := lu.LTo(nil)
		u := lu.UTo(nil)
		var p Dense
		pivot := lu.Pivot(nil)
		p.Permutation(n, pivot)
		var got Dense
		got.Product(&p, l, u)
		if !EqualApprox(&got, &want, 1e-12) {
			t.Errorf("PLU does not equal original matrix.\nWant: %v\n Got: %v", want, got)
		}
	}
}

func TestLURankOne(t *testing.T) {
	for _, pivoting := range []bool{true} {
		for _, n := range []int{3, 10, 50} {
			// Construct a random LU factorization
			lu := &LU{}
			lu.lu = NewDense(n, n, nil)
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					lu.lu.Set(i, j, rand.Float64())
				}
			}
			lu.pivot = make([]int, n)
			for i := range lu.pivot {
				lu.pivot[i] = i
			}
			if pivoting {
				// For each row, randomly swap with itself or a row after (like is done)
				// in the actual LU factorization.
				for i := range lu.pivot {
					idx := i + rand.Intn(n-i)
					lu.pivot[i], lu.pivot[idx] = lu.pivot[idx], lu.pivot[i]
				}
			}
			// Apply a rank one update. Ensure the update magnitude is larger than
			// the equal tolerance.
			alpha := rand.Float64() + 1
			x := NewVecDense(n, nil)
			y := NewVecDense(n, nil)
			for i := 0; i < n; i++ {
				x.setVec(i, rand.Float64()+1)
				y.setVec(i, rand.Float64()+1)
			}
			a := luReconstruct(lu)
			a.RankOne(a, alpha, x, y)

			var luNew LU
			luNew.RankOne(lu, alpha, x, y)
			lu.RankOne(lu, alpha, x, y)

			aR1New := luReconstruct(&luNew)
			aR1 := luReconstruct(lu)

			if !Equal(aR1, aR1New) {
				t.Error("Different answer when new receiver")
			}
			if !EqualApprox(aR1, a, 1e-10) {
				t.Errorf("Rank one mismatch, pivot %v.\nWant: %v\nGot:%v\n", pivoting, a, aR1)
			}
		}
	}
}

// luReconstruct reconstructs the original A matrix from an LU decomposition.
func luReconstruct(lu *LU) *Dense {
	var L, U TriDense
	lu.LTo(&L)
	lu.UTo(&U)
	var P Dense
	pivot := lu.Pivot(nil)
	P.Permutation(len(pivot), pivot)

	var a Dense
	a.Mul(&L, &U)
	a.Mul(&P, &a)
	return &a
}

func TestSolveLU(t *testing.T) {
	for _, test := range []struct {
		n, bc int
	}{
		{5, 5},
		{5, 10},
		{10, 5},
	} {
		n := test.n
		bc := test.bc
		a := NewDense(n, n, nil)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				a.Set(i, j, rand.NormFloat64())
			}
		}
		b := NewDense(n, bc, nil)
		for i := 0; i < n; i++ {
			for j := 0; j < bc; j++ {
				b.Set(i, j, rand.NormFloat64())
			}
		}
		var lu LU
		lu.Factorize(a)
		var x Dense
		if err := lu.Solve(&x, false, b); err != nil {
			continue
		}
		var got Dense
		got.Mul(a, &x)
		if !EqualApprox(&got, b, 1e-12) {
			t.Errorf("Solve mismatch for non-singular matrix. n = %v, bc = %v.\nWant: %v\nGot: %v", n, bc, b, got)
		}
	}
	// TODO(btracey): Add testOneInput test when such a function exists.
}

func TestSolveLUCond(t *testing.T) {
	for _, test := range []*Dense{
		NewDense(2, 2, []float64{1, 0, 0, 1e-20}),
	} {
		m, _ := test.Dims()
		var lu LU
		lu.Factorize(test)
		b := NewDense(m, 2, nil)
		var x Dense
		if err := lu.Solve(&x, false, b); err == nil {
			t.Error("No error for near-singular matrix in matrix solve.")
		}

		bvec := NewVecDense(m, nil)
		var xvec VecDense
		if err := lu.SolveVec(&xvec, false, bvec); err == nil {
			t.Error("No error for near-singular matrix in matrix solve.")
		}
	}
}

func TestSolveLUVec(t *testing.T) {
	for _, n := range []int{5, 10} {
		a := NewDense(n, n, nil)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				a.Set(i, j, rand.NormFloat64())
			}
		}
		b := NewVecDense(n, nil)
		for i := 0; i < n; i++ {
			b.SetVec(i, rand.NormFloat64())
		}
		var lu LU
		lu.Factorize(a)
		var x VecDense
		if err := lu.SolveVec(&x, false, b); err != nil {
			continue
		}
		var got VecDense
		got.MulVec(a, &x)
		if !EqualApprox(&got, b, 1e-12) {
			t.Errorf("Solve mismatch n = %v.\nWant: %v\nGot: %v", n, b, got)
		}
	}
	// TODO(btracey): Add testOneInput test when such a function exists.
}

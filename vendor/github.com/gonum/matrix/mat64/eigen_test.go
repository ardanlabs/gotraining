// Copyright ©2013 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat64

import (
	"math/rand"
	"testing"

	"github.com/gonum/floats"
)

func TestEigen(t *testing.T) {
	for i, test := range []struct {
		a *Dense

		values []complex128
		left   *Dense
		right  *Dense
	}{
		{
			a: NewDense(3, 3, []float64{
				1, 0, 0,
				0, 1, 0,
				0, 0, 1,
			}),
			values: []complex128{1, 1, 1},
			left: NewDense(3, 3, []float64{
				1, 0, 0,
				0, 1, 0,
				0, 0, 1,
			}),
			right: NewDense(3, 3, []float64{
				1, 0, 0,
				0, 1, 0,
				0, 0, 1,
			}),
		},
	} {
		var e1, e2, e3, e4 Eigen
		ok := e1.Factorize(test.a, true, true)
		if !ok {
			panic("bad factorization")
		}
		e2.Factorize(test.a, false, true)
		e3.Factorize(test.a, true, false)
		e4.Factorize(test.a, false, false)

		v1 := e1.Values(nil)
		if !cmplxEqual(v1, test.values) {
			t.Errorf("eigenvector mismatch. Case %v", i)
		}
		if !Equal(e1.LeftVectors(), test.left) {
			t.Errorf("left eigenvector mismatch. Case %v", i)
		}
		if !Equal(e1.Vectors(), test.right) {
			t.Errorf("right eigenvector mismatch. Case %v", i)
		}

		// Check that the eigenvectors and values are the same in all combinations.
		if !cmplxEqual(v1, e2.Values(nil)) {
			t.Errorf("eigenvector mismatch. Case %v", i)
		}
		if !cmplxEqual(v1, e3.Values(nil)) {
			t.Errorf("eigenvector mismatch. Case %v", i)
		}
		if !cmplxEqual(v1, e4.Values(nil)) {
			t.Errorf("eigenvector mismatch. Case %v", i)
		}
		if !Equal(e1.Vectors(), e2.Vectors()) {
			t.Errorf("right eigenvector mismatch. Case %v", i)
		}
		if !Equal(e1.LeftVectors(), e3.LeftVectors()) {
			t.Errorf("right eigenvector mismatch. Case %v", i)
		}

		// TODO(btracey): Also add in a test for correctness when #308 is
		// resolved and we have a CMat.Mul().
	}
}

func cmplxEqual(v1, v2 []complex128) bool {
	for i, v := range v1 {
		if v != v2[i] {
			return false
		}
	}
	return true
}

func TestSymEigen(t *testing.T) {
	// Hand coded tests with results from lapack.
	for _, test := range []struct {
		mat *SymDense

		values  []float64
		vectors *Dense
	}{
		{
			mat:    NewSymDense(3, []float64{8, 2, 4, 2, 6, 10, 4, 10, 5}),
			values: []float64{-4.707679201365891, 6.294580208480216, 17.413098992885672},
			vectors: NewDense(3, 3, []float64{
				-0.127343483135656, -0.902414161226903, -0.411621572466779,
				-0.664177720955769, 0.385801900032553, -0.640331827193739,
				0.736648893495999, 0.191847792659746, -0.648492738712395,
			}),
		},
	} {
		var es EigenSym
		ok := es.Factorize(test.mat, true)
		if !ok {
			t.Errorf("bad factorization")
		}
		if !floats.EqualApprox(test.values, es.values, 1e-14) {
			t.Errorf("Eigenvalue mismatch")
		}
		if !EqualApprox(test.vectors, es.vectors, 1e-14) {
			t.Errorf("Eigenvector mismatch")
		}

		var es2 EigenSym
		es2.Factorize(test.mat, false)
		if !floats.EqualApprox(es2.values, es.values, 1e-14) {
			t.Errorf("Eigenvalue mismatch when no vectors computed")
		}
	}

	// Randomized tests
	rnd := rand.New(rand.NewSource(1))
	for _, n := range []int{3, 5, 10, 70} {
		for cas := 0; cas < 10; cas++ {
			a := make([]float64, n*n)
			for i := range a {
				a[i] = rnd.NormFloat64()
			}
			s := NewSymDense(n, a)
			var es EigenSym
			ok := es.Factorize(s, true)
			if !ok {
				t.Errorf("Bad test")
			}

			// Check that the eigenvectors are orthonormal.
			if !isOrthonormal(es.vectors, 1e-8) {
				t.Errorf("Eigenvectors not orthonormal")
			}

			// Check that the eigenvalues are actually eigenvalues.
			for i := 0; i < n; i++ {
				v := NewVector(n, Col(nil, i, es.vectors))
				var m Vector
				m.MulVec(s, v)

				var scal Vector
				scal.ScaleVec(es.values[i], v)

				if !EqualApprox(&m, &scal, 1e-8) {
					t.Errorf("Eigenvalue does not match")
				}
			}
		}
	}
}

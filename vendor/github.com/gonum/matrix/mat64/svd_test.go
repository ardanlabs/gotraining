// Copyright ©2013 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat64

import (
	"math/rand"
	"testing"

	"github.com/gonum/floats"
	"github.com/gonum/matrix"
)

func TestSVD(t *testing.T) {
	// Hand coded tests
	for _, test := range []struct {
		a *Dense
		u *Dense
		v *Dense
		s []float64
	}{
		{
			a: NewDense(4, 2, []float64{2, 4, 1, 3, 0, 0, 0, 0}),
			u: NewDense(4, 2, []float64{
				-0.8174155604703632, -0.5760484367663209,
				-0.5760484367663209, 0.8174155604703633,
				0, 0,
				0, 0,
			}),
			v: NewDense(2, 2, []float64{
				-0.4045535848337571, -0.9145142956773044,
				-0.9145142956773044, 0.4045535848337571,
			}),
			s: []float64{5.464985704219041, 0.365966190626258},
		},
		{
			// Issue #5.
			a: NewDense(3, 11, []float64{
				1, 1, 0, 1, 0, 0, 0, 0, 0, 11, 1,
				1, 0, 0, 0, 0, 0, 1, 0, 0, 12, 2,
				1, 1, 0, 0, 0, 0, 0, 0, 1, 13, 3,
			}),
			u: NewDense(3, 3, []float64{
				-0.5224167862273765, 0.7864430360363114, 0.3295270133658976,
				-0.5739526766688285, -0.03852203026050301, -0.8179818935216693,
				-0.6306021141833781, -0.6164603833618163, 0.4715056408282468,
			}),
			v: NewDense(11, 3, []float64{
				-0.08123293141915189, 0.08528085505260324, -0.013165501690885152,
				-0.05423546426886932, 0.1102707844980355, 0.622210623111631,
				0, 0, 0,
				-0.0245733326078166, 0.510179651760153, 0.25596360803140994,
				0, 0, 0,
				0, 0, 0,
				-0.026997467150282436, -0.024989929445430496, -0.6353761248025164,
				0, 0, 0,
				-0.029662131661052707, -0.3999088672621176, 0.3662470150802212,
				-0.9798839760830571, 0.11328174160898856, -0.047702613241813366,
				-0.16755466189153964, -0.7395268089170608, 0.08395240366704032,
			}),
			s: []float64{21.259500881097434, 1.5415021616856566, 1.2873979074613628},
		},
	} {
		var svd SVD
		ok := svd.Factorize(test.a, matrix.SVDThin)
		if !ok {
			t.Errorf("SVD failed")
		}
		s, u, v := extractSVD(&svd)
		if !floats.EqualApprox(s, test.s, 1e-10) {
			t.Errorf("Singular value mismatch. Got %v, want %v.", s, test.s)
		}
		if !EqualApprox(u, test.u, 1e-10) {
			t.Errorf("U mismatch.\nGot:\n%v\nWant:\n%v", Formatted(u), Formatted(test.u))
		}
		if !EqualApprox(v, test.v, 1e-10) {
			t.Errorf("V mismatch.\nGot:\n%v\nWant:\n%v", Formatted(v), Formatted(test.v))
		}
		m, n := test.a.Dims()
		sigma := NewDense(min(m, n), min(m, n), nil)
		for i := 0; i < min(m, n); i++ {
			sigma.Set(i, i, s[i])
		}

		var ans Dense
		ans.Product(u, sigma, v.T())
		if !EqualApprox(test.a, &ans, 1e-10) {
			t.Errorf("A reconstruction mismatch.\nGot:\n%v\nWant:\n%v\n", Formatted(&ans), Formatted(test.a))
		}
	}

	for _, test := range []struct {
		m, n int
	}{
		{5, 5},
		{5, 3},
		{3, 5},
		{150, 150},
		{200, 150},
		{150, 200},
	} {
		m := test.m
		n := test.n
		for trial := 0; trial < 10; trial++ {
			a := NewDense(m, n, nil)
			for i := range a.mat.Data {
				a.mat.Data[i] = rand.NormFloat64()
			}
			aCopy := DenseCopyOf(a)

			// Test Full decomposition.
			var svd SVD
			ok := svd.Factorize(a, matrix.SVDFull)
			if !ok {
				t.Errorf("SVD factorization failed")
			}
			if !Equal(a, aCopy) {
				t.Errorf("A changed during call to SVD with full")
			}
			s, u, v := extractSVD(&svd)
			sigma := NewDense(m, n, nil)
			for i := 0; i < min(m, n); i++ {
				sigma.Set(i, i, s[i])
			}
			var ansFull Dense
			ansFull.Product(u, sigma, v.T())
			if !EqualApprox(&ansFull, a, 1e-8) {
				t.Errorf("Answer mismatch when SVDFull")
			}

			// Test Thin decomposition.
			ok = svd.Factorize(a, matrix.SVDThin)
			if !ok {
				t.Errorf("SVD factorization failed")
			}
			if !Equal(a, aCopy) {
				t.Errorf("A changed during call to SVD with Thin")
			}
			sThin, u, v := extractSVD(&svd)
			if !floats.EqualApprox(s, sThin, 1e-8) {
				t.Errorf("Singular value mismatch between Full and Thin decomposition")
			}
			sigma = NewDense(min(m, n), min(m, n), nil)
			for i := 0; i < min(m, n); i++ {
				sigma.Set(i, i, sThin[i])
			}
			ansFull.Reset()
			ansFull.Product(u, sigma, v.T())
			if !EqualApprox(&ansFull, a, 1e-8) {
				t.Errorf("Answer mismatch when SVDFull")
			}

			// Test None decomposition.
			ok = svd.Factorize(a, matrix.SVDNone)
			if !ok {
				t.Errorf("SVD factorization failed")
			}
			if !Equal(a, aCopy) {
				t.Errorf("A changed during call to SVD with none")
			}
			sNone := make([]float64, min(m, n))
			svd.Values(sNone)
			if !floats.EqualApprox(s, sNone, 1e-8) {
				t.Errorf("Singular value mismatch between Full and None decomposition")
			}
		}
	}
}

func extractSVD(svd *SVD) (s []float64, u, v *Dense) {
	var um, vm Dense
	um.UFromSVD(svd)
	vm.VFromSVD(svd)
	s = svd.Values(nil)
	return s, &um, &vm
}

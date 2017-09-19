// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat64

import (
	"math/rand"
	"testing"
)

func TestSolve(t *testing.T) {
	// Hand-coded cases.
	for _, test := range []struct {
		a         [][]float64
		b         [][]float64
		ans       [][]float64
		shouldErr bool
	}{
		{
			a:         [][]float64{{6}},
			b:         [][]float64{{3}},
			ans:       [][]float64{{0.5}},
			shouldErr: false,
		},
		{
			a: [][]float64{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			b: [][]float64{
				{3},
				{2},
				{1},
			},
			ans: [][]float64{
				{3},
				{2},
				{1},
			},
			shouldErr: false,
		},
		{
			a: [][]float64{
				{0.8147, 0.9134, 0.5528},
				{0.9058, 0.6324, 0.8723},
				{0.1270, 0.0975, 0.7612},
			},
			b: [][]float64{
				{0.278},
				{0.547},
				{0.958},
			},
			ans: [][]float64{
				{-0.932687281002860},
				{0.303963920182067},
				{1.375216503507109},
			},
			shouldErr: false,
		},
		{
			a: [][]float64{
				{0.8147, 0.9134, 0.5528},
				{0.9058, 0.6324, 0.8723},
			},
			b: [][]float64{
				{0.278},
				{0.547},
			},
			ans: [][]float64{
				{0.25919787248965376},
				{-0.25560256266441034},
				{0.5432324059702451},
			},
			shouldErr: false,
		},
		{
			a: [][]float64{
				{0.8147, 0.9134, 0.9},
				{0.9058, 0.6324, 0.9},
				{0.1270, 0.0975, 0.1},
				{1.6, 2.8, -3.5},
			},
			b: [][]float64{
				{0.278},
				{0.547},
				{-0.958},
				{1.452},
			},
			ans: [][]float64{
				{0.820970340787782},
				{-0.218604626527306},
				{-0.212938815234215},
			},
			shouldErr: false,
		},
		{
			a: [][]float64{
				{0.8147, 0.9134, 0.231, -1.65},
				{0.9058, 0.6324, 0.9, 0.72},
				{0.1270, 0.0975, 0.1, 1.723},
				{1.6, 2.8, -3.5, 0.987},
				{7.231, 9.154, 1.823, 0.9},
			},
			b: [][]float64{
				{0.278, 8.635},
				{0.547, 9.125},
				{-0.958, -0.762},
				{1.452, 1.444},
				{1.999, -7.234},
			},
			ans: [][]float64{
				{1.863006789511373, 44.467887791812750},
				{-1.127270935407224, -34.073794226035126},
				{-0.527926457947330, -8.032133759788573},
				{-0.248621916204897, -2.366366415805275},
			},
			shouldErr: false,
		},
		{
			a: [][]float64{
				{0, 0},
				{0, 0},
			},
			b: [][]float64{
				{3},
				{2},
			},
			ans:       nil,
			shouldErr: true,
		},
		{
			a: [][]float64{
				{0, 0},
				{0, 0},
				{0, 0},
			},
			b: [][]float64{
				{3},
				{2},
				{1},
			},
			ans:       nil,
			shouldErr: true,
		},
		{
			a: [][]float64{
				{0, 0, 0},
				{0, 0, 0},
			},
			b: [][]float64{
				{3},
				{2},
			},
			ans:       nil,
			shouldErr: true,
		},
	} {
		a := NewDense(flatten(test.a))
		b := NewDense(flatten(test.b))

		var ans *Dense
		if test.ans != nil {
			ans = NewDense(flatten(test.ans))
		}

		var x Dense
		err := x.Solve(a, b)
		if err != nil {
			if !test.shouldErr {
				t.Errorf("Unexpected solve error: %s", err)
			}
			continue
		}
		if err == nil && test.shouldErr {
			t.Errorf("Did not error during solve.")
			continue
		}
		if !EqualApprox(&x, ans, 1e-12) {
			t.Errorf("Solve answer mismatch. Want %v, got %v", ans, x)
		}
	}

	// Random Cases.
	for _, test := range []struct {
		m, n, bc int
	}{
		{5, 5, 1},
		{5, 10, 1},
		{10, 5, 1},
		{5, 5, 7},
		{5, 10, 7},
		{10, 5, 7},
		{5, 5, 12},
		{5, 10, 12},
		{10, 5, 12},
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
		b := NewDense(br, bc, nil)
		for i := 0; i < br; i++ {
			for j := 0; j < bc; j++ {
				b.Set(i, j, rand.Float64())
			}
		}
		var x Dense
		x.Solve(a, b)

		// Test that the normal equations hold.
		// A^T * A * x = A^T * b
		var tmp, lhs, rhs Dense
		tmp.Mul(a.T(), a)
		lhs.Mul(&tmp, &x)
		rhs.Mul(a.T(), b)
		if !EqualApprox(&lhs, &rhs, 1e-10) {
			t.Errorf("Normal equations do not hold.\nLHS: %v\n, RHS: %v\n", lhs, rhs)
		}
	}

	// Use testTwoInput.
	method := func(receiver, a, b Matrix) {
		type Solver interface {
			Solve(a, b Matrix) error
		}
		rd := receiver.(Solver)
		rd.Solve(a, b)
	}
	denseComparison := func(receiver, a, b *Dense) {
		receiver.Solve(a, b)
	}
	testTwoInput(t, "Solve", &Dense{}, method, denseComparison, legalTypesAll, legalSizeSolve, 1e-7)
}

func TestSolveVec(t *testing.T) {
	for _, test := range []struct {
		m, n int
	}{
		{5, 5},
		{5, 10},
		{10, 5},
		{5, 5},
		{5, 10},
		{10, 5},
		{5, 5},
		{5, 10},
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
		b := NewVector(br, nil)
		for i := 0; i < br; i++ {
			b.SetVec(i, rand.Float64())
		}
		var x Vector
		x.SolveVec(a, b)

		// Test that the normal equations hold.
		// A^T * A * x = A^T * b
		var tmp, lhs, rhs Dense
		tmp.Mul(a.T(), a)
		lhs.Mul(&tmp, &x)
		rhs.Mul(a.T(), b)
		if !EqualApprox(&lhs, &rhs, 1e-10) {
			t.Errorf("Normal equations do not hold.\nLHS: %v\n, RHS: %v\n", lhs, rhs)
		}
	}

	// Use testTwoInput
	method := func(receiver, a, b Matrix) {
		type SolveVecer interface {
			SolveVec(a Matrix, b *Vector) error
		}
		rd := receiver.(SolveVecer)
		rd.SolveVec(a, b.(*Vector))
	}
	denseComparison := func(receiver, a, b *Dense) {
		receiver.Solve(a, b)
	}
	testTwoInput(t, "SolveVec", &Vector{}, method, denseComparison, legalTypesNotVecVec, legalSizeSolve, 1e-12)
}

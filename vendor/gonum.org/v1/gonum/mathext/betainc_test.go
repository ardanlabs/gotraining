// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mathext

import (
	"testing"

	"gonum.org/v1/gonum/floats"
)

func TestIncBeta(t *testing.T) {
	tol := 1e-14
	tol2 := 1e-10
	// Test against values from scipy
	for i, test := range []struct {
		a, b, x, ans float64
	}{
		{1, 1, 0.8, 0.8},
		{1, 5, 0.8, 0.99968000000000001},
		{10, 10, 0.8, 0.99842087945083291},
		{10, 10, 0.1, 3.929882327128003e-06},
		{10, 2, 0.4, 0.00073400320000000028},
		{0.1, 0.2, 0.6, 0.69285678232066683},
		{1, 10, 0.7489, 0.99999900352334858},
	} {
		y := RegIncBeta(test.a, test.b, test.x)
		if !floats.EqualWithinAbsOrRel(y, test.ans, tol, tol) {
			t.Errorf("Incomplete beta mismatch. Case %v: Got %v, want %v", i, y, test.ans)
		}

		yc := 1 - RegIncBeta(test.b, test.a, 1-test.x)
		if !floats.EqualWithinAbsOrRel(y, yc, tol, tol) {
			t.Errorf("Incomplete beta complementary mismatch. Case %v: Got %v, want %v", i, y, yc)
		}

		x := InvRegIncBeta(test.a, test.b, y)
		if !floats.EqualWithinAbsOrRel(x, test.x, tol2, tol2) {
			t.Errorf("Inverse incomplete beta mismatch. Case %v: Got %v, want %v", i, x, test.x)
		}
	}

	// Confirm that Invincbeta and Incbeta agree. Sweep over a variety of
	// a, b, and y values.
	tol = 1e-6
	steps := 201
	ints := make([]float64, steps)
	floats.Span(ints, 0, 1)

	sz := 51
	min := 1e-2
	max := 1e2
	as := make([]float64, sz)
	floats.LogSpan(as, min, max)
	bs := make([]float64, sz)
	floats.LogSpan(bs, min, max)

	for _, a := range as {
		for _, b := range bs {
			for _, yr := range ints {
				x := InvRegIncBeta(a, b, yr)
				if x > 1-1e-6 {
					// Numerical error too large
					continue
				}
				y := RegIncBeta(a, b, x)
				if !floats.EqualWithinAbsOrRel(yr, y, tol, tol) {
					t.Errorf("Mismatch between inv inc beta and inc beta. a = %v, b = %v, x = %v, got %v, want %v.", a, b, x, y, yr)
					break
				}
			}
		}
	}
}

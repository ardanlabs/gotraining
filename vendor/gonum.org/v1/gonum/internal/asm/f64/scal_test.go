// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package f64

import (
	"fmt"
	"math/rand"
	"testing"
)

var scalTests = []struct {
	alpha float64
	x     []float64
	want  []float64
}{
	{
		alpha: 0,
		x:     []float64{},
		want:  []float64{},
	},
	{
		alpha: 0,
		x:     []float64{1},
		want:  []float64{0},
	},
	{
		alpha: 1,
		x:     []float64{1},
		want:  []float64{1},
	},
	{
		alpha: 2,
		x:     []float64{1, -2},
		want:  []float64{2, -4},
	},
	{
		alpha: 2,
		x:     []float64{1, -2, 3},
		want:  []float64{2, -4, 6},
	},
	{
		alpha: 2,
		x:     []float64{1, -2, 3, 4},
		want:  []float64{2, -4, 6, 8},
	},
	{
		alpha: 2,
		x:     []float64{1, -2, 3, 4, -5},
		want:  []float64{2, -4, 6, 8, -10},
	},
	{
		alpha: 2,
		x:     []float64{0, 1, -2, 3, 4, -5, 6, -7},
		want:  []float64{0, 2, -4, 6, 8, -10, 12, -14},
	},
	{
		alpha: 2,
		x:     []float64{0, 1, -2, 3, 4, -5, 6, -7, 8},
		want:  []float64{0, 2, -4, 6, 8, -10, 12, -14, 16},
	},
	{
		alpha: 2,
		x:     []float64{0, 1, -2, 3, 4, -5, 6, -7, 8, 9},
		want:  []float64{0, 2, -4, 6, 8, -10, 12, -14, 16, 18},
	},
	{
		alpha: 3,
		x:     []float64{0, 1, -2, 3, 4, -5, 6, -7, 8, 9, 12},
		want:  []float64{0, 3, -6, 9, 12, -15, 18, -21, 24, 27, 36},
	},
}

func TestScalUnitary(t *testing.T) {
	const xGdVal = -0.5
	for i, test := range scalTests {
		for _, align := range align1 {
			prefix := fmt.Sprintf("Test %v (x:%v)", i, align)
			xgLn := 4 + align
			xg := guardVector(test.x, xGdVal, xgLn)
			x := xg[xgLn : len(xg)-xgLn]

			ScalUnitary(test.alpha, x)

			for i := range test.want {
				if !same(x[i], test.want[i]) {
					t.Errorf(msgVal, prefix, i, x[i], test.want[i])
				}
			}
			if !isValidGuard(xg, xGdVal, xgLn) {
				t.Errorf(msgGuard, prefix, "x", xg[:xgLn], xg[len(xg)-xgLn:])
			}
		}
	}
}

func TestScalUnitaryTo(t *testing.T) {
	const xGdVal, dstGdVal = -1, 0.5
	rng := rand.New(rand.NewSource(42))
	for i, test := range scalTests {
		n := len(test.x)
		for _, align := range align2 {
			prefix := fmt.Sprintf("Test %v (x:%v dst:%v)", i, align.x, align.y)
			xgLn, dgLn := 4+align.x, 4+align.y
			xg := guardVector(test.x, xGdVal, xgLn)
			dg := guardVector(randSlice(n, 1, rng), dstGdVal, dgLn)
			x, dst := xg[xgLn:len(xg)-xgLn], dg[dgLn:len(dg)-dgLn]

			ScalUnitaryTo(dst, test.alpha, x)

			for i := range test.want {
				if !same(dst[i], test.want[i]) {
					t.Errorf(msgVal, prefix, i, dst[i], test.want[i])
				}
			}
			if !isValidGuard(xg, xGdVal, xgLn) {
				t.Errorf(msgGuard, prefix, "x", xg[:xgLn], xg[len(xg)-xgLn:])
			}
			if !isValidGuard(dg, dstGdVal, dgLn) {
				t.Errorf(msgGuard, prefix, "y", dg[:dgLn], dg[len(dg)-dgLn:])
			}
			if !equalStrided(test.x, x, 1) {
				t.Errorf("%v: modified read-only x argument", prefix)
			}
		}
	}
}

func TestScalInc(t *testing.T) {
	const xGdVal = -0.5
	gdLn := 4
	for i, test := range scalTests {
		n := len(test.x)
		for _, incX := range []int{1, 2, 3, 4, 7, 10} {
			prefix := fmt.Sprintf("Test %v (x:%v)", i, incX)
			xg := guardIncVector(test.x, xGdVal, incX, gdLn)
			x := xg[gdLn : len(xg)-gdLn]

			ScalInc(test.alpha, x, uintptr(n), uintptr(incX))

			for i := range test.want {
				if !same(x[i*incX], test.want[i]) {
					t.Errorf(msgVal, prefix, i, x[i*incX], test.want[i])
				}
			}
			checkValidIncGuard(t, xg, xGdVal, incX, gdLn)
		}
	}
}

func TestScalIncTo(t *testing.T) {
	const xGdVal, dstGdVal = -1, 0.5
	gdLn := 4
	rng := rand.New(rand.NewSource(42))
	for i, test := range scalTests {
		n := len(test.x)
		for _, inc := range newIncSet(1, 2, 3, 4, 7, 10) {
			prefix := fmt.Sprintf("test %v (x:%v dst:%v)", i, inc.x, inc.y)
			xg := guardIncVector(test.x, xGdVal, inc.x, gdLn)
			dg := guardIncVector(randSlice(n, 1, rng), dstGdVal, inc.y, gdLn)
			x, dst := xg[gdLn:len(xg)-gdLn], dg[gdLn:len(dg)-gdLn]

			ScalIncTo(dst, uintptr(inc.y), test.alpha, x, uintptr(n), uintptr(inc.x))

			for i := range test.want {
				if !same(dst[i*inc.y], test.want[i]) {
					t.Errorf(msgVal, prefix, i, dst[i*inc.y], test.want[i])
				}
			}
			checkValidIncGuard(t, xg, xGdVal, inc.x, gdLn)
			checkValidIncGuard(t, dg, dstGdVal, inc.y, gdLn)
			if !equalStrided(test.x, x, inc.x) {
				t.Errorf("%v: modified read-only x argument", prefix)
			}

		}
	}
}

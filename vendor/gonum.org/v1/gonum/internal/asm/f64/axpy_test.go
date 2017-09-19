// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package f64

import (
	"fmt"
	"testing"
)

const (
	msgVal   = "%v: unexpected value at %v Got: %v Expected: %v"
	msgGuard = "%v: Guard violated in %s vector %v %v"
)

var axpyTests = []struct {
	alpha   float64
	x       []float64
	y       []float64
	want    []float64
	wantRev []float64 // Result when x is traversed in reverse direction.
}{
	{
		alpha:   0,
		x:       []float64{},
		y:       []float64{},
		want:    []float64{},
		wantRev: []float64{},
	},
	{
		alpha:   0,
		x:       []float64{2},
		y:       []float64{-3},
		want:    []float64{-3},
		wantRev: []float64{-3},
	},
	{
		alpha:   1,
		x:       []float64{2},
		y:       []float64{-3},
		want:    []float64{-1},
		wantRev: []float64{-1},
	},
	{
		alpha:   3,
		x:       []float64{2},
		y:       []float64{-3},
		want:    []float64{3},
		wantRev: []float64{3},
	},
	{
		alpha:   -3,
		x:       []float64{2},
		y:       []float64{-3},
		want:    []float64{-9},
		wantRev: []float64{-9},
	},
	{
		alpha:   1,
		x:       []float64{1, 5},
		y:       []float64{2, -3},
		want:    []float64{3, 2},
		wantRev: []float64{7, -2},
	},
	{
		alpha:   1,
		x:       []float64{2, 3, 4},
		y:       []float64{-3, -2, -1},
		want:    []float64{-1, 1, 3},
		wantRev: []float64{1, 1, 1},
	},
	{
		alpha:   0,
		x:       []float64{0, 0, 1, 1, 2, -3, -4},
		y:       []float64{0, 1, 0, 3, -4, 5, -6},
		want:    []float64{0, 1, 0, 3, -4, 5, -6},
		wantRev: []float64{0, 1, 0, 3, -4, 5, -6},
	},
	{
		alpha:   1,
		x:       []float64{0, 0, 1, 1, 2, -3, -4},
		y:       []float64{0, 1, 0, 3, -4, 5, -6},
		want:    []float64{0, 1, 1, 4, -2, 2, -10},
		wantRev: []float64{-4, -2, 2, 4, -3, 5, -6},
	},
	{
		alpha:   3,
		x:       []float64{0, 0, 1, 1, 2, -3, -4},
		y:       []float64{0, 1, 0, 3, -4, 5, -6},
		want:    []float64{0, 1, 3, 6, 2, -4, -18},
		wantRev: []float64{-12, -8, 6, 6, -1, 5, -6},
	},
	{
		alpha:   -3,
		x:       []float64{0, 0, 1, 1, 2, -3, -4, 0, 0, 1, 1, 2, -3, -4},
		y:       []float64{0, 1, 0, 3, -4, 5, -6, 0, 1, 0, 3, -4, 5, -6},
		want:    []float64{0, 1, -3, 0, -10, 14, 6, 0, 1, -3, 0, -10, 14, 6},
		wantRev: []float64{12, 10, -6, 0, -7, 5, -6, 12, 10, -6, 0, -7, 5, -6},
	},
	{
		alpha:   -5,
		x:       []float64{0, 0, 1, 1, 2, -3, -4, 5, 1, 2, -3, -4, 5},
		y:       []float64{0, 1, 0, 3, -4, 5, -6, 7, 3, -4, 5, -6, 7},
		want:    []float64{0, 1, -5, -2, -14, 20, 14, -18, -2, -14, 20, 14, -18},
		wantRev: []float64{-25, 21, 15, -7, -9, -20, 14, 22, -7, -9, 0, -6, 7},
	},
}

func TestAxpyUnitary(t *testing.T) {
	const xGdVal, yGdVal = -1, 0.5
	for i, test := range axpyTests {
		for _, align := range align2 {
			prefix := fmt.Sprintf("Test %v (x:%v y:%v)", i, align.x, align.y)
			xgLn, ygLn := 4+align.x, 4+align.y
			xg, yg := guardVector(test.x, xGdVal, xgLn), guardVector(test.y, yGdVal, ygLn)
			x, y := xg[xgLn:len(xg)-xgLn], yg[ygLn:len(yg)-ygLn]
			AxpyUnitary(test.alpha, x, y)
			for i := range test.want {
				if !same(y[i], test.want[i]) {
					t.Errorf(msgVal, prefix, i, y[i], test.want[i])
				}
			}
			if !isValidGuard(xg, xGdVal, xgLn) {
				t.Errorf(msgGuard, prefix, "x", xg[:xgLn], xg[len(xg)-xgLn:])
			}
			if !isValidGuard(yg, yGdVal, ygLn) {
				t.Errorf(msgGuard, prefix, "y", yg[:ygLn], yg[len(yg)-ygLn:])
			}
			if !equalStrided(test.x, x, 1) {
				t.Errorf("%v: modified read-only x argument", prefix)
			}
		}
	}
}

func TestAxpyUnitaryTo(t *testing.T) {
	const dstGdVal, xGdVal, yGdVal = 1, -1, 0.5
	for i, test := range axpyTests {
		for _, align := range align3 {
			prefix := fmt.Sprintf("Test %v (x:%v y:%v dst:%v)", i, align.x, align.y, align.dst)

			dgLn, xgLn, ygLn := 4+align.dst, 4+align.x, 4+align.y
			dstOrig := make([]float64, len(test.x))
			xg, yg := guardVector(test.x, xGdVal, xgLn), guardVector(test.y, yGdVal, ygLn)
			dstg := guardVector(dstOrig, dstGdVal, dgLn)
			x, y := xg[xgLn:len(xg)-xgLn], yg[ygLn:len(yg)-ygLn]
			dst := dstg[dgLn : len(dstg)-dgLn]

			AxpyUnitaryTo(dst, test.alpha, x, y)
			for i := range test.want {
				if !same(dst[i], test.want[i]) {
					t.Errorf(msgVal, prefix, i, dst[i], test.want[i])
				}
			}
			if !isValidGuard(xg, xGdVal, xgLn) {
				t.Errorf(msgGuard, prefix, "x", xg[:xgLn], xg[len(xg)-xgLn:])
			}
			if !isValidGuard(yg, yGdVal, ygLn) {
				t.Errorf(msgGuard, prefix, "y", yg[:ygLn], yg[len(yg)-ygLn:])
			}
			if !isValidGuard(dstg, dstGdVal, dgLn) {
				t.Errorf(msgGuard, prefix, "dst", dstg[:dgLn], dstg[len(dstg)-dgLn:])
			}
			if !equalStrided(test.x, x, 1) {
				t.Errorf("%v: modified read-only x argument", prefix)
			}
			if !equalStrided(test.y, y, 1) {
				t.Errorf("%v: modified read-only y argument", prefix)
			}
		}
	}
}

func TestAxpyInc(t *testing.T) {
	const xGdVal, yGdVal = -1, 0.5
	gdLn := 4
	for i, test := range axpyTests {
		n := len(test.x)
		for _, inc := range newIncSet(-7, -4, -3, -2, -1, 1, 2, 3, 4, 7) {
			var ix, iy int
			if inc.x < 0 {
				ix = (-n + 1) * inc.x
			}
			if inc.y < 0 {
				iy = (-n + 1) * inc.y
			}
			prefix := fmt.Sprintf("test %v, inc.x = %v, inc.y = %v", i, inc.x, inc.y)
			xg := guardIncVector(test.x, xGdVal, inc.x, gdLn)
			yg := guardIncVector(test.y, yGdVal, inc.y, gdLn)
			x, y := xg[gdLn:len(xg)-gdLn], yg[gdLn:len(yg)-gdLn]

			AxpyInc(test.alpha, x, y, uintptr(n),
				uintptr(inc.x), uintptr(inc.y), uintptr(ix), uintptr(iy))

			want := test.want
			if inc.x*inc.y < 0 {
				want = test.wantRev
			}
			if inc.y < 0 {
				inc.y = -inc.y
			}
			for i := range want {
				if !same(y[i*inc.y], want[i]) {
					t.Errorf(msgVal, prefix, i, y[iy+i*inc.y], want[i])
				}
			}
			if !equalStrided(test.x, x, inc.x) {
				t.Errorf("%v: modified read-only x argument", prefix)
			}
			checkValidIncGuard(t, xg, xGdVal, inc.x, gdLn)
			checkValidIncGuard(t, yg, yGdVal, inc.y, gdLn)
		}
	}
}

func TestAxpyIncTo(t *testing.T) {
	const dstGdVal, xGdVal, yGdVal = 1, -1, 0.5
	var want []float64
	gdLn := 4
	for i, test := range axpyTests {
		n := len(test.x)
		for _, inc := range newIncToSet(-7, -4, -3, -2, -1, 1, 2, 3, 4, 7) {
			var ix, iy, idst uintptr
			if inc.x < 0 {
				ix = uintptr((-n + 1) * inc.x)
			}
			if inc.y < 0 {
				iy = uintptr((-n + 1) * inc.y)
			}
			if inc.dst < 0 {
				idst = uintptr((-n + 1) * inc.dst)
			}

			prefix := fmt.Sprintf("Test %v: (x: %v, y: %v, dst:%v)", i, inc.x, inc.y, inc.dst)
			dstOrig := make([]float64, len(test.want))
			xg := guardIncVector(test.x, xGdVal, inc.x, gdLn)
			yg := guardIncVector(test.y, yGdVal, inc.y, gdLn)
			dstg := guardIncVector(dstOrig, dstGdVal, inc.dst, gdLn)
			x, y := xg[gdLn:len(xg)-gdLn], yg[gdLn:len(yg)-gdLn]
			dst := dstg[gdLn : len(dstg)-gdLn]

			AxpyIncTo(dst, uintptr(inc.dst), idst,
				test.alpha, x, y, uintptr(n),
				uintptr(inc.x), uintptr(inc.y), ix, iy)
			want = test.want
			if inc.x*inc.y < 0 {
				want = test.wantRev
			}
			var iW, incW int = 0, 1
			if inc.y*inc.dst < 0 {
				iW, incW = len(want)-1, -1
			}
			if inc.dst < 0 {
				inc.dst = -inc.dst
			}
			for i := range want {
				if !same(dst[i*inc.dst], want[iW+i*incW]) {
					t.Errorf(msgVal, prefix, i, dst[i*inc.dst], want[iW+i*incW])
				}
			}

			checkValidIncGuard(t, xg, xGdVal, inc.x, gdLn)
			checkValidIncGuard(t, yg, yGdVal, inc.y, gdLn)
			checkValidIncGuard(t, dstg, dstGdVal, inc.dst, gdLn)
			if !equalStrided(test.x, x, inc.x) {
				t.Errorf("%v: modified read-only x argument", prefix)
			}
			if !equalStrided(test.y, y, inc.y) {
				t.Errorf("%v: modified read-only y argument", prefix)
			}
		}
	}
}

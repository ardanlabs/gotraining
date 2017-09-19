// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package f32

import (
	"fmt"
	"math"
	"testing"
)

const (
	msgRes   = "%v: unexpected result Got: %v Expected: %v"
	msgGuard = "%v: Guard violated in %s vector %v %v"
)

var dotTests = []struct {
	x, y     []float32
	sWant    float32 // single-precision
	dWant    float64 // double-precision
	sWantRev float32 // single-precision increment
	dWantRev float64 // double-precision increment
	n        int
	ix, iy   int
}{
	{ // 0
		x:     []float32{},
		y:     []float32{},
		n:     0,
		sWant: 0, dWant: 0,
		sWantRev: 0, dWantRev: 0,
		ix: 0, iy: 0,
	},
	{ // 1
		x:     []float32{0},
		y:     []float32{0},
		n:     1,
		sWant: 0, dWant: 0,
		sWantRev: 0, dWantRev: 0,
		ix: 0, iy: 0,
	},
	{ // 2
		x:     []float32{1},
		y:     []float32{1},
		n:     1,
		sWant: 1, dWant: 1,
		sWantRev: 1, dWantRev: 1,
		ix: 0, iy: 0,
	},
	{ // 3
		x:     []float32{1, 2, 3, 4, 5, 6, 7, 8},
		y:     []float32{2, 2, 2, 2, 2, 2, 2, 2},
		n:     8,
		sWant: 72, dWant: 72,
		sWantRev: 72, dWantRev: 72,
		ix: 1, iy: 1,
	},
	{ // 4
		x:     []float32{math.MaxFloat32},
		y:     []float32{2},
		n:     1,
		sWant: inf, dWant: 2 * float64(math.MaxFloat32),
		sWantRev: inf, dWantRev: 2 * float64(math.MaxFloat32),
		ix: 0, iy: 0,
	},
	{ // 5
		x:     []float32{1, 1, 2, 2, 1, 1, 2, 2, 1, 1, 2, 2, 1, 1, 2, 2, 1, 1, 2, 2},
		y:     []float32{3, 3, 2, 2, 3, 3, 2, 2, 3, 3, 2, 2, 3, 3, 2, 2, 3, 3, 2, 2},
		n:     20,
		sWant: 70, dWant: 70,
		sWantRev: 80, dWantRev: 80,
		ix: 0, iy: 0,
	},
}

func TestDotUnitary(t *testing.T) {
	const xGdVal, yGdVal = 0.5, 0.25
	for i, test := range dotTests {
		for _, align := range align2 {
			prefix := fmt.Sprintf("Test %v (x:%v y:%v)", i, align.x, align.y)
			xgLn, ygLn := 8+align.x, 8+align.y
			xg, yg := guardVector(test.x, xGdVal, xgLn), guardVector(test.y, yGdVal, ygLn)
			x, y := xg[xgLn:len(xg)-xgLn], yg[ygLn:len(yg)-ygLn]
			res := DotUnitary(x, y)
			if !same(res, test.sWant) {
				t.Errorf(msgRes, prefix, res, test.sWant)
			}
			if !isValidGuard(xg, xGdVal, xgLn) {
				t.Errorf(msgGuard, prefix, "x", xg[:xgLn], xg[len(xg)-xgLn:])
			}
			if !isValidGuard(yg, yGdVal, ygLn) {
				t.Errorf(msgGuard, prefix, "y", yg[:ygLn], yg[len(yg)-ygLn:])
			}
		}
	}
}

func TestDotInc(t *testing.T) {
	const xGdVal, yGdVal, gdLn = 0.5, 0.25, 8
	for i, test := range dotTests {
		for _, inc := range newIncSet(1, 2, 3, 4, 7, 10, -1, -2, -5, -10) {
			xg, yg := guardIncVector(test.x, xGdVal, inc.x, gdLn), guardIncVector(test.y, yGdVal, inc.y, gdLn)
			x, y := xg[gdLn:len(xg)-gdLn], yg[gdLn:len(yg)-gdLn]
			want := test.sWant
			var ix, iy int
			if inc.x < 0 {
				ix = -inc.x * (test.n - 1)
			}
			if inc.y < 0 {
				iy = -inc.y * (test.n - 1)
			}
			if inc.x*inc.y < 0 {
				want = test.sWantRev
			}
			prefix := fmt.Sprintf("Test %v (x:%v y:%v) (ix:%v iy:%v)", i, inc.x, inc.y, ix, iy)
			res := DotInc(x, y, uintptr(test.n), uintptr(inc.x), uintptr(inc.y), uintptr(ix), uintptr(iy))
			if !same(res, want) {
				t.Errorf(msgRes, prefix, res, want)
			}
			checkValidIncGuard(t, xg, xGdVal, inc.x, gdLn)
			checkValidIncGuard(t, yg, yGdVal, inc.y, gdLn)
		}
	}
}

func TestDdotUnitary(t *testing.T) {
	const xGdVal, yGdVal = 0.5, 0.25
	for i, test := range dotTests {
		for _, align := range align2 {
			prefix := fmt.Sprintf("Test %v (x:%v y:%v)", i, align.x, align.y)
			xgLn, ygLn := 8+align.x, 8+align.y
			xg, yg := guardVector(test.x, xGdVal, xgLn), guardVector(test.y, yGdVal, ygLn)
			x, y := xg[xgLn:len(xg)-xgLn], yg[ygLn:len(yg)-ygLn]
			res := DdotUnitary(x, y)
			if !same64(res, test.dWant) {
				t.Errorf(msgRes, prefix, res, test.dWant)
			}
			if !isValidGuard(xg, xGdVal, xgLn) {
				t.Errorf(msgGuard, prefix, "x", xg[:xgLn], xg[len(xg)-xgLn:])
			}
			if !isValidGuard(yg, yGdVal, ygLn) {
				t.Errorf(msgGuard, prefix, "y", yg[:ygLn], yg[len(yg)-ygLn:])
			}
		}
	}
}

func TestDdotInc(t *testing.T) {
	const xGdVal, yGdVal, gdLn = 0.5, 0.25, 8
	for i, test := range dotTests {
		for _, inc := range newIncSet(1, 2, 3, 4, 7, 10, -1, -2, -5, -10) {
			prefix := fmt.Sprintf("Test %v (x:%v y:%v)", i, inc.x, inc.y)
			xg, yg := guardIncVector(test.x, xGdVal, inc.x, gdLn), guardIncVector(test.y, yGdVal, inc.y, gdLn)
			x, y := xg[gdLn:len(xg)-gdLn], yg[gdLn:len(yg)-gdLn]
			want := test.dWant
			var ix, iy int
			if inc.x < 0 {
				ix = -inc.x * (test.n - 1)
			}
			if inc.y < 0 {
				iy = -inc.y * (test.n - 1)
			}
			if inc.x*inc.y < 0 {
				want = test.dWantRev
			}
			res := DdotInc(x, y, uintptr(test.n), uintptr(inc.x), uintptr(inc.y), uintptr(ix), uintptr(iy))
			if !same64(res, want) {
				t.Errorf(msgRes, prefix, res, want)
			}
			checkValidIncGuard(t, xg, xGdVal, inc.x, gdLn)
			checkValidIncGuard(t, yg, yGdVal, inc.y, gdLn)
		}
	}
}

// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package c128

import (
	"fmt"
	"testing"
)

var tests = []struct {
	incX, incY, incDst int
	ix, iy, idst       uintptr
	a                  complex128
	dst, x, y          []complex128
	ex                 []complex128
}{
	{incX: 2, incY: 2, incDst: 3, ix: 0, iy: 0, idst: 0,
		a:   1 + 1i,
		dst: []complex128{5},
		x:   []complex128{1},
		y:   []complex128{1i},
		ex:  []complex128{1 + 2i}},
	{incX: 2, incY: 2, incDst: 3, ix: 0, iy: 0, idst: 0,
		a:   1 + 2i,
		dst: []complex128{0, 0, 0},
		x:   []complex128{0, 0, 0},
		y:   []complex128{1, 1, 1},
		ex:  []complex128{1, 1, 1}},
	{incX: 2, incY: 2, incDst: 3, ix: 0, iy: 0, idst: 0,
		a:   1 + 2i,
		dst: []complex128{0, 0, 0},
		x:   []complex128{0, 0},
		y:   []complex128{1, 1, 1},
		ex:  []complex128{1, 1}},
	{incX: 2, incY: 2, incDst: 3, ix: 0, iy: 0, idst: 0,
		a:   1 + 2i,
		dst: []complex128{1i, 1i, 1i},
		x:   []complex128{1i, 1i, 1i},
		y:   []complex128{1, 2, 1},
		ex:  []complex128{-1 + 1i, 1i, -1 + 1i}},
	{incX: 2, incY: 2, incDst: 3, ix: 0, iy: 0, idst: 0,
		a:   -1i,
		dst: []complex128{1i, 1i, 1i},
		x:   []complex128{1i, 1i, 1i},
		y:   []complex128{1, 2, 1},
		ex:  []complex128{2, 3, 2}},
	{incX: 2, incY: 2, incDst: 3, ix: 0, iy: 0, idst: 0,
		a:   -1i,
		dst: []complex128{1i, 1i, 1i},
		x:   []complex128{1i, 1i, 1i, 1i, 1i}[1:4],
		y:   []complex128{1, 1, 2, 1, 1}[1:4],
		ex:  []complex128{2, 3, 2}},
	{incX: 2, incY: 4, incDst: 3, ix: 0, iy: 0, idst: 0,
		a:   -2,
		dst: []complex128{1i, 1i, 1i, 1i, 1i},
		x:   []complex128{2 + 1i, 2 + 1i, 2 + 1i, 2 + 1i, 2 + 1i},
		y:   []complex128{1, 1, 2, 1, 1},
		ex:  []complex128{-3 - 2i, -3 - 2i, -2 - 2i, -3 - 2i, -3 - 2i}},
	// Run big test twice, once aligned once unaligned.
	{incX: 2, incY: 2, incDst: 3, ix: 0, iy: 0, idst: 0,
		a:   1 - 1i,
		dst: make([]complex128, 10),
		x:   []complex128{1i, 1i, 1i, 1i, 1i, 1i, 1i, 1i, 1i, 1i},
		y:   []complex128{1, 1, 2, 1, 1, 1, 1, 2, 1, 1},
		ex:  []complex128{2 + 1i, 2 + 1i, 3 + 1i, 2 + 1i, 2 + 1i, 2 + 1i, 2 + 1i, 3 + 1i, 2 + 1i, 2 + 1i}},
	{incX: 2, incY: 2, incDst: 3, ix: 0, iy: 0, idst: 0,
		a:   1 - 1i,
		dst: make([]complex128, 10),
		x:   []complex128{1i, 1i, 1i, 1i, 1i, 1i, 1i, 1i, 1i, 1i},
		y:   []complex128{1, 1, 2, 1, 1, 1, 1, 2, 1, 1},
		ex:  []complex128{2 + 1i, 2 + 1i, 3 + 1i, 2 + 1i, 2 + 1i, 2 + 1i, 2 + 1i, 3 + 1i, 2 + 1i, 2 + 1i}},
	{incX: -2, incY: -2, incDst: -3, ix: 18, iy: 18, idst: 27,
		a:   1 - 1i,
		dst: make([]complex128, 10),
		x:   []complex128{1i, 1i, 1i, 1i, 1i, 1i, 1i, 1i, 1i, 1i},
		y:   []complex128{1, 1, 2, 1, 1, 1, 1, 2, 1, 1},
		ex:  []complex128{2 + 1i, 2 + 1i, 3 + 1i, 2 + 1i, 2 + 1i, 2 + 1i, 2 + 1i, 3 + 1i, 2 + 1i, 2 + 1i}},
	{incX: -2, incY: 2, incDst: -3, ix: 18, iy: 0, idst: 27,
		a:   1 - 1i,
		dst: make([]complex128, 10),
		x:   []complex128{1i, 1i, 1i, 1i, 1i, 1i, 1i, 1i, 1i, 1i},
		y:   []complex128{1, 1, 2, 1, 1, 1, 1, 2, 1, 1},
		ex:  []complex128{2 + 1i, 2 + 1i, 3 + 1i, 2 + 1i, 2 + 1i, 2 + 1i, 2 + 1i, 3 + 1i, 2 + 1i, 2 + 1i}},
}

func TestAxpyUnitary(t *testing.T) {
	const xGdVal, yGdVal = 1, 1
	for cas, test := range tests {
		xgLn, ygLn := 4+cas%2, 4+cas%3
		test.x, test.y = guardVector(test.x, xGdVal, xgLn), guardVector(test.y, yGdVal, ygLn)
		x, y := test.x[xgLn:len(test.x)-xgLn], test.y[ygLn:len(test.y)-ygLn]
		AxpyUnitary(test.a, x, y)
		for i := range test.ex {
			if y[i] != test.ex[i] {
				t.Errorf("Test %d Unexpected result at %d Got: %v Expected: %v", cas, i, y[i], test.ex[i])
			}
		}
		if !isValidGuard(test.x, xGdVal, xgLn) {
			t.Errorf("Test %d Guard violated in x vector %v %v", cas, test.x[:xgLn], test.x[len(test.x)-xgLn:])
		}
		if !isValidGuard(test.y, yGdVal, ygLn) {
			t.Errorf("Test %d Guard violated in y vector %v %v", cas, test.y[:ygLn], test.y[len(test.y)-ygLn:])
		}
	}
}

func TestAxpyUnitaryTo(t *testing.T) {
	const xGdVal, yGdVal, dstGdVal = 1, 1, 0
	for cas, test := range tests {
		xgLn, ygLn := 4+cas%2, 4+cas%3
		test.x, test.y = guardVector(test.x, xGdVal, xgLn), guardVector(test.y, yGdVal, ygLn)
		test.dst = guardVector(test.dst, dstGdVal, xgLn)
		x, y := test.x[xgLn:len(test.x)-xgLn], test.y[ygLn:len(test.y)-ygLn]
		dst := test.dst[xgLn : len(test.dst)-xgLn]
		AxpyUnitaryTo(dst, test.a, x, y)
		for i := range test.ex {
			if dst[i] != test.ex[i] {
				t.Errorf("Test %d Unexpected result at %d Got: %v Expected: %v", cas, i, dst[i], test.ex[i])
			}
		}
		if !isValidGuard(test.x, xGdVal, xgLn) {
			t.Errorf("Test %d Guard violated in x vector %v %v", cas, test.x[:xgLn], test.x[len(test.x)-xgLn:])
		}
		if !isValidGuard(test.y, yGdVal, ygLn) {
			t.Errorf("Test %d Guard violated in y vector %v %v", cas, test.y[:ygLn], test.y[len(test.y)-ygLn:])
		}
		if !isValidGuard(test.dst, dstGdVal, xgLn) {
			t.Errorf("Test %d Guard violated in dst vector %v %v", cas, test.dst[:xgLn], test.dst[len(test.dst)-xgLn:])
		}

	}
}

func TestAxpyInc(t *testing.T) {
	const xGdVal, yGdVal = 1, 1
	for cas, test := range tests {
		xgLn, ygLn := 4+cas%2, 4+cas%3
		test.x, test.y = guardIncVector(test.x, xGdVal, test.incX, xgLn), guardIncVector(test.y, yGdVal, test.incY, ygLn)
		x, y := test.x[xgLn:len(test.x)-xgLn], test.y[ygLn:len(test.y)-ygLn]
		AxpyInc(test.a, x, y, uintptr(len(test.ex)), uintptr(test.incX), uintptr(test.incY), test.ix, test.iy)
		for i := range test.ex {
			if y[int(test.iy)+i*int(test.incY)] != test.ex[i] {
				t.Errorf("Test %d Unexpected result at %d Got: %v Expected: %v", cas, i, y[i*int(test.incY)], test.ex[i])
			}
		}
		checkValidIncGuard(t, test.x, xGdVal, test.incX, xgLn)
		checkValidIncGuard(t, test.y, yGdVal, test.incY, ygLn)
	}
}

func TestAxpyIncTo(t *testing.T) {
	const xGdVal, yGdVal, dstGdVal = 1, 1, 0
	for cas, test := range tests {
		xgLn, ygLn := 4+cas%2, 4+cas%3
		test.x, test.y = guardIncVector(test.x, xGdVal, test.incX, xgLn), guardIncVector(test.y, yGdVal, test.incY, ygLn)
		test.dst = guardIncVector(test.dst, dstGdVal, test.incDst, xgLn)
		x, y := test.x[xgLn:len(test.x)-xgLn], test.y[ygLn:len(test.y)-ygLn]
		dst := test.dst[xgLn : len(test.dst)-xgLn]
		AxpyIncTo(dst, uintptr(test.incDst), test.idst, test.a, x, y, uintptr(len(test.ex)), uintptr(test.incX), uintptr(test.incY), test.ix, test.iy)
		for i := range test.ex {
			if dst[int(test.idst)+i*int(test.incDst)] != test.ex[i] {
				t.Errorf("Test %d Unexpected result at %d Got: %v Expected: %v", cas, i, dst[i*int(test.incDst)], test.ex[i])
			}
		}
		checkValidIncGuard(t, test.x, xGdVal, test.incX, xgLn)
		checkValidIncGuard(t, test.y, yGdVal, test.incY, ygLn)
		checkValidIncGuard(t, test.dst, dstGdVal, test.incDst, xgLn)
	}
}

var dscalTests = []struct {
	alpha float64
	x     []complex128
	want  []complex128
}{
	{
		alpha: 0,
		x:     []complex128{},
		want:  []complex128{},
	},
	{
		alpha: 1,
		x:     []complex128{1 + 2i},
		want:  []complex128{1 + 2i},
	},
	{
		alpha: 2,
		x:     []complex128{1 + 2i},
		want:  []complex128{2 + 4i},
	},
	{
		alpha: 2,
		x:     []complex128{1 + 2i, 3 + 5i, 6 + 11i, 12 - 23i},
		want:  []complex128{2 + 4i, 6 + 10i, 12 + 22i, 24 - 46i},
	},
	{
		alpha: 3,
		x:     []complex128{1 + 2i, 5 + 4i, 3 + 6i, 8 + 12i, -3 - 2i, -5 + 5i},
		want:  []complex128{3 + 6i, 15 + 12i, 9 + 18i, 24 + 36i, -9 - 6i, -15 + 15i},
	},
	{
		alpha: 5,
		x:     []complex128{1 + 2i, 5 + 4i, 3 + 6i, 8 + 12i, -3 - 2i, -5 + 5i, 1 + 2i, 5 + 4i, 3 + 6i, 8 + 12i, -3 - 2i, -5 + 5i},
		want:  []complex128{5 + 10i, 25 + 20i, 15 + 30i, 40 + 60i, -15 - 10i, -25 + 25i, 5 + 10i, 25 + 20i, 15 + 30i, 40 + 60i, -15 - 10i, -25 + 25i},
	},
}

func TestDscalUnitary(t *testing.T) {
	const xGdVal = -0.5
	for i, test := range dscalTests {
		for _, align := range align1 {
			prefix := fmt.Sprintf("Test %v (x:%v)", i, align)
			xgLn := 4 + align
			xg := guardVector(test.x, xGdVal, xgLn)
			x := xg[xgLn : len(xg)-xgLn]

			DscalUnitary(test.alpha, x)

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

func TestDscalInc(t *testing.T) {
	const xGdVal = -0.5
	gdLn := 4
	for i, test := range dscalTests {
		n := len(test.x)
		for _, incX := range []int{1, 2, 3, 4, 7, 10} {
			prefix := fmt.Sprintf("Test %v (x:%v)", i, incX)
			xg := guardIncVector(test.x, xGdVal, incX, gdLn)
			x := xg[gdLn : len(xg)-gdLn]

			DscalInc(test.alpha, x, uintptr(n), uintptr(incX))

			for i := range test.want {
				if !same(x[i*incX], test.want[i]) {
					t.Errorf(msgVal, prefix, i, x[i*incX], test.want[i])
				}
			}
			checkValidIncGuard(t, xg, xGdVal, incX, gdLn)
		}
	}
}

var scalTests = []struct {
	alpha complex128
	x     []complex128
	want  []complex128
}{
	{
		alpha: 0,
		x:     []complex128{},
		want:  []complex128{},
	},
	{
		alpha: 1 + 1i,
		x:     []complex128{1 + 2i},
		want:  []complex128{-1 + 3i},
	},
	{
		alpha: 2 + 3i,
		x:     []complex128{1 + 2i},
		want:  []complex128{-4 + 7i},
	},
	{
		alpha: 2 - 4i,
		x:     []complex128{1 + 2i},
		want:  []complex128{10},
	},
	{
		alpha: 2 + 8i,
		x:     []complex128{1 + 2i, 5 + 4i, 3 + 6i, 8 + 12i, -3 - 2i, -5 + 5i},
		want:  []complex128{-14 + 12i, -22 + 48i, -42 + 36i, -80 + 88i, 10 - 28i, -50 - 30i},
	},
	{
		alpha: 5 - 10i,
		x:     []complex128{1 + 2i, 5 + 4i, 3 + 6i, 8 + 12i, -3 - 2i, -5 + 5i, 1 + 2i, 5 + 4i, 3 + 6i, 8 + 12i, -3 - 2i, -5 + 5i},
		want:  []complex128{25, 65 - 30i, 75, 160 - 20i, -35 + 20i, 25 + 75i, 25, 65 - 30i, 75, 160 - 20i, -35 + 20i, 25 + 75i},
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

func TestScalInc(t *testing.T) {
	const xGdVal = -0.5
	gdLn := 4
	for i, test := range scalTests {
		n := len(test.x)
		for _, inc := range []int{1, 2, 3, 4, 7, 10} {
			prefix := fmt.Sprintf("Test %v (x:%v)", i, inc)
			xg := guardIncVector(test.x, xGdVal, inc, gdLn)
			x := xg[gdLn : len(xg)-gdLn]

			ScalInc(test.alpha, x, uintptr(n), uintptr(inc))

			for i := range test.want {
				if !same(x[i*inc], test.want[i]) {
					t.Errorf(msgVal, prefix, i, x[i*inc], test.want[i])
				}
			}
			checkValidIncGuard(t, xg, xGdVal, inc, gdLn)
		}
	}
}

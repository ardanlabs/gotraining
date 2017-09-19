// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package f32

import (
	"math"
	"testing"
)

var (
	nan = float32(math.NaN())
	inf = float32(math.Inf(1))
)

var tests = []struct {
	incX, incY, incDst uintptr
	ix, iy, idst       uintptr
	a                  float32
	dst, x, y          []float32
	ex                 []float32
}{
	{incX: 2, incY: 2, incDst: 3, ix: 0, iy: 0, idst: 0,
		a:   3,
		dst: []float32{5},
		x:   []float32{2},
		y:   []float32{1},
		ex:  []float32{7}},
	{incX: 2, incY: 2, incDst: 3, ix: 0, iy: 0, idst: 0,
		a:   5,
		dst: []float32{0, 0, 0},
		x:   []float32{0, 0, 0},
		y:   []float32{1, 1, 1},
		ex:  []float32{1, 1, 1}},
	{incX: 2, incY: 2, incDst: 3, ix: 0, iy: 0, idst: 0,
		a:   5,
		dst: []float32{0, 0, 0},
		x:   []float32{0, 0},
		y:   []float32{1, 1, 1},
		ex:  []float32{1, 1}},
	{incX: 2, incY: 2, incDst: 3, ix: 0, iy: 0, idst: 0,
		a:   -1,
		dst: []float32{-1, -1, -1},
		x:   []float32{1, 1, 1},
		y:   []float32{1, 2, 1},
		ex:  []float32{0, 1, 0}},
	{incX: 2, incY: 2, incDst: 3, ix: 0, iy: 0, idst: 0,
		a:   -1,
		dst: []float32{1, 1, 1},
		x:   []float32{1, 2, 1},
		y:   []float32{-1, -2, -1},
		ex:  []float32{-2, -4, -2}},
	{incX: 2, incY: 2, incDst: 3, ix: 0, iy: 0, idst: 0,
		a:   2.5,
		dst: []float32{1, 1, 1, 1, 1},
		x:   []float32{1, 2, 3, 2, 1},
		y:   []float32{0, 0, 0, 0, 0},
		ex:  []float32{2.5, 5, 7.5, 5, 2.5}},
	{incX: 2, incY: 2, incDst: 3, ix: 0, iy: 0, idst: 0, // Run big test twice, once aligned once unaligned.
		a:   16.5,
		dst: make([]float32, 20),
		x:   []float32{.5, .5, .5, .5, .5, .5, .5, .5, .5, .5, .5, .5, .5, .5, .5, .5, .5, .5, .5, .5},
		y:   []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		ex:  []float32{9.25, 10.25, 11.25, 12.25, 13.25, 14.25, 15.25, 16.25, 17.25, 18.25, 9.25, 10.25, 11.25, 12.25, 13.25, 14.25, 15.25, 16.25, 17.25, 18.25}},
	{incX: 2, incY: 2, incDst: 3, ix: 0, iy: 0, idst: 0,
		a:   16.5,
		dst: make([]float32, 10),
		x:   []float32{.5, .5, .5, .5, .5, .5, .5, .5, .5, .5},
		y:   []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		ex:  []float32{9.25, 10.25, 11.25, 12.25, 13.25, 14.25, 15.25, 16.25, 17.25, 18.25}},
}

func guardVector(vec []float32, guard_val float32, guard_len int) (guarded []float32) {
	guarded = make([]float32, len(vec)+guard_len*2)
	copy(guarded[guard_len:], vec)
	for i := 0; i < guard_len; i++ {
		guarded[i] = guard_val
		guarded[len(guarded)-1-i] = guard_val
	}
	return guarded
}

func isValidGuard(vec []float32, guard_val float32, guard_len int) bool {
	for i := 0; i < guard_len; i++ {
		if vec[i] != guard_val || vec[len(vec)-1-i] != guard_val {
			return false
		}
	}
	return true
}

func same(x, y float32) bool {
	a, b := float64(x), float64(y)
	return a == b || (math.IsNaN(a) && math.IsNaN(b))
}

func TestAxpyUnitary(t *testing.T) {
	var x_gd, y_gd float32 = 1, 1
	for cas, test := range tests {
		xg_ln, yg_ln := 4+cas%2, 4+cas%3
		test.x, test.y = guardVector(test.x, x_gd, xg_ln), guardVector(test.y, y_gd, yg_ln)
		x, y := test.x[xg_ln:len(test.x)-xg_ln], test.y[yg_ln:len(test.y)-yg_ln]
		AxpyUnitary(test.a, x, y)
		for i := range test.ex {
			if !same(y[i], test.ex[i]) {
				t.Errorf("Test %d Unexpected result at %d Got: %v Expected: %v", cas, i, y[i], test.ex[i])
			}
		}
		if !isValidGuard(test.x, x_gd, xg_ln) {
			t.Errorf("Test %d Guard violated in x vector %v %v", cas, test.x[:xg_ln], test.x[len(test.x)-xg_ln:])
		}
		if !isValidGuard(test.y, y_gd, yg_ln) {
			t.Errorf("Test %d Guard violated in y vector %v %v", cas, test.y[:yg_ln], test.y[len(test.y)-yg_ln:])
		}
	}
}

func TestAxpyUnitaryTo(t *testing.T) {
	var x_gd, y_gd, dst_gd float32 = 1, 1, 0
	for cas, test := range tests {
		xg_ln, yg_ln := 4+cas%2, 4+cas%3
		test.x, test.y = guardVector(test.x, x_gd, xg_ln), guardVector(test.y, y_gd, yg_ln)
		test.dst = guardVector(test.dst, dst_gd, xg_ln)
		x, y := test.x[xg_ln:len(test.x)-xg_ln], test.y[yg_ln:len(test.y)-yg_ln]
		dst := test.dst[xg_ln : len(test.dst)-xg_ln]
		AxpyUnitaryTo(dst, test.a, x, y)
		for i := range test.ex {
			if !same(test.ex[i], dst[i]) {
				t.Errorf("Test %d Unexpected result at %d Got: %v Expected: %v", cas, i, dst[i], test.ex[i])
			}
		}
		if !isValidGuard(test.x, x_gd, xg_ln) {
			t.Errorf("Test %d Guard violated in x vector %v %v", cas, test.x[:xg_ln], test.x[len(test.x)-xg_ln:])
		}
		if !isValidGuard(test.y, y_gd, yg_ln) {
			t.Errorf("Test %d Guard violated in y vector %v %v", cas, test.y[:yg_ln], test.y[len(test.y)-yg_ln:])
		}
		if !isValidGuard(test.dst, dst_gd, xg_ln) {
			t.Errorf("Test %d Guard violated in dst vector %v %v", cas, test.dst[:xg_ln], test.dst[len(test.dst)-xg_ln:])
		}
	}
}

func guardIncVector(vec []float32, guard_val float32, incV uintptr, guard_len int) (guarded []float32) {
	inc := int(incV)
	s_ln := len(vec) * (inc)
	guarded = make([]float32, s_ln+guard_len*2)
	for i, j := 0, 0; i < len(guarded); i++ {
		switch {
		case i < guard_len, i > guard_len+s_ln:
			guarded[i] = guard_val
		case (i-guard_len)%(inc) == 0 && j < len(vec):
			guarded[i] = vec[j]
			j++
		default:
			guarded[i] = guard_val
		}
	}
	return guarded
}

func checkValidIncGuard(t *testing.T, vec []float32, guard_val float32, incV uintptr, guard_len int) {
	inc := int(incV)
	s_ln := len(vec) - 2*guard_len
	for i := range vec {
		switch {
		case same(vec[i], guard_val):
			// Correct value
		case i < guard_len:
			t.Errorf("Front guard violated at %d %v", i, vec[:guard_len])
		case i > guard_len+s_ln:
			t.Errorf("Back guard violated at %d %v", i-guard_len-s_ln, vec[guard_len+s_ln:])
		case (i-guard_len)%inc == 0 && (i-guard_len)/inc < len(vec):
			// Ignore input values
		default:
			t.Errorf("Internal guard violated at %d %v", i-guard_len, vec[guard_len:guard_len+s_ln])
		}
	}
}

func TestAxpyInc(t *testing.T) {
	var x_gd, y_gd float32 = 1, 1
	for cas, test := range tests {
		xg_ln, yg_ln := 4+cas%2, 4+cas%3
		test.x, test.y = guardIncVector(test.x, x_gd, uintptr(test.incX), xg_ln), guardIncVector(test.y, y_gd, uintptr(test.incY), yg_ln)
		x, y := test.x[xg_ln:len(test.x)-xg_ln], test.y[yg_ln:len(test.y)-yg_ln]
		AxpyInc(test.a, x, y, uintptr(len(test.ex)), test.incX, test.incY, test.ix, test.iy)
		for i := range test.ex {
			if !same(y[i*int(test.incY)], test.ex[i]) {
				t.Errorf("Test %d Unexpected result at %d Got: %v Expected: %v", cas, i, y[i*int(test.incY)], test.ex[i])
			}
		}
		checkValidIncGuard(t, test.x, x_gd, uintptr(test.incX), xg_ln)
		checkValidIncGuard(t, test.y, y_gd, uintptr(test.incY), yg_ln)
	}
}

func TestAxpyIncTo(t *testing.T) {
	var x_gd, y_gd, dst_gd float32 = 1, 1, 0
	for cas, test := range tests {
		xg_ln, yg_ln := 4+cas%2, 4+cas%3
		test.x, test.y = guardIncVector(test.x, x_gd, uintptr(test.incX), xg_ln), guardIncVector(test.y, y_gd, uintptr(test.incY), yg_ln)
		test.dst = guardIncVector(test.dst, dst_gd, uintptr(test.incDst), xg_ln)
		x, y := test.x[xg_ln:len(test.x)-xg_ln], test.y[yg_ln:len(test.y)-yg_ln]
		dst := test.dst[xg_ln : len(test.dst)-xg_ln]
		AxpyIncTo(dst, test.incDst, test.idst, test.a, x, y, uintptr(len(test.ex)), test.incX, test.incY, test.ix, test.iy)
		for i := range test.ex {
			if !same(dst[i*int(test.incDst)], test.ex[i]) {
				t.Errorf("Test %d Unexpected result at %d Got: %v Expected: %v", cas, i, dst[i*int(test.incDst)], test.ex[i])
			}
		}
		checkValidIncGuard(t, test.x, x_gd, uintptr(test.incX), xg_ln)
		checkValidIncGuard(t, test.y, y_gd, uintptr(test.incY), yg_ln)
		checkValidIncGuard(t, test.dst, dst_gd, uintptr(test.incDst), xg_ln)
	}
}

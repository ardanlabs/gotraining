// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package f32

import "testing"

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

func TestAxpyUnitary(t *testing.T) {
	for j, v := range tests {
		gdLn := 4 + j%2
		v.x, v.y = guardVector(v.x, 1, gdLn), guardVector(v.y, 1, gdLn)
		x, y := v.x[gdLn:len(v.x)-gdLn], v.y[gdLn:len(v.y)-gdLn]
		AxpyUnitary(v.a, x, y)
		for i := range v.ex {
			if !same(y[i], v.ex[i]) {
				t.Error("Test", j, "Unexpected result at", i, "Got:", int(y[i]), "Expected:", v.ex[i])
			}
		}
		if !isValidGuard(v.x, 1, gdLn) {
			t.Error("Test", j, "Guard violated in x vector", v.x[:gdLn], v.x[len(v.x)-gdLn:])
		}
		if !isValidGuard(v.y, 1, gdLn) {
			t.Error("Test", j, "Guard violated in y vector", v.y[:gdLn], v.y[len(v.x)-gdLn:])
		}
	}
}

func TestAxpyUnitaryTo(t *testing.T) {
	for j, v := range tests {
		gdLn := 4 + j%2
		v.x, v.y = guardVector(v.x, 1, gdLn), guardVector(v.y, 1, gdLn)
		v.dst = guardVector(v.dst, 0, gdLn)
		x, y := v.x[gdLn:len(v.x)-gdLn], v.y[gdLn:len(v.y)-gdLn]
		dst := v.dst[gdLn : len(v.dst)-gdLn]
		AxpyUnitaryTo(dst, v.a, x, y)
		for i := range v.ex {
			if !same(v.ex[i], dst[i]) {
				t.Error("Test", j, "Unexpected result at", i, "Got:", dst[i], "Expected:", v.ex[i])
			}
		}
		if !isValidGuard(v.x, 1, gdLn) {
			t.Error("Test", j, "Guard violated in x vector", v.x[:gdLn], v.x[len(v.x)-gdLn:])
		}
		if !isValidGuard(v.y, 1, gdLn) {
			t.Error("Test", j, "Guard violated in y vector", v.y[:gdLn], v.y[len(v.x)-gdLn:])
		}
		if !isValidGuard(v.dst, 0, gdLn) {
			t.Error("Test", j, "Guard violated in x vector", v.x[:gdLn], v.x[len(v.x)-gdLn:])
		}
	}
}

func TestAxpyInc(t *testing.T) {
	for j, v := range tests {
		gdLn := 4 + j%2
		v.x, v.y = guardIncVector(v.x, 1, int(v.incX), gdLn), guardIncVector(v.y, 1, int(v.incY), gdLn)
		x, y := v.x[gdLn:len(v.x)-gdLn], v.y[gdLn:len(v.y)-gdLn]
		AxpyInc(v.a, x, y, uintptr(len(v.ex)), v.incX, v.incY, v.ix, v.iy)
		for i := range v.ex {
			if !same(y[i*int(v.incY)], v.ex[i]) {
				t.Error("Test", j, "Unexpected result at", i, "Got:", y[i*int(v.incY)], "Expected:", v.ex[i])
				t.Error("Result:", y)
				t.Error("Expect:", v.ex)
			}
		}
		checkValidIncGuard(t, v.x, 1, int(v.incX), gdLn)
		checkValidIncGuard(t, v.y, 1, int(v.incY), gdLn)
	}
}

func TestAxpyIncTo(t *testing.T) {
	for j, v := range tests {
		gdLn := 4 + j%2
		v.x, v.y = guardIncVector(v.x, 1, int(v.incX), gdLn), guardIncVector(v.y, 1, int(v.incY), gdLn)
		v.dst = guardIncVector(v.dst, 0, int(v.incDst), gdLn)
		x, y := v.x[gdLn:len(v.x)-gdLn], v.y[gdLn:len(v.y)-gdLn]
		dst := v.dst[gdLn : len(v.dst)-gdLn]
		AxpyIncTo(dst, v.incDst, v.idst, v.a, x, y, uintptr(len(v.ex)), v.incX, v.incY, v.ix, v.iy)
		for i := range v.ex {
			if !same(dst[i*int(v.incDst)], v.ex[i]) {
				t.Error("Test", j, "Unexpected result at", i, "Got:", dst[i*int(v.incDst)], "Expected:", v.ex[i])
				t.Error(v.dst)
				t.Error(v.ex)
			}
		}
		checkValidIncGuard(t, v.x, 1, int(v.incX), gdLn)
		checkValidIncGuard(t, v.y, 1, int(v.incY), gdLn)
		checkValidIncGuard(t, v.dst, 0, int(v.incDst), gdLn)
	}
}

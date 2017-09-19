// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat

import (
	"reflect"
	"testing"

	"gonum.org/v1/gonum/blas/blas64"
)

func TestNewBand(t *testing.T) {
	for i, test := range []struct {
		data   []float64
		r, c   int
		kl, ku int
		mat    *BandDense
		dense  *Dense
	}{
		{
			data: []float64{
				-1, 1, 2, 3,
				4, 5, 6, 7,
				8, 9, 10, 11,
				12, 13, 14, 15,
				16, 17, 18, -1,
				19, 20, -1, -1,
			},
			r: 6, c: 6,
			kl: 1, ku: 2,
			mat: &BandDense{
				mat: blas64.Band{
					Rows:   6,
					Cols:   6,
					KL:     1,
					KU:     2,
					Stride: 4,
					Data: []float64{
						-1, 1, 2, 3,
						4, 5, 6, 7,
						8, 9, 10, 11,
						12, 13, 14, 15,
						16, 17, 18, -1,
						19, 20, -1, -1,
					},
				},
			},
			dense: NewDense(6, 6, []float64{
				1, 2, 3, 0, 0, 0,
				4, 5, 6, 7, 0, 0,
				0, 8, 9, 10, 11, 0,
				0, 0, 12, 13, 14, 15,
				0, 0, 0, 16, 17, 18,
				0, 0, 0, 0, 19, 20,
			}),
		},
		{
			data: []float64{
				-1, 1, 2, 3,
				4, 5, 6, 7,
				8, 9, 10, 11,
				12, 13, 14, 15,
				16, 17, 18, -1,
				19, 20, -1, -1,
				21, -1, -1, -1,
			},
			r: 10, c: 6,
			kl: 1, ku: 2,
			mat: &BandDense{
				mat: blas64.Band{
					Rows:   10,
					Cols:   6,
					KL:     1,
					KU:     2,
					Stride: 4,
					Data: []float64{
						-1, 1, 2, 3,
						4, 5, 6, 7,
						8, 9, 10, 11,
						12, 13, 14, 15,
						16, 17, 18, -1,
						19, 20, -1, -1,
						21, -1, -1, -1,
					},
				},
			},
			dense: NewDense(10, 6, []float64{
				1, 2, 3, 0, 0, 0,
				4, 5, 6, 7, 0, 0,
				0, 8, 9, 10, 11, 0,
				0, 0, 12, 13, 14, 15,
				0, 0, 0, 16, 17, 18,
				0, 0, 0, 0, 19, 20,
				0, 0, 0, 0, 0, 21,
				0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0,
			}),
		},
		{
			data: []float64{
				-1, 1, 2, 3,
				4, 5, 6, 7,
				8, 9, 10, 11,
				12, 13, 14, 15,
				16, 17, 18, 19,
				20, 21, 22, 23,
			},
			r: 6, c: 10,
			kl: 1, ku: 2,
			mat: &BandDense{
				mat: blas64.Band{
					Rows:   6,
					Cols:   10,
					KL:     1,
					KU:     2,
					Stride: 4,
					Data: []float64{
						-1, 1, 2, 3,
						4, 5, 6, 7,
						8, 9, 10, 11,
						12, 13, 14, 15,
						16, 17, 18, 19,
						20, 21, 22, 23,
					},
				},
			},
			dense: NewDense(6, 10, []float64{
				1, 2, 3, 0, 0, 0, 0, 0, 0, 0,
				4, 5, 6, 7, 0, 0, 0, 0, 0, 0,
				0, 8, 9, 10, 11, 0, 0, 0, 0, 0,
				0, 0, 12, 13, 14, 15, 0, 0, 0, 0,
				0, 0, 0, 16, 17, 18, 19, 0, 0, 0,
				0, 0, 0, 0, 20, 21, 22, 23, 0, 0,
			}),
		},
	} {
		band := NewBandDense(test.r, test.c, test.kl, test.ku, test.data)
		rows, cols := band.Dims()

		if rows != test.r {
			t.Errorf("unexpected number of rows for test %d: got: %d want: %d", i, rows, test.r)
		}
		if cols != test.c {
			t.Errorf("unexpected number of cols for test %d: got: %d want: %d", i, cols, test.c)
		}
		if !reflect.DeepEqual(band, test.mat) {
			t.Errorf("unexpected value via reflect for test %d: got: %v want: %v", i, band, test.mat)
		}
		if !Equal(band, test.mat) {
			t.Errorf("unexpected value via mat.Equal for test %d: got: %v want: %v", i, band, test.mat)
		}
		if !Equal(band, test.dense) {
			t.Errorf("unexpected value via mat.Equal(band, dense) for test %d:\ngot:\n% v\nwant:\n% v", i, Formatted(band), Formatted(test.dense))
		}
	}
}

func TestNewDiagonalRect(t *testing.T) {
	for i, test := range []struct {
		data  []float64
		r, c  int
		mat   *BandDense
		dense *Dense
	}{
		{
			data: []float64{1, 2, 3, 4, 5, 6},
			r:    6, c: 6,
			mat: &BandDense{
				mat: blas64.Band{
					Rows:   6,
					Cols:   6,
					Stride: 1,
					Data:   []float64{1, 2, 3, 4, 5, 6},
				},
			},
			dense: NewDense(6, 6, []float64{
				1, 0, 0, 0, 0, 0,
				0, 2, 0, 0, 0, 0,
				0, 0, 3, 0, 0, 0,
				0, 0, 0, 4, 0, 0,
				0, 0, 0, 0, 5, 0,
				0, 0, 0, 0, 0, 6,
			}),
		},
		{
			data: []float64{1, 2, 3, 4, 5, 6},
			r:    7, c: 6,
			mat: &BandDense{
				mat: blas64.Band{
					Rows:   7,
					Cols:   6,
					Stride: 1,
					Data:   []float64{1, 2, 3, 4, 5, 6},
				},
			},
			dense: NewDense(7, 6, []float64{
				1, 0, 0, 0, 0, 0,
				0, 2, 0, 0, 0, 0,
				0, 0, 3, 0, 0, 0,
				0, 0, 0, 4, 0, 0,
				0, 0, 0, 0, 5, 0,
				0, 0, 0, 0, 0, 6,
				0, 0, 0, 0, 0, 0,
			}),
		},
		{
			data: []float64{1, 2, 3, 4, 5, 6},
			r:    6, c: 7,
			mat: &BandDense{
				mat: blas64.Band{
					Rows:   6,
					Cols:   7,
					Stride: 1,
					Data:   []float64{1, 2, 3, 4, 5, 6},
				},
			},
			dense: NewDense(6, 7, []float64{
				1, 0, 0, 0, 0, 0, 0,
				0, 2, 0, 0, 0, 0, 0,
				0, 0, 3, 0, 0, 0, 0,
				0, 0, 0, 4, 0, 0, 0,
				0, 0, 0, 0, 5, 0, 0,
				0, 0, 0, 0, 0, 6, 0,
			}),
		},
	} {
		band := NewDiagonalRect(test.r, test.c, test.data)
		rows, cols := band.Dims()

		if rows != test.r {
			t.Errorf("unexpected number of rows for test %d: got: %d want: %d", i, rows, test.r)
		}
		if cols != test.c {
			t.Errorf("unexpected number of cols for test %d: got: %d want: %d", i, cols, test.c)
		}
		if !reflect.DeepEqual(band, test.mat) {
			t.Errorf("unexpected value via reflect for test %d: got: %v want: %v", i, band, test.mat)
		}
		if !Equal(band, test.mat) {
			t.Errorf("unexpected value via mat.Equal for test %d: got: %v want: %v", i, band, test.mat)
		}
		if !Equal(band, test.dense) {
			t.Errorf("unexpected value via mat.Equal(band, dense) for test %d:\ngot:\n% v\nwant:\n% v", i, Formatted(band), Formatted(test.dense))
		}
	}
}

func TestBandAtSet(t *testing.T) {
	// 2  3  4  0  0  0
	// 5  6  7  8  0  0
	// 0  9 10 11 12  0
	// 0  0 13 14 15 16
	// 0  0  0 17 18 19
	// 0  0  0  0 21 22
	band := NewBandDense(6, 6, 1, 2, []float64{
		-1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
		17, 18, 19, -1,
		21, 22, -1, -1,
	})

	rows, cols := band.Dims()
	kl, ku := band.Bandwidth()

	// Explicitly test all indexes.
	want := bandImplicit{rows, cols, kl, ku, func(i, j int) float64 {
		return float64(i*(kl+ku) + j + kl + 1)
	}}
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			if band.At(i, j) != want.At(i, j) {
				t.Errorf("unexpected value for band.At(%d, %d): got:%v want:%v", i, j, band.At(i, j), want.At(i, j))
			}
		}
	}
	// Do that same thing via a call to Equal.
	if !Equal(band, want) {
		t.Errorf("unexpected value via mat.Equal:\ngot:\n% v\nwant:\n% v", Formatted(band), Formatted(want))
	}

	// Check At out of bounds
	for _, row := range []int{-1, rows, rows + 1} {
		panicked, message := panics(func() { band.At(row, 0) })
		if !panicked || message != ErrRowAccess.Error() {
			t.Errorf("expected panic for invalid row access N=%d r=%d", rows, row)
		}
	}
	for _, col := range []int{-1, cols, cols + 1} {
		panicked, message := panics(func() { band.At(0, col) })
		if !panicked || message != ErrColAccess.Error() {
			t.Errorf("expected panic for invalid column access N=%d c=%d", cols, col)
		}
	}

	// Check Set out of bounds
	for _, row := range []int{-1, rows, rows + 1} {
		panicked, message := panics(func() { band.SetBand(row, 0, 1.2) })
		if !panicked || message != ErrRowAccess.Error() {
			t.Errorf("expected panic for invalid row access N=%d r=%d", rows, row)
		}
	}
	for _, col := range []int{-1, cols, cols + 1} {
		panicked, message := panics(func() { band.SetBand(0, col, 1.2) })
		if !panicked || message != ErrColAccess.Error() {
			t.Errorf("expected panic for invalid column access N=%d c=%d", cols, col)
		}
	}

	for _, st := range []struct {
		row, col int
	}{
		{row: 0, col: 3},
		{row: 0, col: 4},
		{row: 0, col: 5},
		{row: 1, col: 4},
		{row: 1, col: 5},
		{row: 2, col: 5},
		{row: 2, col: 0},
		{row: 3, col: 1},
		{row: 4, col: 2},
		{row: 5, col: 3},
	} {
		panicked, message := panics(func() { band.SetBand(st.row, st.col, 1.2) })
		if !panicked || message != ErrBandSet.Error() {
			t.Errorf("expected panic for %+v %s", st, message)
		}
	}

	for _, st := range []struct {
		row, col  int
		orig, new float64
	}{
		{row: 1, col: 2, orig: 7, new: 15},
		{row: 2, col: 3, orig: 11, new: 15},
	} {
		if e := band.At(st.row, st.col); e != st.orig {
			t.Errorf("unexpected value for At(%d, %d): got: %v want: %v", st.row, st.col, e, st.orig)
		}
		band.SetBand(st.row, st.col, st.new)
		if e := band.At(st.row, st.col); e != st.new {
			t.Errorf("unexpected value for At(%d, %d) after SetBand(%[1]d, %d, %v): got: %v want: %[3]v", st.row, st.col, st.new, e)
		}
	}
}

// bandImplicit is an implicit band matrix returning val(i, j)
// for the value at (i, j).
type bandImplicit struct {
	r, c, kl, ku int
	val          func(i, j int) float64
}

func (b bandImplicit) Dims() (r, c int) {
	return b.r, b.c
}

func (b bandImplicit) T() Matrix {
	return Transpose{b}
}

func (b bandImplicit) At(i, j int) float64 {
	if i < 0 || b.r <= i {
		panic("row")
	}
	if j < 0 || b.c <= j {
		panic("col")
	}
	if j < i-b.kl || i+b.ku < j {
		return 0
	}
	return b.val(i, j)
}

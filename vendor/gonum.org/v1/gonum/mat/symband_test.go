// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat

import (
	"reflect"
	"testing"

	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/blas/blas64"
)

func TestNewSymBand(t *testing.T) {
	for i, test := range []struct {
		data  []float64
		n     int
		k     int
		mat   *SymBandDense
		dense *Dense
	}{
		{
			data: []float64{
				1, 2, 3,
				4, 5, 6,
				7, 8, 9,
				10, 11, 12,
				13, 14, -1,
				15, -1, -1,
			},
			n: 6,
			k: 2,
			mat: &SymBandDense{
				mat: blas64.SymmetricBand{
					N:      6,
					K:      2,
					Stride: 3,
					Uplo:   blas.Upper,
					Data: []float64{
						1, 2, 3,
						4, 5, 6,
						7, 8, 9,
						10, 11, 12,
						13, 14, -1,
						15, -1, -1,
					},
				},
			},
			dense: NewDense(6, 6, []float64{
				1, 2, 3, 0, 0, 0,
				2, 4, 5, 6, 0, 0,
				3, 5, 7, 8, 9, 0,
				0, 6, 8, 10, 11, 12,
				0, 0, 9, 11, 13, 14,
				0, 0, 0, 12, 14, 15,
			}),
		},
	} {
		band := NewSymBandDense(test.n, test.k, test.data)
		rows, cols := band.Dims()

		if rows != test.n {
			t.Errorf("unexpected number of rows for test %d: got: %d want: %d", i, rows, test.n)
		}
		if cols != test.n {
			t.Errorf("unexpected number of cols for test %d: got: %d want: %d", i, cols, test.n)
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

func TestNewDiagonal(t *testing.T) {
	for i, test := range []struct {
		data  []float64
		n     int
		mat   *SymBandDense
		dense *Dense
	}{
		{
			data: []float64{1, 2, 3, 4, 5, 6},
			n:    6,
			mat: &SymBandDense{
				mat: blas64.SymmetricBand{
					N:      6,
					Stride: 1,
					Uplo:   blas.Upper,
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
	} {
		band := NewDiagonal(test.n, test.data)
		rows, cols := band.Dims()

		if rows != test.n {
			t.Errorf("unexpected number of rows for test %d: got: %d want: %d", i, rows, test.n)
		}
		if cols != test.n {
			t.Errorf("unexpected number of cols for test %d: got: %d want: %d", i, cols, test.n)
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

func TestSymBandAtSet(t *testing.T) {
	// 1  2  3  0  0  0
	// 2  4  5  6  0  0
	// 3  5  7  8  9  0
	// 0  6  8 10 11 12
	// 0  0  9 11 13 14
	// 0  0  0 12 14 16
	band := NewSymBandDense(6, 2, []float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
		10, 11, 12,
		13, 14, -1,
		16, -1, -1,
	})

	rows, cols := band.Dims()
	kl, ku := band.Bandwidth()

	// Explicitly test all indexes.
	want := bandImplicit{rows, cols, kl, ku, func(i, j int) float64 {
		if i > j {
			i, j = j, i
		}
		return float64(i*ku + j + 1)
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
		panicked, message := panics(func() { band.SetSymBand(row, 0, 1.2) })
		if !panicked || message != ErrRowAccess.Error() {
			t.Errorf("expected panic for invalid row access N=%d r=%d", rows, row)
		}
	}
	for _, col := range []int{-1, cols, cols + 1} {
		panicked, message := panics(func() { band.SetSymBand(0, col, 1.2) })
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
		{row: 3, col: 0},
		{row: 4, col: 1},
		{row: 5, col: 2},
	} {
		panicked, message := panics(func() { band.SetSymBand(st.row, st.col, 1.2) })
		if !panicked || message != ErrBandSet.Error() {
			t.Errorf("expected panic for %+v %s", st, message)
		}
	}

	for _, st := range []struct {
		row, col  int
		orig, new float64
	}{
		{row: 1, col: 2, orig: 5, new: 15},
		{row: 2, col: 3, orig: 8, new: 15},
	} {
		if e := band.At(st.row, st.col); e != st.orig {
			t.Errorf("unexpected value for At(%d, %d): got: %v want: %v", st.row, st.col, e, st.orig)
		}
		band.SetSymBand(st.row, st.col, st.new)
		if e := band.At(st.row, st.col); e != st.new {
			t.Errorf("unexpected value for At(%d, %d) after SetSymBand(%[1]d, %d, %v): got: %v want: %[3]v", st.row, st.col, st.new, e)
		}
	}
}

// Copyright ©2013 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat64

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"testing"

	"github.com/gonum/blas/blas64"
	"github.com/gonum/floats"
	"github.com/gonum/matrix"
)

func asBasicMatrix(d *Dense) Matrix            { return (*basicMatrix)(d) }
func asBasicSymmetric(s *SymDense) Matrix      { return (*basicSymmetric)(s) }
func asBasicTriangular(t *TriDense) Triangular { return (*basicTriangular)(t) }

func TestNewDense(t *testing.T) {
	for i, test := range []struct {
		a          []float64
		rows, cols int
		min, max   float64
		fro        float64
		mat        *Dense
	}{
		{
			[]float64{
				0, 0, 0,
				0, 0, 0,
				0, 0, 0,
			},
			3, 3,
			0, 0,
			0,
			&Dense{
				mat: blas64.General{
					Rows: 3, Cols: 3,
					Stride: 3,
					Data:   []float64{0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
				capRows: 3, capCols: 3,
			},
		},
		{
			[]float64{
				1, 1, 1,
				1, 1, 1,
				1, 1, 1,
			},
			3, 3,
			1, 1,
			3,
			&Dense{
				mat: blas64.General{
					Rows: 3, Cols: 3,
					Stride: 3,
					Data:   []float64{1, 1, 1, 1, 1, 1, 1, 1, 1},
				},
				capRows: 3, capCols: 3,
			},
		},
		{
			[]float64{
				1, 0, 0,
				0, 1, 0,
				0, 0, 1,
			},
			3, 3,
			0, 1,
			1.7320508075688772,
			&Dense{
				mat: blas64.General{
					Rows: 3, Cols: 3,
					Stride: 3,
					Data:   []float64{1, 0, 0, 0, 1, 0, 0, 0, 1},
				},
				capRows: 3, capCols: 3,
			},
		},
		{
			[]float64{
				-1, 0, 0,
				0, -1, 0,
				0, 0, -1,
			},
			3, 3,
			-1, 0,
			1.7320508075688772,
			&Dense{
				mat: blas64.General{
					Rows: 3, Cols: 3,
					Stride: 3,
					Data:   []float64{-1, 0, 0, 0, -1, 0, 0, 0, -1},
				},
				capRows: 3, capCols: 3,
			},
		},
		{
			[]float64{
				1, 2, 3,
				4, 5, 6,
			},
			2, 3,
			1, 6,
			9.539392014169458,
			&Dense{
				mat: blas64.General{
					Rows: 2, Cols: 3,
					Stride: 3,
					Data:   []float64{1, 2, 3, 4, 5, 6},
				},
				capRows: 2, capCols: 3,
			},
		},
		{
			[]float64{
				1, 2,
				3, 4,
				5, 6,
			},
			3, 2,
			1, 6,
			9.539392014169458,
			&Dense{
				mat: blas64.General{
					Rows: 3, Cols: 2,
					Stride: 2,
					Data:   []float64{1, 2, 3, 4, 5, 6},
				},
				capRows: 3, capCols: 2,
			},
		},
	} {
		m := NewDense(test.rows, test.cols, test.a)
		rows, cols := m.Dims()
		if rows != test.rows {
			t.Errorf("unexpected number of rows for test %d: got: %d want: %d", i, rows, test.rows)
		}
		if cols != test.cols {
			t.Errorf("unexpected number of cols for test %d: got: %d want: %d", i, cols, test.cols)
		}
		if min := Min(m); min != test.min {
			t.Errorf("unexpected min for test %d: got: %v want: %v", i, min, test.min)
		}
		if max := Max(m); max != test.max {
			t.Errorf("unexpected max for test %d: got: %v want: %v", i, max, test.max)
		}
		if fro := Norm(m, 2); math.Abs(Norm(m, 2)-test.fro) > 1e-14 {
			t.Errorf("unexpected Frobenius norm for test %d: got: %v want: %v", i, fro, test.fro)
		}
		if !reflect.DeepEqual(m, test.mat) {
			t.Errorf("unexpected matrix for test %d", i)
		}
		if !Equal(m, test.mat) {
			t.Errorf("matrix does not equal expected matrix for test %d", i)
		}
	}
}

func TestAtSet(t *testing.T) {
	for test, af := range [][][]float64{
		{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, // even
		{{1, 2}, {4, 5}, {7, 8}},          // wide
		{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, //skinny
	} {
		m := NewDense(flatten(af))
		rows, cols := m.Dims()
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				if m.At(i, j) != af[i][j] {
					t.Errorf("unexpected value for At(%d, %d) for test %d: got: %v want: %v",
						i, j, test, m.At(i, j), af[i][j])
				}

				v := float64(i * j)
				m.Set(i, j, v)
				if m.At(i, j) != v {
					t.Errorf("unexpected value for At(%d, %d) after Set(%[1]d, %d, %v) for test %d: got: %v want: %[3]v",
						i, j, v, test, m.At(i, j))
				}
			}
		}
		// Check access out of bounds fails
		for _, row := range []int{-1, rows, rows + 1} {
			panicked, message := panics(func() { m.At(row, 0) })
			if !panicked || message != matrix.ErrRowAccess.Error() {
				t.Errorf("expected panic for invalid row access N=%d r=%d", rows, row)
			}
		}
		for _, col := range []int{-1, cols, cols + 1} {
			panicked, message := panics(func() { m.At(0, col) })
			if !panicked || message != matrix.ErrColAccess.Error() {
				t.Errorf("expected panic for invalid column access N=%d c=%d", cols, col)
			}
		}

		// Check Set out of bounds
		for _, row := range []int{-1, rows, rows + 1} {
			panicked, message := panics(func() { m.Set(row, 0, 1.2) })
			if !panicked || message != matrix.ErrRowAccess.Error() {
				t.Errorf("expected panic for invalid row access N=%d r=%d", rows, row)
			}
		}
		for _, col := range []int{-1, cols, cols + 1} {
			panicked, message := panics(func() { m.Set(0, col, 1.2) })
			if !panicked || message != matrix.ErrColAccess.Error() {
				t.Errorf("expected panic for invalid column access N=%d c=%d", cols, col)
			}
		}
	}
}

func TestSetRowColumn(t *testing.T) {
	for _, as := range [][][]float64{
		{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}},
		{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
	} {
		for ri, row := range as {
			a := NewDense(flatten(as))
			m := &Dense{}
			m.Clone(a)
			a.SetRow(ri, make([]float64, a.mat.Cols))
			m.Sub(m, a)
			nt := Norm(m, 2)
			nr := floats.Norm(row, 2)
			if math.Abs(nt-nr) > 1e-14 {
				t.Errorf("Row %d norm mismatch, want: %g, got: %g", ri, nr, nt)
			}
		}

		for ci := range as[0] {
			a := NewDense(flatten(as))
			m := &Dense{}
			m.Clone(a)
			a.SetCol(ci, make([]float64, a.mat.Rows))
			col := make([]float64, a.mat.Rows)
			for j := range col {
				col[j] = float64(ci + 1 + j*a.mat.Cols)
			}
			m.Sub(m, a)
			nt := Norm(m, 2)
			nc := floats.Norm(col, 2)
			if math.Abs(nt-nc) > 1e-14 {
				t.Errorf("Column %d norm mismatch, want: %g, got: %g", ci, nc, nt)
			}
		}
	}
}

func TestRowColView(t *testing.T) {
	for _, test := range []struct {
		mat [][]float64
	}{
		{
			mat: [][]float64{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25},
			},
		},
		{
			mat: [][]float64{
				{1, 2, 3, 4},
				{6, 7, 8, 9},
				{11, 12, 13, 14},
				{16, 17, 18, 19},
				{21, 22, 23, 24},
			},
		},
		{
			mat: [][]float64{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
			},
		},
	} {
		// This over cautious approach to building a matrix data
		// slice is to ensure that changes to flatten in the future
		// do not mask a regression to the issue identified in
		// gonum/matrix#110.
		rows, cols, flat := flatten(test.mat)
		m := NewDense(rows, cols, flat[:len(flat):len(flat)])

		for _, row := range []int{-1, rows, rows + 1} {
			panicked, message := panics(func() { m.At(row, 0) })
			if !panicked || message != matrix.ErrRowAccess.Error() {
				t.Errorf("expected panic for invalid row access rows=%d r=%d", rows, row)
			}
		}
		for _, col := range []int{-1, cols, cols + 1} {
			panicked, message := panics(func() { m.At(0, col) })
			if !panicked || message != matrix.ErrColAccess.Error() {
				t.Errorf("expected panic for invalid column access cols=%d c=%d", cols, col)
			}
		}

		for i := 0; i < rows; i++ {
			vr := m.RowView(i)
			if vr.Len() != cols {
				t.Errorf("unexpected number of columns: got: %d want: %d", vr.Len(), cols)
			}
			for j := 0; j < cols; j++ {
				if got := vr.At(j, 0); got != test.mat[i][j] {
					t.Errorf("unexpected value for row.At(%d, 0): got: %v want: %v",
						j, got, test.mat[i][j])
				}
			}
		}
		for j := 0; j < cols; j++ {
			vc := m.ColView(j)
			if vc.Len() != rows {
				t.Errorf("unexpected number of rows: got: %d want: %d", vc.Len(), rows)
			}
			for i := 0; i < rows; i++ {
				if got := vc.At(i, 0); got != test.mat[i][j] {
					t.Errorf("unexpected value for col.At(%d, 0): got: %v want: %v",
						i, got, test.mat[i][j])
				}
			}
		}
		m = m.Slice(1, rows-1, 1, cols-1).(*Dense)
		for i := 1; i < rows-1; i++ {
			vr := m.RowView(i - 1)
			if vr.Len() != cols-2 {
				t.Errorf("unexpected number of columns: got: %d want: %d", vr.Len(), cols-2)
			}
			for j := 1; j < cols-1; j++ {
				if got := vr.At(j-1, 0); got != test.mat[i][j] {
					t.Errorf("unexpected value for row.At(%d, 0): got: %v want: %v",
						j-1, got, test.mat[i][j])
				}
			}
		}
		for j := 1; j < cols-1; j++ {
			vc := m.ColView(j - 1)
			if vc.Len() != rows-2 {
				t.Errorf("unexpected number of rows: got: %d want: %d", vc.Len(), rows-2)
			}
			for i := 1; i < rows-1; i++ {
				if got := vc.At(i-1, 0); got != test.mat[i][j] {
					t.Errorf("unexpected value for col.At(%d, 0): got: %v want: %v",
						i-1, got, test.mat[i][j])
				}
			}
		}
	}
}

func TestGrow(t *testing.T) {
	m := &Dense{}
	m = m.Grow(10, 10).(*Dense)
	rows, cols := m.Dims()
	capRows, capCols := m.Caps()
	if rows != 10 {
		t.Errorf("unexpected value for rows: got: %d want: 10", rows)
	}
	if cols != 10 {
		t.Errorf("unexpected value for cols: got: %d want: 10", cols)
	}
	if capRows != 10 {
		t.Errorf("unexpected value for capRows: got: %d want: 10", capRows)
	}
	if capCols != 10 {
		t.Errorf("unexpected value for capCols: got: %d want: 10", capCols)
	}

	// Test grow within caps is in-place.
	m.Set(1, 1, 1)
	v := m.Slice(1, 5, 1, 5).(*Dense)
	if v.At(0, 0) != m.At(1, 1) {
		t.Errorf("unexpected viewed element value: got: %v want: %v", v.At(0, 0), m.At(1, 1))
	}
	v = v.Grow(5, 5).(*Dense)
	if !Equal(v, m.Slice(1, 10, 1, 10)) {
		t.Error("unexpected view value after grow")
	}

	// Test grow bigger than caps copies.
	v = v.Grow(5, 5).(*Dense)
	if !Equal(v.Slice(0, 9, 0, 9), m.Slice(1, 10, 1, 10)) {
		t.Error("unexpected mismatched common view value after grow")
	}
	v.Set(0, 0, 0)
	if Equal(v.Slice(0, 9, 0, 9), m.Slice(1, 10, 1, 10)) {
		t.Error("unexpected matching view value after grow past capacity")
	}

	// Test grow uses existing data slice when matrix is zero size.
	v.Reset()
	p, l := &v.mat.Data[:1][0], cap(v.mat.Data)
	*p = 1 // This element is at position (-1, -1) relative to v and so should not be visible.
	v = v.Grow(5, 5).(*Dense)
	if &v.mat.Data[:1][0] != p {
		t.Error("grow unexpectedly copied slice within cap limit")
	}
	if cap(v.mat.Data) != l {
		t.Errorf("unexpected change in data slice capacity: got: %d want: %d", cap(v.mat.Data), l)
	}
	if v.At(0, 0) != 0 {
		t.Errorf("unexpected value for At(0, 0): got: %v want: 0", v.At(0, 0))
	}
}

func TestAdd(t *testing.T) {
	for i, test := range []struct {
		a, b, r [][]float64
	}{
		{
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		},
		{
			[][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			[][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			[][]float64{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}},
		},
		{
			[][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			[][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			[][]float64{{2, 0, 0}, {0, 2, 0}, {0, 0, 2}},
		},
		{
			[][]float64{{-1, 0, 0}, {0, -1, 0}, {0, 0, -1}},
			[][]float64{{-1, 0, 0}, {0, -1, 0}, {0, 0, -1}},
			[][]float64{{-2, 0, 0}, {0, -2, 0}, {0, 0, -2}},
		},
		{
			[][]float64{{1, 2, 3}, {4, 5, 6}},
			[][]float64{{1, 2, 3}, {4, 5, 6}},
			[][]float64{{2, 4, 6}, {8, 10, 12}},
		},
	} {
		a := NewDense(flatten(test.a))
		b := NewDense(flatten(test.b))
		r := NewDense(flatten(test.r))

		var temp Dense
		temp.Add(a, b)
		if !Equal(&temp, r) {
			t.Errorf("unexpected result from Add for test %d %v Add %v: got: %v want: %v",
				i, test.a, test.b, unflatten(temp.mat.Rows, temp.mat.Cols, temp.mat.Data), test.r)
		}

		zero(temp.mat.Data)
		temp.Add(a, b)
		if !Equal(&temp, r) {
			t.Errorf("unexpected result from Add for test %d %v Add %v: got: %v want: %v",
				i, test.a, test.b, unflatten(temp.mat.Rows, temp.mat.Cols, temp.mat.Data), test.r)
		}

		// These probably warrant a better check and failure. They should never happen in the wild though.
		temp.mat.Data = nil
		panicked, message := panics(func() { temp.Add(a, b) })
		if !panicked || message != "runtime error: index out of range" {
			t.Error("exected runtime panic for nil data slice")
		}

		a.Add(a, b)
		if !Equal(a, r) {
			t.Errorf("unexpected result from Add for test %d %v Add %v: got: %v want: %v",
				i, test.a, test.b, unflatten(a.mat.Rows, a.mat.Cols, a.mat.Data), test.r)
		}
	}

	panicked, message := panics(func() {
		m := NewDense(10, 10, nil)
		a := NewDense(5, 5, nil)
		m.Slice(1, 6, 1, 6).(*Dense).Add(a, m.Slice(2, 7, 2, 7))
	})
	if !panicked {
		t.Error("expected panic for overlapping matrices")
	}
	if message != regionOverlap {
		t.Errorf("unexpected panic message: got: %q want: %q", message, regionOverlap)
	}

	method := func(receiver, a, b Matrix) {
		type Adder interface {
			Add(a, b Matrix)
		}
		rd := receiver.(Adder)
		rd.Add(a, b)
	}
	denseComparison := func(receiver, a, b *Dense) {
		receiver.Add(a, b)
	}
	testTwoInput(t, "Add", &Dense{}, method, denseComparison, legalTypesAll, legalSizeSameRectangular, 1e-14)
}

func TestSub(t *testing.T) {
	for i, test := range []struct {
		a, b, r [][]float64
	}{
		{
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		},
		{
			[][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			[][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		},
		{
			[][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			[][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		},
		{
			[][]float64{{-1, 0, 0}, {0, -1, 0}, {0, 0, -1}},
			[][]float64{{-1, 0, 0}, {0, -1, 0}, {0, 0, -1}},
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		},
		{
			[][]float64{{1, 2, 3}, {4, 5, 6}},
			[][]float64{{1, 2, 3}, {4, 5, 6}},
			[][]float64{{0, 0, 0}, {0, 0, 0}},
		},
	} {
		a := NewDense(flatten(test.a))
		b := NewDense(flatten(test.b))
		r := NewDense(flatten(test.r))

		var temp Dense
		temp.Sub(a, b)
		if !Equal(&temp, r) {
			t.Errorf("unexpected result from Sub for test %d %v Sub %v: got: %v want: %v",
				i, test.a, test.b, unflatten(temp.mat.Rows, temp.mat.Cols, temp.mat.Data), test.r)
		}

		zero(temp.mat.Data)
		temp.Sub(a, b)
		if !Equal(&temp, r) {
			t.Errorf("unexpected result from Sub for test %d %v Sub %v: got: %v want: %v",
				i, test.a, test.b, unflatten(temp.mat.Rows, temp.mat.Cols, temp.mat.Data), test.r)
		}

		// These probably warrant a better check and failure. They should never happen in the wild though.
		temp.mat.Data = nil
		panicked, message := panics(func() { temp.Sub(a, b) })
		if !panicked || message != "runtime error: index out of range" {
			t.Error("exected runtime panic for nil data slice")
		}

		a.Sub(a, b)
		if !Equal(a, r) {
			t.Errorf("unexpected result from Sub for test %d %v Sub %v: got: %v want: %v",
				i, test.a, test.b, unflatten(a.mat.Rows, a.mat.Cols, a.mat.Data), test.r)
		}
	}

	panicked, message := panics(func() {
		m := NewDense(10, 10, nil)
		a := NewDense(5, 5, nil)
		m.Slice(1, 6, 1, 6).(*Dense).Sub(a, m.Slice(2, 7, 2, 7))
	})
	if !panicked {
		t.Error("expected panic for overlapping matrices")
	}
	if message != regionOverlap {
		t.Errorf("unexpected panic message: got: %q want: %q", message, regionOverlap)
	}

	method := func(receiver, a, b Matrix) {
		type Suber interface {
			Sub(a, b Matrix)
		}
		rd := receiver.(Suber)
		rd.Sub(a, b)
	}
	denseComparison := func(receiver, a, b *Dense) {
		receiver.Sub(a, b)
	}
	testTwoInput(t, "Sub", &Dense{}, method, denseComparison, legalTypesAll, legalSizeSameRectangular, 1e-14)
}

func TestMulElem(t *testing.T) {
	for i, test := range []struct {
		a, b, r [][]float64
	}{
		{
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		},
		{
			[][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			[][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			[][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
		},
		{
			[][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			[][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			[][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
		},
		{
			[][]float64{{-1, 0, 0}, {0, -1, 0}, {0, 0, -1}},
			[][]float64{{-1, 0, 0}, {0, -1, 0}, {0, 0, -1}},
			[][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
		},
		{
			[][]float64{{1, 2, 3}, {4, 5, 6}},
			[][]float64{{1, 2, 3}, {4, 5, 6}},
			[][]float64{{1, 4, 9}, {16, 25, 36}},
		},
	} {
		a := NewDense(flatten(test.a))
		b := NewDense(flatten(test.b))
		r := NewDense(flatten(test.r))

		var temp Dense
		temp.MulElem(a, b)
		if !Equal(&temp, r) {
			t.Errorf("unexpected result from MulElem for test %d %v MulElem %v: got: %v want: %v",
				i, test.a, test.b, unflatten(temp.mat.Rows, temp.mat.Cols, temp.mat.Data), test.r)
		}

		zero(temp.mat.Data)
		temp.MulElem(a, b)
		if !Equal(&temp, r) {
			t.Errorf("unexpected result from MulElem for test %d %v MulElem %v: got: %v want: %v",
				i, test.a, test.b, unflatten(temp.mat.Rows, temp.mat.Cols, temp.mat.Data), test.r)
		}

		// These probably warrant a better check and failure. They should never happen in the wild though.
		temp.mat.Data = nil
		panicked, message := panics(func() { temp.MulElem(a, b) })
		if !panicked || message != "runtime error: index out of range" {
			t.Error("exected runtime panic for nil data slice")
		}

		a.MulElem(a, b)
		if !Equal(a, r) {
			t.Errorf("unexpected result from MulElem for test %d %v MulElem %v: got: %v want: %v",
				i, test.a, test.b, unflatten(a.mat.Rows, a.mat.Cols, a.mat.Data), test.r)
		}
	}

	panicked, message := panics(func() {
		m := NewDense(10, 10, nil)
		a := NewDense(5, 5, nil)
		m.Slice(1, 6, 1, 6).(*Dense).MulElem(a, m.Slice(2, 7, 2, 7))
	})
	if !panicked {
		t.Error("expected panic for overlapping matrices")
	}
	if message != regionOverlap {
		t.Errorf("unexpected panic message: got: %q want: %q", message, regionOverlap)
	}

	method := func(receiver, a, b Matrix) {
		type ElemMuler interface {
			MulElem(a, b Matrix)
		}
		rd := receiver.(ElemMuler)
		rd.MulElem(a, b)
	}
	denseComparison := func(receiver, a, b *Dense) {
		receiver.MulElem(a, b)
	}
	testTwoInput(t, "MulElem", &Dense{}, method, denseComparison, legalTypesAll, legalSizeSameRectangular, 1e-14)
}

// A comparison that treats NaNs as equal, for testing.
func (m *Dense) same(b Matrix) bool {
	br, bc := b.Dims()
	if br != m.mat.Rows || bc != m.mat.Cols {
		return false
	}
	for r := 0; r < br; r++ {
		for c := 0; c < bc; c++ {
			if av, bv := m.At(r, c), b.At(r, c); av != bv && !(math.IsNaN(av) && math.IsNaN(bv)) {
				return false
			}
		}
	}
	return true
}

func TestDivElem(t *testing.T) {
	for i, test := range []struct {
		a, b, r [][]float64
	}{
		{
			[][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			[][]float64{{math.Inf(1), math.NaN(), math.NaN()}, {math.NaN(), math.Inf(1), math.NaN()}, {math.NaN(), math.NaN(), math.Inf(1)}},
		},
		{
			[][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			[][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			[][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
		},
		{
			[][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			[][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			[][]float64{{1, math.NaN(), math.NaN()}, {math.NaN(), 1, math.NaN()}, {math.NaN(), math.NaN(), 1}},
		},
		{
			[][]float64{{-1, 0, 0}, {0, -1, 0}, {0, 0, -1}},
			[][]float64{{-1, 0, 0}, {0, -1, 0}, {0, 0, -1}},
			[][]float64{{1, math.NaN(), math.NaN()}, {math.NaN(), 1, math.NaN()}, {math.NaN(), math.NaN(), 1}},
		},
		{
			[][]float64{{1, 2, 3}, {4, 5, 6}},
			[][]float64{{1, 2, 3}, {4, 5, 6}},
			[][]float64{{1, 1, 1}, {1, 1, 1}},
		},
	} {
		a := NewDense(flatten(test.a))
		b := NewDense(flatten(test.b))
		r := NewDense(flatten(test.r))

		var temp Dense
		temp.DivElem(a, b)
		if !temp.same(r) {
			t.Errorf("unexpected result from DivElem for test %d %v DivElem %v: got: %v want: %v",
				i, test.a, test.b, unflatten(temp.mat.Rows, temp.mat.Cols, temp.mat.Data), test.r)
		}

		zero(temp.mat.Data)
		temp.DivElem(a, b)
		if !temp.same(r) {
			t.Errorf("unexpected result from DivElem for test %d %v DivElem %v: got: %v want: %v",
				i, test.a, test.b, unflatten(temp.mat.Rows, temp.mat.Cols, temp.mat.Data), test.r)
		}

		// These probably warrant a better check and failure. They should never happen in the wild though.
		temp.mat.Data = nil
		panicked, message := panics(func() { temp.DivElem(a, b) })
		if !panicked || message != "runtime error: index out of range" {
			t.Error("exected runtime panic for nil data slice")
		}

		a.DivElem(a, b)
		if !a.same(r) {
			t.Errorf("unexpected result from DivElem for test %d %v DivElem %v: got: %v want: %v",
				i, test.a, test.b, unflatten(a.mat.Rows, a.mat.Cols, a.mat.Data), test.r)
		}
	}

	panicked, message := panics(func() {
		m := NewDense(10, 10, nil)
		a := NewDense(5, 5, nil)
		m.Slice(1, 6, 1, 6).(*Dense).DivElem(a, m.Slice(2, 7, 2, 7))
	})
	if !panicked {
		t.Error("expected panic for overlapping matrices")
	}
	if message != regionOverlap {
		t.Errorf("unexpected panic message: got: %q want: %q", message, regionOverlap)
	}

	method := func(receiver, a, b Matrix) {
		type ElemDiver interface {
			DivElem(a, b Matrix)
		}
		rd := receiver.(ElemDiver)
		rd.DivElem(a, b)
	}
	denseComparison := func(receiver, a, b *Dense) {
		receiver.DivElem(a, b)
	}
	testTwoInput(t, "DivElem", &Dense{}, method, denseComparison, legalTypesAll, legalSizeSameRectangular, 1e-14)
}

func TestMul(t *testing.T) {
	for i, test := range []struct {
		a, b, r [][]float64
	}{
		{
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		},
		{
			[][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			[][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			[][]float64{{3, 3, 3}, {3, 3, 3}, {3, 3, 3}},
		},
		{
			[][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			[][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			[][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
		},
		{
			[][]float64{{-1, 0, 0}, {0, -1, 0}, {0, 0, -1}},
			[][]float64{{-1, 0, 0}, {0, -1, 0}, {0, 0, -1}},
			[][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
		},
		{
			[][]float64{{1, 2, 3}, {4, 5, 6}},
			[][]float64{{1, 2}, {3, 4}, {5, 6}},
			[][]float64{{22, 28}, {49, 64}},
		},
		{
			[][]float64{{0, 1, 1}, {0, 1, 1}, {0, 1, 1}},
			[][]float64{{0, 1, 1}, {0, 1, 1}, {0, 1, 1}},
			[][]float64{{0, 2, 2}, {0, 2, 2}, {0, 2, 2}},
		},
	} {
		a := NewDense(flatten(test.a))
		b := NewDense(flatten(test.b))
		r := NewDense(flatten(test.r))

		var temp Dense
		temp.Mul(a, b)
		if !Equal(&temp, r) {
			t.Errorf("unexpected result from Mul for test %d %v Mul %v: got: %v want: %v",
				i, test.a, test.b, unflatten(temp.mat.Rows, temp.mat.Cols, temp.mat.Data), test.r)
		}

		zero(temp.mat.Data)
		temp.Mul(a, b)
		if !Equal(&temp, r) {
			t.Errorf("unexpected result from Mul for test %d %v Mul %v: got: %v want: %v",
				i, test.a, test.b, unflatten(temp.mat.Rows, temp.mat.Cols, temp.mat.Data), test.r)
		}

		// These probably warrant a better check and failure. They should never happen in the wild though.
		temp.mat.Data = nil
		panicked, message := panics(func() { temp.Mul(a, b) })
		if !panicked || message != "blas: insufficient matrix slice length" {
			if message != "" {
				t.Errorf("expected runtime panic for nil data slice: got %q", message)
			} else {
				t.Error("expected runtime panic for nil data slice")
			}
		}
	}

	panicked, message := panics(func() {
		m := NewDense(10, 10, nil)
		a := NewDense(5, 5, nil)
		m.Slice(1, 6, 1, 6).(*Dense).Mul(a, m.Slice(2, 7, 2, 7))
	})
	if !panicked {
		t.Error("expected panic for overlapping matrices")
	}
	if message != regionOverlap {
		t.Errorf("unexpected panic message: got: %q want: %q", message, regionOverlap)
	}

	method := func(receiver, a, b Matrix) {
		type Muler interface {
			Mul(a, b Matrix)
		}
		rd := receiver.(Muler)
		rd.Mul(a, b)
	}
	denseComparison := func(receiver, a, b *Dense) {
		receiver.Mul(a, b)
	}
	legalSizeMul := func(ar, ac, br, bc int) bool {
		return ac == br
	}
	testTwoInput(t, "Mul", &Dense{}, method, denseComparison, legalTypesAll, legalSizeMul, 1e-14)
}

func randDense(size int, rho float64, rnd func() float64) (*Dense, error) {
	if size == 0 {
		return nil, matrix.ErrZeroLength
	}
	d := &Dense{
		mat: blas64.General{
			Rows: size, Cols: size, Stride: size,
			Data: make([]float64, size*size),
		},
		capRows: size, capCols: size,
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if rand.Float64() < rho {
				d.Set(i, j, rnd())
			}
		}
	}
	return d, nil
}

func TestExp(t *testing.T) {
	for i, test := range []struct {
		a    [][]float64
		want [][]float64
		mod  func(*Dense)
	}{
		{
			a:    [][]float64{{-49, 24}, {-64, 31}},
			want: [][]float64{{-0.7357587581474017, 0.5518190996594223}, {-1.4715175990917921, 1.103638240717339}},
		},
		{
			a:    [][]float64{{-49, 24}, {-64, 31}},
			want: [][]float64{{-0.7357587581474017, 0.5518190996594223}, {-1.4715175990917921, 1.103638240717339}},
			mod: func(a *Dense) {
				d := make([]float64, 100)
				for i := range d {
					d[i] = math.NaN()
				}
				*a = *NewDense(10, 10, d).Slice(1, 3, 1, 3).(*Dense)
			},
		},
		{
			a:    [][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			want: [][]float64{{2.71828182845905, 0, 0}, {0, 2.71828182845905, 0}, {0, 0, 2.71828182845905}},
		},
	} {
		var got Dense
		if test.mod != nil {
			test.mod(&got)
		}
		got.Exp(NewDense(flatten(test.a)))
		if !EqualApprox(&got, NewDense(flatten(test.want)), 1e-12) {
			t.Errorf("unexpected result for Exp test %d", i)
		}
	}
}

func TestPow(t *testing.T) {
	for i, test := range []struct {
		a    [][]float64
		n    int
		mod  func(*Dense)
		want [][]float64
	}{
		{
			a:    [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			n:    0,
			want: [][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
		},
		{
			a:    [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			n:    0,
			want: [][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			mod: func(a *Dense) {
				d := make([]float64, 100)
				for i := range d {
					d[i] = math.NaN()
				}
				*a = *NewDense(10, 10, d).Slice(1, 4, 1, 4).(*Dense)
			},
		},
		{
			a:    [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			n:    1,
			want: [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
		{
			a:    [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			n:    1,
			want: [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			mod: func(a *Dense) {
				d := make([]float64, 100)
				for i := range d {
					d[i] = math.NaN()
				}
				*a = *NewDense(10, 10, d).Slice(1, 4, 1, 4).(*Dense)
			},
		},
		{
			a:    [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			n:    2,
			want: [][]float64{{30, 36, 42}, {66, 81, 96}, {102, 126, 150}},
		},
		{
			a:    [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			n:    2,
			want: [][]float64{{30, 36, 42}, {66, 81, 96}, {102, 126, 150}},
			mod: func(a *Dense) {
				d := make([]float64, 100)
				for i := range d {
					d[i] = math.NaN()
				}
				*a = *NewDense(10, 10, d).Slice(1, 4, 1, 4).(*Dense)
			},
		},
		{
			a:    [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			n:    3,
			want: [][]float64{{468, 576, 684}, {1062, 1305, 1548}, {1656, 2034, 2412}},
		},
		{
			a:    [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			n:    3,
			want: [][]float64{{468, 576, 684}, {1062, 1305, 1548}, {1656, 2034, 2412}},
			mod: func(a *Dense) {
				d := make([]float64, 100)
				for i := range d {
					d[i] = math.NaN()
				}
				*a = *NewDense(10, 10, d).Slice(1, 4, 1, 4).(*Dense)
			},
		},
	} {
		var got Dense
		if test.mod != nil {
			test.mod(&got)
		}
		got.Pow(NewDense(flatten(test.a)), test.n)
		if !EqualApprox(&got, NewDense(flatten(test.want)), 1e-12) {
			t.Errorf("unexpected result for Pow test %d", i)
		}
	}
}

func TestScale(t *testing.T) {
	for _, f := range []float64{0.5, 1, 3} {
		method := func(receiver, a Matrix) {
			type Scaler interface {
				Scale(f float64, a Matrix)
			}
			rd := receiver.(Scaler)
			rd.Scale(f, a)
		}
		denseComparison := func(receiver, a *Dense) {
			receiver.Scale(f, a)
		}
		testOneInput(t, "Scale", &Dense{}, method, denseComparison, isAnyType, isAnySize, 1e-14)
	}
}

func TestPowN(t *testing.T) {
	for i, test := range []struct {
		a    [][]float64
		mod  func(*Dense)
		want [][]float64
	}{
		{
			a: [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
		{
			a: [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			mod: func(a *Dense) {
				d := make([]float64, 100)
				for i := range d {
					d[i] = math.NaN()
				}
				*a = *NewDense(10, 10, d).Slice(1, 4, 1, 4).(*Dense)
			},
		},
	} {
		for n := 1; n <= 14; n++ {
			var got, want Dense
			if test.mod != nil {
				test.mod(&got)
			}
			got.Pow(NewDense(flatten(test.a)), n)
			want.iterativePow(NewDense(flatten(test.a)), n)
			if !Equal(&got, &want) {
				t.Errorf("unexpected result for iterative Pow test %d", i)
			}
		}
	}
}

func (m *Dense) iterativePow(a Matrix, n int) {
	m.Clone(a)
	for i := 1; i < n; i++ {
		m.Mul(m, a)
	}
}

func TestCloneT(t *testing.T) {
	for i, test := range []struct {
		a, want [][]float64
	}{
		{
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		},
		{
			[][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			[][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
		},
		{
			[][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			[][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
		},
		{
			[][]float64{{-1, 0, 0}, {0, -1, 0}, {0, 0, -1}},
			[][]float64{{-1, 0, 0}, {0, -1, 0}, {0, 0, -1}},
		},
		{
			[][]float64{{1, 2, 3}, {4, 5, 6}},
			[][]float64{{1, 4}, {2, 5}, {3, 6}},
		},
	} {
		a := NewDense(flatten(test.a))
		want := NewDense(flatten(test.want))

		var got, gotT Dense

		for j := 0; j < 2; j++ {
			got.Clone(a.T())
			if !Equal(&got, want) {
				t.Errorf("expected transpose for test %d iteration %d: %v transpose = %v",
					i, j, test.a, test.want)
			}
			gotT.Clone(got.T())
			if !Equal(&gotT, a) {
				t.Errorf("expected transpose for test %d iteration %d: %v transpose = %v",
					i, j, test.a, test.want)
			}

			zero(got.mat.Data)
		}
	}
}

func TestCopyT(t *testing.T) {
	for i, test := range []struct {
		a, want [][]float64
	}{
		{
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		},
		{
			[][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			[][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
		},
		{
			[][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			[][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
		},
		{
			[][]float64{{-1, 0, 0}, {0, -1, 0}, {0, 0, -1}},
			[][]float64{{-1, 0, 0}, {0, -1, 0}, {0, 0, -1}},
		},
		{
			[][]float64{{1, 2, 3}, {4, 5, 6}},
			[][]float64{{1, 4}, {2, 5}, {3, 6}},
		},
	} {
		a := NewDense(flatten(test.a))
		want := NewDense(flatten(test.want))

		ar, ac := a.Dims()
		got := NewDense(ac, ar, nil)
		rr := NewDense(ar, ac, nil)

		for j := 0; j < 2; j++ {
			got.Copy(a.T())
			if !Equal(got, want) {
				t.Errorf("expected transpose for test %d iteration %d: %v transpose = %v",
					i, j, test.a, test.want)
			}
			rr.Copy(got.T())
			if !Equal(rr, a) {
				t.Errorf("expected transpose for test %d iteration %d: %v transpose = %v",
					i, j, test.a, test.want)
			}

			zero(got.mat.Data)
		}
	}
}

func TestCopyDenseAlias(t *testing.T) {
	for _, trans := range []bool{false, true} {
		for di := 0; di < 2; di++ {
			for dj := 0; dj < 2; dj++ {
				for si := 0; si < 2; si++ {
					for sj := 0; sj < 2; sj++ {
						a := NewDense(3, 3, []float64{
							1, 2, 3,
							4, 5, 6,
							7, 8, 9,
						})
						src := a.Slice(si, si+2, sj, sj+2)
						want := DenseCopyOf(src)
						got := a.Slice(di, di+2, dj, dj+2).(*Dense)

						if trans {
							panicked, _ := panics(func() { got.Copy(src.T()) })
							if !panicked {
								t.Errorf("expected panic for transpose aliased copy with offsets dst(%d,%d) src(%d,%d):\ngot:\n%v\nwant:\n%v",
									di, dj, si, sj, Formatted(got), Formatted(want),
								)
							}
							continue
						}

						got.Copy(src)
						if !Equal(got, want) {
							t.Errorf("unexpected aliased copy result with offsets dst(%d,%d) src(%d,%d):\ngot:\n%v\nwant:\n%v",
								di, dj, si, sj, Formatted(got), Formatted(want),
							)
						}
					}
				}
			}
		}
	}
}

func TestCopyVectorAlias(t *testing.T) {
	for _, horiz := range []bool{false, true} {
		for do := 0; do < 2; do++ {
			for di := 0; di < 3; di++ {
				for si := 0; si < 3; si++ {
					a := NewDense(3, 3, []float64{
						1, 2, 3,
						4, 5, 6,
						7, 8, 9,
					})
					var src *Vector
					var want *Dense
					if horiz {
						src = a.RowView(si)
						want = DenseCopyOf(a.Slice(si, si+1, 0, 2))
					} else {
						src = a.ColView(si)
						want = DenseCopyOf(a.Slice(0, 2, si, si+1))
					}

					var got *Dense
					if horiz {
						got = a.Slice(di, di+1, do, do+2).(*Dense)
						got.Copy(src.T())
					} else {
						got = a.Slice(do, do+2, di, di+1).(*Dense)
						got.Copy(src)
					}

					if !Equal(got, want) {
						t.Errorf("unexpected aliased copy result with offsets dst(%d) src(%d):\ngot:\n%v\nwant:\n%v",
							di, si, Formatted(got), Formatted(want),
						)
					}
				}
			}
		}
	}
}

func identity(r, c int, v float64) float64 { return v }

func TestApply(t *testing.T) {
	for i, test := range []struct {
		a, want [][]float64
		fn      func(r, c int, v float64) float64
	}{
		{
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			identity,
		},
		{
			[][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			[][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			identity,
		},
		{
			[][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			[][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			identity,
		},
		{
			[][]float64{{-1, 0, 0}, {0, -1, 0}, {0, 0, -1}},
			[][]float64{{-1, 0, 0}, {0, -1, 0}, {0, 0, -1}},
			identity,
		},
		{
			[][]float64{{1, 2, 3}, {4, 5, 6}},
			[][]float64{{1, 2, 3}, {4, 5, 6}},
			identity,
		},
		{
			[][]float64{{1, 2, 3}, {4, 5, 6}},
			[][]float64{{2, 4, 6}, {8, 10, 12}},
			func(r, c int, v float64) float64 { return v * 2 },
		},
		{
			[][]float64{{1, 2, 3}, {4, 5, 6}},
			[][]float64{{0, 2, 0}, {0, 5, 0}},
			func(r, c int, v float64) float64 {
				if c == 1 {
					return v
				}
				return 0
			},
		},
		{
			[][]float64{{1, 2, 3}, {4, 5, 6}},
			[][]float64{{0, 0, 0}, {4, 5, 6}},
			func(r, c int, v float64) float64 {
				if r == 1 {
					return v
				}
				return 0
			},
		},
	} {
		a := NewDense(flatten(test.a))
		want := NewDense(flatten(test.want))

		var got Dense

		for j := 0; j < 2; j++ {
			got.Apply(test.fn, a)
			if !Equal(&got, want) {
				t.Errorf("unexpected result for test %d iteration %d: got: %v want: %v", i, j, got.mat.Data, want.mat.Data)
			}
		}
	}

	for _, fn := range []func(r, c int, v float64) float64{
		identity,
		func(r, c int, v float64) float64 {
			if r < c {
				return v
			}
			return -v
		},
		func(r, c int, v float64) float64 {
			if r%2 == 0 && c%2 == 0 {
				return v
			}
			return -v
		},
		func(_, _ int, v float64) float64 { return v * v },
		func(_, _ int, v float64) float64 { return -v },
	} {
		method := func(receiver, x Matrix) {
			type Applier interface {
				Apply(func(r, c int, v float64) float64, Matrix)
			}
			rd := receiver.(Applier)
			rd.Apply(fn, x)
		}
		denseComparison := func(receiver, x *Dense) {
			receiver.Apply(fn, x)
		}
		testOneInput(t, "Apply", &Dense{}, method, denseComparison, isAnyType, isAnySize, 0)
	}
}

func TestClone(t *testing.T) {
	for i, test := range []struct {
		a    [][]float64
		i, j int
		v    float64
	}{
		{
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			1, 1,
			1,
		},
		{
			[][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			0, 0,
			0,
		},
	} {
		a := NewDense(flatten(test.a))
		b := *a
		a.Clone(a)
		a.Set(test.i, test.j, test.v)

		if Equal(&b, a) {
			t.Errorf("unexpected mirror of write to cloned matrix for test %d: %v cloned and altered = %v",
				i, a, &b)
		}
	}
}

// TODO(kortschak) Roll this into testOneInput when it exists.
func TestCopyPanic(t *testing.T) {
	for _, a := range []*Dense{
		{},
		{mat: blas64.General{Rows: 1}},
		{mat: blas64.General{Cols: 1}},
	} {
		var rows, cols int
		m := NewDense(1, 1, nil)
		panicked, message := panics(func() { rows, cols = m.Copy(a) })
		if panicked {
			t.Errorf("unexpected panic: %v", message)
		}
		if rows != 0 {
			t.Errorf("unexpected rows: got: %d want: 0", rows)
		}
		if cols != 0 {
			t.Errorf("unexpected cols: got: %d want: 0", cols)
		}
	}
}

func TestStack(t *testing.T) {
	for i, test := range []struct {
		a, b, e [][]float64
	}{
		{
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		},
		{
			[][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			[][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			[][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
		},
		{
			[][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			[][]float64{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}},
			[][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}, {0, 1, 0}, {0, 0, 1}, {1, 0, 0}},
		},
	} {
		a := NewDense(flatten(test.a))
		b := NewDense(flatten(test.b))

		var s Dense
		s.Stack(a, b)

		if !Equal(&s, NewDense(flatten(test.e))) {
			t.Errorf("unexpected result for Stack test %d: %v stack %v = %v", i, a, b, s)
		}
	}

	method := func(receiver, a, b Matrix) {
		type Stacker interface {
			Stack(a, b Matrix)
		}
		rd := receiver.(Stacker)
		rd.Stack(a, b)
	}
	denseComparison := func(receiver, a, b *Dense) {
		receiver.Stack(a, b)
	}
	testTwoInput(t, "Stack", &Dense{}, method, denseComparison, legalTypesAll, legalSizeSameWidth, 0)
}

func TestAugment(t *testing.T) {
	for i, test := range []struct {
		a, b, e [][]float64
	}{
		{
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			[][]float64{{0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0}},
		},
		{
			[][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			[][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			[][]float64{{1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}},
		},
		{
			[][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			[][]float64{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}},
			[][]float64{{1, 0, 0, 0, 1, 0}, {0, 1, 0, 0, 0, 1}, {0, 0, 1, 1, 0, 0}},
		},
	} {
		a := NewDense(flatten(test.a))
		b := NewDense(flatten(test.b))

		var s Dense
		s.Augment(a, b)

		if !Equal(&s, NewDense(flatten(test.e))) {
			t.Errorf("unexpected result for Augment test %d: %v augment %v = %v", i, a, b, s)
		}
	}

	method := func(receiver, a, b Matrix) {
		type Augmenter interface {
			Augment(a, b Matrix)
		}
		rd := receiver.(Augmenter)
		rd.Augment(a, b)
	}
	denseComparison := func(receiver, a, b *Dense) {
		receiver.Augment(a, b)
	}
	testTwoInput(t, "Augment", &Dense{}, method, denseComparison, legalTypesAll, legalSizeSameHeight, 0)
}

func TestRankOne(t *testing.T) {
	for i, test := range []struct {
		x     []float64
		y     []float64
		m     [][]float64
		alpha float64
	}{
		{
			x:     []float64{5},
			y:     []float64{10},
			m:     [][]float64{{2}},
			alpha: -3,
		},
		{
			x:     []float64{5, 6, 1},
			y:     []float64{10},
			m:     [][]float64{{2}, {-3}, {5}},
			alpha: -3,
		},

		{
			x:     []float64{5},
			y:     []float64{10, 15, 8},
			m:     [][]float64{{2, -3, 5}},
			alpha: -3,
		},
		{
			x: []float64{1, 5},
			y: []float64{10, 15},
			m: [][]float64{
				{2, -3},
				{4, -1},
			},
			alpha: -3,
		},
		{
			x: []float64{2, 3, 9},
			y: []float64{8, 9},
			m: [][]float64{
				{2, 3},
				{4, 5},
				{6, 7},
			},
			alpha: -3,
		},
		{
			x: []float64{2, 3},
			y: []float64{8, 9, 9},
			m: [][]float64{
				{2, 3, 6},
				{4, 5, 7},
			},
			alpha: -3,
		},
	} {
		want := &Dense{}
		xm := NewDense(len(test.x), 1, test.x)
		ym := NewDense(1, len(test.y), test.y)

		want.Mul(xm, ym)
		want.Scale(test.alpha, want)
		want.Add(want, NewDense(flatten(test.m)))

		a := NewDense(flatten(test.m))
		m := &Dense{}
		// Check with a new matrix
		m.RankOne(a, test.alpha, NewVector(len(test.x), test.x), NewVector(len(test.y), test.y))
		if !Equal(m, want) {
			t.Errorf("unexpected result for RankOne test %d iteration 0: got: %+v want: %+v", i, m, want)
		}
		// Check with the same matrix
		a.RankOne(a, test.alpha, NewVector(len(test.x), test.x), NewVector(len(test.y), test.y))
		if !Equal(a, want) {
			t.Errorf("unexpected result for Outer test %d iteration 1: got: %+v want: %+v", i, m, want)
		}
	}
}

func TestOuter(t *testing.T) {
	for i, test := range []struct {
		x []float64
		y []float64
	}{
		{
			x: []float64{5},
			y: []float64{10},
		},
		{
			x: []float64{5, 6, 1},
			y: []float64{10},
		},

		{
			x: []float64{5},
			y: []float64{10, 15, 8},
		},
		{
			x: []float64{1, 5},
			y: []float64{10, 15},
		},
		{
			x: []float64{2, 3, 9},
			y: []float64{8, 9},
		},
		{
			x: []float64{2, 3},
			y: []float64{8, 9, 9},
		},
	} {
		for _, f := range []float64{0.5, 1, 3} {
			want := &Dense{}
			xm := NewDense(len(test.x), 1, test.x)
			ym := NewDense(1, len(test.y), test.y)

			want.Mul(xm, ym)
			want.Scale(f, want)

			var m Dense
			for j := 0; j < 2; j++ {
				// Check with a new matrix - and then again.
				m.Outer(f, NewVector(len(test.x), test.x), NewVector(len(test.y), test.y))
				if !Equal(&m, want) {
					t.Errorf("unexpected result for Outer test %d iteration %d scale %v: got: %+v want: %+v", i, j, f, m, want)
				}
			}
		}
	}
}

func TestInverse(t *testing.T) {
	for i, test := range []struct {
		a    Matrix
		want Matrix // nil indicates that a is singular.
		tol  float64
	}{
		{
			a: NewDense(3, 3, []float64{
				8, 1, 6,
				3, 5, 7,
				4, 9, 2,
			}),
			want: NewDense(3, 3, []float64{
				0.147222222222222, -0.144444444444444, 0.063888888888889,
				-0.061111111111111, 0.022222222222222, 0.105555555555556,
				-0.019444444444444, 0.188888888888889, -0.102777777777778,
			}),
			tol: 1e-14,
		},
		{
			a: NewDense(3, 3, []float64{
				8, 1, 6,
				3, 5, 7,
				4, 9, 2,
			}).T(),
			want: NewDense(3, 3, []float64{
				0.147222222222222, -0.144444444444444, 0.063888888888889,
				-0.061111111111111, 0.022222222222222, 0.105555555555556,
				-0.019444444444444, 0.188888888888889, -0.102777777777778,
			}).T(),
			tol: 1e-14,
		},

		// This case does not fail, but we do not guarantee that. The success
		// is because the receiver and the input are aligned in the call to
		// inverse. If there was a misalignment, the result would likely be
		// incorrect and no shadowing panic would occur.
		{
			a: asBasicMatrix(NewDense(3, 3, []float64{
				8, 1, 6,
				3, 5, 7,
				4, 9, 2,
			})),
			want: NewDense(3, 3, []float64{
				0.147222222222222, -0.144444444444444, 0.063888888888889,
				-0.061111111111111, 0.022222222222222, 0.105555555555556,
				-0.019444444444444, 0.188888888888889, -0.102777777777778,
			}),
			tol: 1e-14,
		},

		// The following case fails as it does not follow the shadowing rules.
		// Specifically, the test extracts the underlying *Dense, and uses
		// it as a receiver with the basicMatrix as input. The basicMatrix type
		// allows shadowing of the input data without providing the Raw method
		// required for detection of shadowing.
		//
		// We specifically state we do not check this case.
		//
		// {
		// 	a: asBasicMatrix(NewDense(3, 3, []float64{
		// 		8, 1, 6,
		// 		3, 5, 7,
		// 		4, 9, 2,
		// 	})).T(),
		// 	want: NewDense(3, 3, []float64{
		// 		0.147222222222222, -0.144444444444444, 0.063888888888889,
		// 		-0.061111111111111, 0.022222222222222, 0.105555555555556,
		// 		-0.019444444444444, 0.188888888888889, -0.102777777777778,
		// 	}).T(),
		// 	tol: 1e-14,
		// },

		{
			a: NewDense(4, 4, []float64{
				5, 2, 8, 7,
				4, 5, 8, 2,
				8, 5, 3, 2,
				8, 7, 7, 5,
			}),
			want: NewDense(4, 4, []float64{
				0.100548446069470, 0.021937842778793, 0.334552102376599, -0.283363802559415,
				-0.226691042047532, -0.067641681901280, -0.281535648994515, 0.457038391224863,
				0.080438756855576, 0.217550274223035, 0.067641681901280, -0.226691042047532,
				0.043875685557587, -0.244972577696527, -0.235831809872029, 0.330895795246801,
			}),
			tol: 1e-14,
		},

		// Tests with singular matrix.
		{
			a: NewDense(1, 1, []float64{
				0,
			}),
		},
		{
			a: NewDense(2, 2, []float64{
				0, 0,
				0, 0,
			}),
		},
		{
			a: NewDense(2, 2, []float64{
				0, 0,
				0, 1,
			}),
		},
		{
			a: NewDense(3, 3, []float64{
				0, 0, 0,
				0, 0, 0,
				0, 0, 0,
			}),
		},
		{
			a: NewDense(4, 4, []float64{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			}),
		},
		{
			a: NewDense(4, 4, []float64{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 20, 20,
				0, 0, 20, 20,
			}),
		},
		{
			a: NewDense(4, 4, []float64{
				0, 1, 0, 0,
				0, 0, 1, 0,
				0, 0, 0, 1,
				0, 0, 0, 0,
			}),
		},
		{
			a: NewDense(4, 4, []float64{
				1, 1, 1, 1,
				1, 1, 1, 1,
				1, 1, 1, 1,
				1, 1, 1, 1,
			}),
		},
		{
			a: NewDense(5, 5, []float64{
				0, 1, 0, 0, 0,
				4, 0, 2, 0, 0,
				0, 3, 0, 3, 0,
				0, 0, 2, 0, 4,
				0, 0, 0, 1, 0,
			}),
		},
		{
			a: NewDense(5, 5, []float64{
				4, -1, -1, -1, -1,
				-1, 4, -1, -1, -1,
				-1, -1, 4, -1, -1,
				-1, -1, -1, 4, -1,
				-1, -1, -1, -1, 4,
			}),
		},
		{
			a: NewDense(5, 5, []float64{
				2, -1, 0, 0, -1,
				-1, 2, -1, 0, 0,
				0, -1, 2, -1, 0,
				0, 0, -1, 2, -1,
				-1, 0, 0, -1, 2,
			}),
		},
		{
			a: NewDense(5, 5, []float64{
				1, 2, 3, 5, 8,
				2, 3, 5, 8, 13,
				3, 5, 8, 13, 21,
				5, 8, 13, 21, 34,
				8, 13, 21, 34, 55,
			}),
		},
		{
			a: NewDense(8, 8, []float64{
				611, 196, -192, 407, -8, -52, -49, 29,
				196, 899, 113, -192, -71, -43, -8, -44,
				-192, 113, 899, 196, 61, 49, 8, 52,
				407, -192, 196, 611, 8, 44, 59, -23,
				-8, -71, 61, 8, 411, -599, 208, 208,
				-52, -43, 49, 44, -599, 411, 208, 208,
				-49, -8, 8, 59, 208, 208, 99, -911,
				29, -44, 52, -23, 208, 208, -911, 99,
			}),
		},
	} {
		var got Dense
		err := got.Inverse(test.a)
		if test.want == nil {
			if err == nil {
				t.Errorf("Case %d: expected error for singular matrix", i)
			}
			continue
		}
		if err != nil {
			t.Errorf("Case %d: unexpected error: %v", i, err)
			continue
		}
		if !equalApprox(&got, test.want, test.tol, false) {
			t.Errorf("Case %d, inverse mismatch.", i)
		}
		var m Dense
		m.Mul(&got, test.a)
		r, _ := test.a.Dims()
		d := make([]float64, r*r)
		for i := 0; i < r*r; i += r + 1 {
			d[i] = 1
		}
		eye := NewDense(r, r, d)
		if !equalApprox(eye, &m, 1e-14, false) {
			t.Errorf("Case %d, A^-1 * A != I", i)
		}

		var tmp Dense
		tmp.Clone(test.a)
		aU, transposed := untranspose(test.a)
		if transposed {
			switch aU := aU.(type) {
			case *Dense:
				err = aU.Inverse(test.a)
			case *basicMatrix:
				err = (*Dense)(aU).Inverse(test.a)
			default:
				continue
			}
			m.Mul(aU, &tmp)
		} else {
			switch a := test.a.(type) {
			case *Dense:
				err = a.Inverse(test.a)
				m.Mul(a, &tmp)
			case *basicMatrix:
				err = (*Dense)(a).Inverse(test.a)
				m.Mul(a, &tmp)
			default:
				continue
			}
		}
		if err != nil {
			t.Errorf("Error computing inverse: %v", err)
		}
		if !equalApprox(eye, &m, 1e-14, false) {
			t.Errorf("Case %d, A^-1 * A != I", i)
			fmt.Println(Formatted(&m))
		}
	}
}

var (
	wd *Dense
)

func BenchmarkMulDense100Half(b *testing.B)        { denseMulBench(b, 100, 0.5) }
func BenchmarkMulDense100Tenth(b *testing.B)       { denseMulBench(b, 100, 0.1) }
func BenchmarkMulDense1000Half(b *testing.B)       { denseMulBench(b, 1000, 0.5) }
func BenchmarkMulDense1000Tenth(b *testing.B)      { denseMulBench(b, 1000, 0.1) }
func BenchmarkMulDense1000Hundredth(b *testing.B)  { denseMulBench(b, 1000, 0.01) }
func BenchmarkMulDense1000Thousandth(b *testing.B) { denseMulBench(b, 1000, 0.001) }
func denseMulBench(b *testing.B, size int, rho float64) {
	b.StopTimer()
	a, _ := randDense(size, rho, rand.NormFloat64)
	d, _ := randDense(size, rho, rand.NormFloat64)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var n Dense
		n.Mul(a, d)
		wd = &n
	}
}

func BenchmarkPreMulDense100Half(b *testing.B)        { densePreMulBench(b, 100, 0.5) }
func BenchmarkPreMulDense100Tenth(b *testing.B)       { densePreMulBench(b, 100, 0.1) }
func BenchmarkPreMulDense1000Half(b *testing.B)       { densePreMulBench(b, 1000, 0.5) }
func BenchmarkPreMulDense1000Tenth(b *testing.B)      { densePreMulBench(b, 1000, 0.1) }
func BenchmarkPreMulDense1000Hundredth(b *testing.B)  { densePreMulBench(b, 1000, 0.01) }
func BenchmarkPreMulDense1000Thousandth(b *testing.B) { densePreMulBench(b, 1000, 0.001) }
func densePreMulBench(b *testing.B, size int, rho float64) {
	b.StopTimer()
	a, _ := randDense(size, rho, rand.NormFloat64)
	d, _ := randDense(size, rho, rand.NormFloat64)
	wd = NewDense(size, size, nil)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		wd.Mul(a, d)
	}
}

func BenchmarkRow10(b *testing.B)   { rowBench(b, 10) }
func BenchmarkRow100(b *testing.B)  { rowBench(b, 100) }
func BenchmarkRow1000(b *testing.B) { rowBench(b, 1000) }

func rowBench(b *testing.B, size int) {
	a, _ := randDense(size, 1, rand.NormFloat64)
	_, c := a.Dims()
	dst := make([]float64, c)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Row(dst, 0, a)
	}
}

func BenchmarkExp10(b *testing.B)   { expBench(b, 10) }
func BenchmarkExp100(b *testing.B)  { expBench(b, 100) }
func BenchmarkExp1000(b *testing.B) { expBench(b, 1000) }

func expBench(b *testing.B, size int) {
	a, _ := randDense(size, 1, rand.NormFloat64)

	b.ResetTimer()
	var m Dense
	for i := 0; i < b.N; i++ {
		m.Exp(a)
	}
}

func BenchmarkPow10_3(b *testing.B)   { powBench(b, 10, 3) }
func BenchmarkPow100_3(b *testing.B)  { powBench(b, 100, 3) }
func BenchmarkPow1000_3(b *testing.B) { powBench(b, 1000, 3) }
func BenchmarkPow10_4(b *testing.B)   { powBench(b, 10, 4) }
func BenchmarkPow100_4(b *testing.B)  { powBench(b, 100, 4) }
func BenchmarkPow1000_4(b *testing.B) { powBench(b, 1000, 4) }
func BenchmarkPow10_5(b *testing.B)   { powBench(b, 10, 5) }
func BenchmarkPow100_5(b *testing.B)  { powBench(b, 100, 5) }
func BenchmarkPow1000_5(b *testing.B) { powBench(b, 1000, 5) }
func BenchmarkPow10_6(b *testing.B)   { powBench(b, 10, 6) }
func BenchmarkPow100_6(b *testing.B)  { powBench(b, 100, 6) }
func BenchmarkPow1000_6(b *testing.B) { powBench(b, 1000, 6) }
func BenchmarkPow10_7(b *testing.B)   { powBench(b, 10, 7) }
func BenchmarkPow100_7(b *testing.B)  { powBench(b, 100, 7) }
func BenchmarkPow1000_7(b *testing.B) { powBench(b, 1000, 7) }
func BenchmarkPow10_8(b *testing.B)   { powBench(b, 10, 8) }
func BenchmarkPow100_8(b *testing.B)  { powBench(b, 100, 8) }
func BenchmarkPow1000_8(b *testing.B) { powBench(b, 1000, 8) }
func BenchmarkPow10_9(b *testing.B)   { powBench(b, 10, 9) }
func BenchmarkPow100_9(b *testing.B)  { powBench(b, 100, 9) }
func BenchmarkPow1000_9(b *testing.B) { powBench(b, 1000, 9) }

func powBench(b *testing.B, size, n int) {
	a, _ := randDense(size, 1, rand.NormFloat64)

	b.ResetTimer()
	var m Dense
	for i := 0; i < b.N; i++ {
		m.Pow(a, n)
	}
}

func BenchmarkMulTransDense100Half(b *testing.B)        { denseMulTransBench(b, 100, 0.5) }
func BenchmarkMulTransDense100Tenth(b *testing.B)       { denseMulTransBench(b, 100, 0.1) }
func BenchmarkMulTransDense1000Half(b *testing.B)       { denseMulTransBench(b, 1000, 0.5) }
func BenchmarkMulTransDense1000Tenth(b *testing.B)      { denseMulTransBench(b, 1000, 0.1) }
func BenchmarkMulTransDense1000Hundredth(b *testing.B)  { denseMulTransBench(b, 1000, 0.01) }
func BenchmarkMulTransDense1000Thousandth(b *testing.B) { denseMulTransBench(b, 1000, 0.001) }
func denseMulTransBench(b *testing.B, size int, rho float64) {
	b.StopTimer()
	a, _ := randDense(size, rho, rand.NormFloat64)
	d, _ := randDense(size, rho, rand.NormFloat64)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var n Dense
		n.Mul(a, d.T())
		wd = &n
	}
}

func BenchmarkMulTransDenseSym100Half(b *testing.B)        { denseMulTransSymBench(b, 100, 0.5) }
func BenchmarkMulTransDenseSym100Tenth(b *testing.B)       { denseMulTransSymBench(b, 100, 0.1) }
func BenchmarkMulTransDenseSym1000Half(b *testing.B)       { denseMulTransSymBench(b, 1000, 0.5) }
func BenchmarkMulTransDenseSym1000Tenth(b *testing.B)      { denseMulTransSymBench(b, 1000, 0.1) }
func BenchmarkMulTransDenseSym1000Hundredth(b *testing.B)  { denseMulTransSymBench(b, 1000, 0.01) }
func BenchmarkMulTransDenseSym1000Thousandth(b *testing.B) { denseMulTransSymBench(b, 1000, 0.001) }
func denseMulTransSymBench(b *testing.B, size int, rho float64) {
	b.StopTimer()
	a, _ := randDense(size, rho, rand.NormFloat64)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var n Dense
		n.Mul(a, a.T())
		wd = &n
	}
}

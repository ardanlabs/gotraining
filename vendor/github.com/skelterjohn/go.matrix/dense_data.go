// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

//returns a copy of the row (not a slice)
func (A *DenseMatrix) RowCopy(i int) []float64 {
	row := make([]float64, A.cols)
	for j := 0; j < A.cols; j++ {
		row[j] = A.Get(i, j)
	}
	return row
}

//returns a copy of the column (not a slice)
func (A *DenseMatrix) ColCopy(j int) []float64 {
	col := make([]float64, A.rows)
	for i := 0; i < A.rows; i++ {
		col[i] = A.Get(i, j)
	}
	return col
}

//returns a copy of the diagonal (not a slice)
func (A *DenseMatrix) DiagonalCopy() []float64 {
	span := A.rows
	if A.cols < span {
		span = A.cols
	}
	diag := make([]float64, span)
	for i := 0; i < span; i++ {
		diag[i] = A.Get(i, i)
	}
	return diag
}

func (A *DenseMatrix) BufferRow(i int, buf []float64) {
	for j := 0; j < A.cols; j++ {
		buf[j] = A.Get(i, j)
	}
}

func (A *DenseMatrix) BufferCol(j int, buf []float64) {
	for i := 0; i < A.rows; i++ {
		buf[i] = A.Get(i, j)
	}
}

func (A *DenseMatrix) BufferDiagonal(buf []float64) {
	for i := 0; i < A.rows && i < A.cols; i++ {
		buf[i] = A.Get(i, i)
	}
}

func (A *DenseMatrix) FillRow(i int, buf []float64) {
	for j := 0; j < A.cols; j++ {
		A.Set(i, j, buf[j])
	}
}

func (A *DenseMatrix) FillCol(j int, buf []float64) {
	for i := 0; i < A.rows; i++ {
		A.Set(i, j, buf[i])
	}
}

func (A *DenseMatrix) FillDiagonal(buf []float64) {
	for i := 0; i < A.rows && i < A.cols; i++ {
		A.Set(i, i, buf[i])
	}
}

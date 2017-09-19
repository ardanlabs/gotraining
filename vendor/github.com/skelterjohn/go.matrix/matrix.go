// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//target:gomatrix.googlecode.com/hg/matrix

//Linear algebra.
package matrix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//The MatrixRO interface defines matrix operations that do not change the
//underlying data, such as information requests or the creation of transforms
/*
Read-only matrix types (at the moment, PivotMatrix).
*/
type MatrixRO interface {
	//Returns true if the underlying object is nil.
	Nil() bool

	//The number of rows in this matrix.
	Rows() int
	//The number of columns in this matrix.
	Cols() int

	//The number of elements in this matrix.
	NumElements() int
	//The size pair, (Rows(), Cols())
	GetSize() (int, int)

	//The element in the ith row and jth column.
	Get(i, j int) float64

	Plus(MatrixRO) (Matrix, error)
	Minus(MatrixRO) (Matrix, error)
	Times(MatrixRO) (Matrix, error)

	//The determinant of this matrix.
	Det() float64
	//The trace of this matrix.
	Trace() float64

	//A pretty-print string.
	String() string

	DenseMatrix() *DenseMatrix
	SparseMatrix() *SparseMatrix
}

/*
A mutable matrix.
*/
type Matrix interface {
	MatrixRO

	//Set the element at the ith row and jth column to v.
	Set(i int, j int, v float64)

	Add(MatrixRO) error
	Subtract(MatrixRO) error
	Scale(float64)
}

type matrix struct {
	rows int
	cols int
}

func (A *matrix) Nil() bool { return A == nil }

func (A *matrix) Rows() int { return A.rows }

func (A *matrix) Cols() int { return A.cols }

func (A *matrix) NumElements() int { return A.rows * A.cols }

func (A *matrix) GetSize() (rows, cols int) {
	rows = A.rows
	cols = A.cols
	return
}

/*
	Take a matlab-style matrix representation

	eg [a b c; d e f]
*/
func ParseMatlab(txt string) (A *DenseMatrix, err error) {
	var arrays [][]float64

	spaceSep := strings.Fields(txt)

	tok := func() (t string, eos bool) {
		defer func() {
			for len(spaceSep) != 0 && len(spaceSep[0]) == 0 {
				spaceSep = spaceSep[1:]
			}
		}()

		isNotNumber := func(c byte) bool {
			return c != '[' || c != ']' || c == ';'
		}

		if len(spaceSep) == 0 {
			eos = true
			return
		}

		top := spaceSep[0]

		var lof int
		for ; lof < len(top) && !isNotNumber(top[lof]); lof++ {
		}

		if lof != 0 {
			t = top[:lof]
			spaceSep[0] = top[lof:]
			return
		} else {
			t = top[:1]
			spaceSep[0] = top[1:]
			return
		}

		panic("unreachable")
	}

	stack := func(row []float64) (err error) {
		if len(arrays) == 0 {
			arrays = [][]float64{row}
			return
		}
		if len(arrays[0]) != len(row) {
			err = errors.New("misaligned row")
		}
		arrays = append(arrays, row)
		return
	}

	var row []float64

loop:
	for {
		t, eos := tok()
		if eos {
			break loop
		}
		switch t {
		case "[":
		case ";":
			err = stack(row)
			if err != nil {
				return
			}
			row = []float64{}
		case "]":
			err = stack(row)
			if err != nil {
				return
			}
			break loop
		default:
			var v float64
			v, err = strconv.ParseFloat(t, 64)
			if err != nil {
				return
			}
			row = append(row, v)
		}
	}
	A = MakeDenseMatrixStacked(arrays)
	return
}

func String(A MatrixRO) string {
	condense := func(vs string) string {
		if strings.Index(vs, ".") != -1 {
			for vs[len(vs)-1] == '0' {
				vs = vs[0 : len(vs)-1]
			}
		}
		if vs[len(vs)-1] == '.' {
			vs = vs[0 : len(vs)-1]
		}
		return vs
	}

	if A == nil {
		return "{nil}"
	}
	s := "{"

	maxLen := 0
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			v := A.Get(i, j)
			vs := condense(fmt.Sprintf("%f", v))

			maxLen = maxInt(maxLen, len(vs))
		}
	}

	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			v := A.Get(i, j)

			vs := condense(fmt.Sprintf("%f", v))

			for len(vs) < maxLen {
				vs = " " + vs
			}
			s += vs
			if i != A.Rows()-1 || j != A.Cols()-1 {
				s += ","
			}
			if j != A.Cols()-1 {
				s += " "
			}
		}
		if i != A.Rows()-1 {
			s += "\n "
		}
	}
	s += "}"
	return s
}

// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

/*
A space-optimized structure for pivot matrices, ie a matrix with
exactly one 1 in each row and each column.
*/
type PivotMatrix struct {
	matrix
	pivots    []int
	pivotSign float64
}

func (P *PivotMatrix) Get(i, j int) float64 {
	i = i % P.rows
	if i < 0 {
		i = P.rows - i
	}
	j = j % P.cols
	if j < 0 {
		j = P.cols - j
	}
	if P.pivots[j] == i {
		return 1
	}
	return 0
}

/*
Convert this PivotMatrix into a DenseMatrix.
*/
func (P *PivotMatrix) DenseMatrix() *DenseMatrix {
	A := Zeros(P.rows, P.cols)
	for j := 0; j < P.rows; j++ {
		A.Set(P.pivots[j], j, 1)
	}
	return A
}

/*
Convert this PivotMatrix into a SparseMatrix.
*/
func (P *PivotMatrix) SparseMatrix() *SparseMatrix {
	A := ZerosSparse(P.rows, P.cols)
	for j := 0; j < P.rows; j++ {
		A.Set(P.pivots[j], j, 1)
	}
	return A
}

/*
Make a copy of this PivotMatrix.
*/
func (P *PivotMatrix) Copy() *PivotMatrix { return MakePivotMatrix(P.pivots, P.pivotSign) }

func MakePivotMatrix(pivots []int, pivotSign float64) *PivotMatrix {
	n := len(pivots)
	P := new(PivotMatrix)
	P.rows = n
	P.cols = n
	P.pivots = pivots
	P.pivotSign = pivotSign
	return P
}

func (A *PivotMatrix) String() string { return String(A) }

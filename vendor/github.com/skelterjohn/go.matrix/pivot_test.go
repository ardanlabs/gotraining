// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import "testing"
import "fmt"

func TestTimes_Pivot(t *testing.T) {
	P1 := MakePivotMatrix([]int{2, 1, 0}, -1)
	P2 := MakePivotMatrix([]int{2, 0, 1}, 1)
	P, _ := P1.TimesPivot(P2)
	if !Equals(P, Product(P1, P2)) {
		if verbose {
			fmt.Printf("%v\n%v\n%v\n", P1, P2, P)
		}
		t.Fail()
	}
}

func TestRowPivot(t *testing.T) {
	P := MakePivotMatrix([]int{2, 1, 0}, -1)
	A := Normals(3, 4)
	B, _ := P.RowPivotDense(A)
	Btrue := Product(P, A)
	if !Equals(B, Btrue) {
		t.Fail()
	}
	A = Normals(4, 3)
	_, err := P.RowPivotDense(A)
	if err == nil {
		t.Fail()
	}
	C := Normals(3, 4).SparseMatrix()
	D, _ := P.RowPivotSparse(C)
	Btrue = Product(P, C)
	if !Equals(D, Btrue) {
		t.Fail()
	}
}

func TestColPivot(t *testing.T) {
	P := MakePivotMatrix([]int{2, 1, 0}, -1)
	A := Normals(4, 3)
	B, _ := P.ColPivotDense(A)
	Btrue := Product(A, P)
	if !Equals(B, Btrue) {
		t.Fail()
	}
	A = Normals(3, 4)
	_, err := P.ColPivotDense(A)
	if err == nil {
		t.Fail()
	}
	C := Normals(4, 3).SparseMatrix()
	D, _ := P.ColPivotSparse(C)
	Btrue = Product(C, P)
	if !Equals(D, Btrue) {
		t.Fail()
	}
}

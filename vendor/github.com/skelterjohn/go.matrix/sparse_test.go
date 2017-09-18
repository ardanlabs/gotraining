// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestAdd_Sparse(t *testing.T) {
	A := NormalsSparse(3, 3, 9)
	B := NormalsSparse(3, 3, 9)
	C1, _ := A.Plus(B)
	C2, _ := A.PlusSparse(B)
	if !ApproxEquals(C1, Sum(A, B), ε) {
		t.Fail()
	}
	if !ApproxEquals(C2, Sum(A, B), ε) {
		t.Fail()
	}
}

func TestSubtract_Sparse(t *testing.T) {
	A := NormalsSparse(3, 3, 9)
	B := NormalsSparse(3, 3, 9)
	C1, _ := A.Minus(B)
	C2, _ := A.MinusSparse(B)
	if !ApproxEquals(C1, Difference(A, B), ε) {
		t.Fail()
	}
	if !ApproxEquals(C2, Difference(A, B), ε) {
		t.Fail()
	}
}

func TestTimes_Sparse(t *testing.T) {
	A := Normals(3, 3).SparseMatrix()
	B := Normals(3, 3).SparseMatrix()
	C1, _ := A.Times(B)
	C2, _ := A.TimesSparse(B)
	if !ApproxEquals(C1, Product(A, B), ε) {
		t.Fail()
	}
	if !ApproxEquals(C2, Product(A, B), ε) {
		t.Fail()
	}
}

func TestElementMult_Sparse(t *testing.T) {
	A := Normals(3, 3).SparseMatrix()
	B := Normals(3, 3).SparseMatrix()
	C1, _ := A.ElementMult(B)
	C2, _ := A.ElementMultSparse(B)
	D, _ := A.DenseMatrix().ElementMult(B)
	if !Equals(D, C1) {
		t.Fail()
	}
	if !Equals(D, C2) {
		t.Fail()
	}
}

func TestGetMatrix_Sparse(t *testing.T) {
	A := ZerosSparse(6, 6)
	for i := 0; i < 36; i++ {
		x := rand.Intn(6)
		y := rand.Intn(6)
		A.Set(y, x, 1)
	}
	B := A.GetMatrix(1, 1, 4, 4)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if B.Get(i, j) != A.Get(i+1, j+1) {
				t.Fail()
			}
		}
	}

}

func TestAugment_Sparse(t *testing.T) {
	var A, B, C *SparseMatrix
	A = NormalsSparse(4, 4, 16)
	B = NormalsSparse(4, 4, 16)
	C, _ = A.Augment(B)
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if C.Get(i, j) != A.Get(i, j) {
				t.Fail()
			}
		}
	}
	for i := 0; i < B.Rows(); i++ {
		for j := 0; j < B.Cols(); j++ {
			if C.Get(i, j+A.Cols()) != B.Get(i, j) {
				t.Fail()
			}
		}
	}

	A = NormalsSparse(2, 2, 4)
	B = NormalsSparse(4, 4, 16)
	C, err := A.Augment(B)
	if err == nil {
		t.Fail()
	}

	A = NormalsSparse(4, 4, 16)
	B = NormalsSparse(4, 2, 8)
	C, _ = A.Augment(B)
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if C.Get(i, j) != A.Get(i, j) {
				t.Fail()
			}
		}
	}
	for i := 0; i < B.Rows(); i++ {
		for j := 0; j < B.Cols(); j++ {
			if C.Get(i, j+A.Cols()) != B.Get(i, j) {
				t.Fail()
			}
		}
	}
}

func TestStack_Sparse(t *testing.T) {
	var A, B, C *SparseMatrix
	A = NormalsSparse(4, 4, 16)
	B = NormalsSparse(4, 4, 16)
	C, _ = A.Stack(B)
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if C.Get(i, j) != A.Get(i, j) {
				t.Fail()
			}
		}
	}
	for i := 0; i < B.Rows(); i++ {
		for j := 0; j < B.Cols(); j++ {
			if C.Get(i+A.Rows(), j) != B.Get(i, j) {
				t.Fail()
			}
		}
	}

	A = NormalsSparse(2, 2, 4)
	B = NormalsSparse(4, 4, 16)
	C, err := A.Stack(B)
	if err == nil {
		if verbose {
			fmt.Printf("%v\n", err)
		}
		t.Fail()
	}

	A = NormalsSparse(4, 4, 16)
	B = NormalsSparse(2, 4, 8)
	C, _ = A.Stack(B)
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if C.Get(i, j) != A.Get(i, j) {
				t.Fail()
			}
		}
	}
	for i := 0; i < B.Rows(); i++ {
		for j := 0; j < B.Cols(); j++ {
			if C.Get(i+A.Rows(), j) != B.Get(i, j) {
				t.Fail()
			}
		}
	}
}

// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import (
	"math"
	"errors"
)

func (A *DenseMatrix) Symmetric() bool {
	if A.rows != A.cols {
		return false
	}
	for i := 0; i < A.rows; i++ {
		for j := 0; j < i; j++ {
			if A.Get(i, j) != A.Get(j, i) {
				return false
			}
		}
	}
	return true
}

func (m *DenseMatrix) SwapRows(r1 int, r2 int) {
	index1 := r1 * m.step
	index2 := r2 * m.step
	for j := 0; j < m.cols; j++ {
		m.elements[index1], m.elements[index2] = m.elements[index2], m.elements[index1]
		index1++
		index2++
	}
}

func (m *DenseMatrix) ScaleRow(r int, f float64) {
	index := r * m.step
	for j := 0; j < m.cols; j++ {
		m.elements[index] *= f
		index++
	}
}

func (m *DenseMatrix) ScaleAddRow(rd int, rs int, f float64) {
	indexd := rd * m.step
	indexs := rs * m.step
	for j := 0; j < m.cols; j++ {
		m.elements[indexd] += f * m.elements[indexs]
		indexd++
		indexs++
	}
}

func (A *DenseMatrix) Inverse() (*DenseMatrix, error) {
	if A.Rows() != A.Cols() {
		return nil, ErrorDimensionMismatch
	}
	aug, _ := A.Augment(Eye(A.Rows()))
	for i := 0; i < aug.Rows(); i++ {
		j := i
		for k := i; k < aug.Rows(); k++ {
			if math.Abs(aug.Get(k, i)) > math.Abs(aug.Get(j, i)) {
				j = k
			}
		}
		if j != i {
			aug.SwapRows(i, j)
		}
		if aug.Get(i, i) == 0 {
			return nil, ExceptionSingular
		}
		aug.ScaleRow(i, 1.0/aug.Get(i, i))
		for k := 0; k < aug.Rows(); k++ {
			if k == i {
				continue
			}
			aug.ScaleAddRow(k, i, -aug.Get(k, i))
		}
	}
	inv := aug.GetMatrix(0, A.Cols(), A.Rows(), A.Cols())
	return inv, nil
}

func (A *DenseMatrix) Det() float64 {
	B := A.Copy()
	P := B.LUInPlace()
	return product(B.DiagonalCopy()) * P.Det()
}

func (A *DenseMatrix) Trace() float64 { return sum(A.DiagonalCopy()) }

func (A *DenseMatrix) OneNorm() (ε float64) {
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			ε = max(ε, A.Get(i, j))
		}
	}
	return
}

func (A *DenseMatrix) TwoNorm() float64 {
	var sum float64 = 0
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			v := A.elements[i*A.step+j]
			sum += v * v
		}
	}
	return math.Sqrt(sum)
}

func (A *DenseMatrix) InfinityNorm() (ε float64) {
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			ε += A.Get(i, j)
		}
	}
	return
}

func (A *DenseMatrix) Transpose() *DenseMatrix {
	B := Zeros(A.Cols(), A.Rows())
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			B.Set(j, i, A.Get(i, j))
		}
	}
	return B
}

func (A *DenseMatrix) TransposeInPlace() (err error) {
	if A.rows != A.cols {
		err = errors.New("Can only transpose a square matrix in place")
		return
	}
	for i := 0; i < A.rows; i++ {
		for j := 0; j < i; j++ {
			tmp := A.Get(i, j)
			A.Set(i, j, A.Get(j, i))
			A.Set(j, i, tmp)
		}
	}
	return
}

func solveLower(A *DenseMatrix, b Matrix) *DenseMatrix {
	x := make([]float64, A.Cols())
	for i := 0; i < A.Rows(); i++ {
		x[i] = b.Get(i, 0)
		for j := 0; j < i; j++ {
			x[i] -= x[j] * A.Get(i, j)
		}
		//the diagonal defined to be ones
		//x[i] /= A.Get(i, i);
	}
	return MakeDenseMatrix(x, A.Cols(), 1)
}

func solveUpper(A *DenseMatrix, b Matrix) *DenseMatrix {
	x := make([]float64, A.Cols())
	for i := A.Rows() - 1; i >= 0; i-- {
		x[i] = b.Get(i, 0)
		for j := i + 1; j < A.Cols(); j++ {
			x[i] -= x[j] * A.Get(i, j)
		}
		x[i] /= A.Get(i, i)
	}
	return MakeDenseMatrix(x, A.Cols(), 1)
}

func (A *DenseMatrix) Solve(b MatrixRO) (*DenseMatrix, error) {
	Acopy := A.Copy()
	P := Acopy.LUInPlace()
	Pinv := P.Inverse()
	pb, err := Pinv.Times(b)

	if !(err == nil) {
		return nil, err
	}

	y := solveLower(Acopy, pb)
	x := solveUpper(Acopy, y)
	return x, nil
}

func (A *DenseMatrix) SolveDense(b *DenseMatrix) (*DenseMatrix, error) {
	return A.Solve(b)
}

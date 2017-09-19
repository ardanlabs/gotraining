// Copyright ©2014 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file must be kept in sync with index_bound_checks.go.

//+build !bounds

package mat64

import "github.com/gonum/matrix"

// At returns the element at row i, column j.
func (m *Dense) At(i, j int) float64 {
	if uint(i) >= uint(m.mat.Rows) {
		panic(matrix.ErrRowAccess)
	}
	if uint(j) >= uint(m.mat.Cols) {
		panic(matrix.ErrColAccess)
	}
	return m.at(i, j)
}

func (m *Dense) at(i, j int) float64 {
	return m.mat.Data[i*m.mat.Stride+j]
}

// Set sets the element at row i, column j to the value v.
func (m *Dense) Set(i, j int, v float64) {
	if uint(i) >= uint(m.mat.Rows) {
		panic(matrix.ErrRowAccess)
	}
	if uint(j) >= uint(m.mat.Cols) {
		panic(matrix.ErrColAccess)
	}
	m.set(i, j, v)
}

func (m *Dense) set(i, j int, v float64) {
	m.mat.Data[i*m.mat.Stride+j] = v
}

// At returns the element at row i.
// It panics if i is out of bounds or if j is not zero.
func (v *Vector) At(i, j int) float64 {
	if uint(i) >= uint(v.n) {
		panic(matrix.ErrRowAccess)
	}
	if j != 0 {
		panic(matrix.ErrColAccess)
	}
	return v.at(i)
}

func (v *Vector) at(i int) float64 {
	return v.mat.Data[i*v.mat.Inc]
}

// SetVec sets the element at row i to the value val.
// It panics if i is out of bounds.
func (v *Vector) SetVec(i int, val float64) {
	if uint(i) >= uint(v.n) {
		panic(matrix.ErrVectorAccess)
	}
	v.setVec(i, val)
}

func (v *Vector) setVec(i int, val float64) {
	v.mat.Data[i*v.mat.Inc] = val
}

// At returns the element at row i and column j.
func (s *SymDense) At(i, j int) float64 {
	if uint(i) >= uint(s.mat.N) {
		panic(matrix.ErrRowAccess)
	}
	if uint(j) >= uint(s.mat.N) {
		panic(matrix.ErrColAccess)
	}
	return s.at(i, j)
}

func (s *SymDense) at(i, j int) float64 {
	if i > j {
		i, j = j, i
	}
	return s.mat.Data[i*s.mat.Stride+j]
}

// SetSym sets the elements at (i,j) and (j,i) to the value v.
func (s *SymDense) SetSym(i, j int, v float64) {
	if uint(i) >= uint(s.mat.N) {
		panic(matrix.ErrRowAccess)
	}
	if uint(j) >= uint(s.mat.N) {
		panic(matrix.ErrColAccess)
	}
	s.set(i, j, v)
}

func (s *SymDense) set(i, j int, v float64) {
	if i > j {
		i, j = j, i
	}
	s.mat.Data[i*s.mat.Stride+j] = v
}

// At returns the element at row i, column j.
func (t *TriDense) At(i, j int) float64 {
	if uint(i) >= uint(t.mat.N) {
		panic(matrix.ErrRowAccess)
	}
	if uint(j) >= uint(t.mat.N) {
		panic(matrix.ErrColAccess)
	}
	return t.at(i, j)
}

func (t *TriDense) at(i, j int) float64 {
	isUpper := t.triKind()
	if (isUpper && i > j) || (!isUpper && i < j) {
		return 0
	}
	return t.mat.Data[i*t.mat.Stride+j]
}

// SetTri sets the element at row i, column j to the value v.
// It panics if the location is outside the appropriate half of the matrix.
func (t *TriDense) SetTri(i, j int, v float64) {
	if uint(i) >= uint(t.mat.N) {
		panic(matrix.ErrRowAccess)
	}
	if uint(j) >= uint(t.mat.N) {
		panic(matrix.ErrColAccess)
	}
	isUpper := t.isUpper()
	if (isUpper && i > j) || (!isUpper && i < j) {
		panic(matrix.ErrTriangleSet)
	}
	t.set(i, j, v)
}

func (t *TriDense) set(i, j int, v float64) {
	t.mat.Data[i*t.mat.Stride+j] = v
}

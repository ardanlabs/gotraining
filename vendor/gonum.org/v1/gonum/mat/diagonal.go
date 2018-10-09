// Copyright ©2018 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat

import (
	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/blas/blas64"
)

var (
	diagDense *DiagDense
	_         Matrix          = diagDense
	_         Diagonal        = diagDense
	_         MutableDiagonal = diagDense
	_         Triangular      = diagDense
	_         Symmetric       = diagDense
	_         Banded          = diagDense
	_         RawBander       = diagDense
	_         RawSymBander    = diagDense
)

// Diagonal represents a diagonal matrix, that is a square matrix that only
// has non-zero terms on the diagonal.
type Diagonal interface {
	Matrix
	// Diag returns the number of rows/columns in the matrix
	Diag() int
}

// MutableDiagonal is a Diagonal matrix whose elements can be set.
type MutableDiagonal interface {
	Diagonal
	SetDiag(i int, v float64)
}

// DiagDense represents a diagonal matrix in dense storage format.
type DiagDense struct {
	data []float64
}

// NewDiagonal creates a new Diagonal matrix with n rows and n columns.
// The length of data must be n or data must be nil, otherwise NewDiagonal
// will panic.
func NewDiagonal(n int, data []float64) *DiagDense {
	if n < 0 {
		panic("mat: negative dimension")
	}
	if data == nil {
		data = make([]float64, n)
	}
	if len(data) != n {
		panic(ErrShape)
	}
	return &DiagDense{
		data: data,
	}
}

// Diag returns the dimension of the receiver.
func (d *DiagDense) Diag() int {
	return len(d.data)
}

// Dims returns the dimensions of the matrix.
func (d *DiagDense) Dims() (r, c int) {
	return len(d.data), len(d.data)
}

// T returns the transpose of the matrix.
func (d *DiagDense) T() Matrix {
	return d
}

// TTri returns the transpose of the matrix. Note that Diagonal matrices are
// Upper by default
func (d *DiagDense) TTri() Triangular {
	return TransposeTri{d}
}

func (d *DiagDense) TBand() Banded {
	return TransposeBand{d}
}

func (d *DiagDense) Bandwidth() (kl, ku int) {
	return 0, 0
}

// Symmetric implements the Symmetric interface.
func (d *DiagDense) Symmetric() int {
	return len(d.data)
}

// Triangle implements the Triangular interface.
func (d *DiagDense) Triangle() (int, TriKind) {
	return len(d.data), Upper
}

func (d *DiagDense) RawBand() blas64.Band {
	return blas64.Band{
		Rows:   len(d.data),
		Cols:   len(d.data),
		KL:     0,
		KU:     0,
		Stride: 1,
		Data:   d.data,
	}
}

func (d *DiagDense) RawSymBand() blas64.SymmetricBand {
	return blas64.SymmetricBand{
		N:      len(d.data),
		K:      0,
		Stride: 1,
		Uplo:   blas.Upper,
		Data:   d.data,
	}
}

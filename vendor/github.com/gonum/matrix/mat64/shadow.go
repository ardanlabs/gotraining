// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat64

import (
	"github.com/gonum/blas"
	"github.com/gonum/blas/blas64"
)

const (
	// regionOverlap is the panic string used for the general case
	// of a matrix region overlap between a source and destination.
	regionOverlap = "mat64: bad region: overlap"

	// regionIdentity is the panic string used for the specific
	// case of complete agreement between a source and a destination.
	regionIdentity = "mat64: bad region: identical"

	// mismatchedStrides is the panic string used for overlapping
	// data slices with differing strides.
	mismatchedStrides = "mat64: bad region: different strides"
)

// checkOverlap returns false if the receiver does not overlap data elements
// referenced by the parameter and panics otherwise.
//
// checkOverlap methods return a boolean to allow the check call to be added to a
// boolean expression, making use of short-circuit operators.

func (m *Dense) checkOverlap(a blas64.General) bool {
	mat := m.RawMatrix()
	if cap(mat.Data) == 0 || cap(a.Data) == 0 {
		return false
	}

	off := offset(mat.Data[:1], a.Data[:1])

	if off == 0 {
		// At least one element overlaps.
		if mat.Cols == a.Cols && mat.Rows == a.Rows && mat.Stride == a.Stride {
			panic(regionIdentity)
		}
		panic(regionOverlap)
	}

	if off > 0 && len(mat.Data) <= off {
		// We know m is completely before a.
		return false
	}
	if off < 0 && len(a.Data) <= -off {
		// We know m is completely after a.
		return false
	}

	if mat.Stride != a.Stride {
		// Too hard, so assume the worst.
		panic(mismatchedStrides)
	}

	if off < 0 {
		off = -off
		mat.Cols, a.Cols = a.Cols, mat.Cols
	}
	if rectanglesOverlap(off, mat.Cols, a.Cols, mat.Stride) {
		panic(regionOverlap)
	}
	return false
}

func (s *SymDense) checkOverlap(a blas64.Symmetric) bool {
	mat := s.RawSymmetric()
	if cap(mat.Data) == 0 || cap(a.Data) == 0 {
		return false
	}

	off := offset(mat.Data[:1], a.Data[:1])

	if off == 0 {
		// At least one element overlaps.
		if mat.N == a.N && mat.Stride == a.Stride {
			panic(regionIdentity)
		}
		panic(regionOverlap)
	}

	if off > 0 && len(mat.Data) <= off {
		// We know s is completely before a.
		return false
	}
	if off < 0 && len(a.Data) <= -off {
		// We know s is completely after a.
		return false
	}

	if mat.Stride != a.Stride {
		// Too hard, so assume the worst.
		panic(mismatchedStrides)
	}

	if off < 0 {
		off = -off
		mat.N, a.N = a.N, mat.N
		// If we created the matrix it will always
		// be in the upper triangle, but don't trust
		// that this is the case.
		mat.Uplo, a.Uplo = a.Uplo, mat.Uplo
	}
	if trianglesOverlap(off, mat.N, a.N, mat.Stride, mat.Uplo == blas.Upper, a.Uplo == blas.Upper) {
		panic(regionOverlap)
	}
	return false
}

func (t *TriDense) checkOverlap(a blas64.Triangular) bool {
	mat := t.RawTriangular()
	if cap(mat.Data) == 0 || cap(a.Data) == 0 {
		return false
	}

	off := offset(mat.Data[:1], a.Data[:1])

	if off == 0 {
		// At least one element overlaps.
		if mat.N == a.N && mat.Stride == a.Stride {
			panic(regionIdentity)
		}
		panic(regionOverlap)
	}

	if off > 0 && len(mat.Data) <= off {
		// We know t is completely before a.
		return false
	}
	if off < 0 && len(a.Data) <= -off {
		// We know t is completely after a.
		return false
	}

	if mat.Stride != a.Stride {
		// Too hard, so assume the worst.
		panic(mismatchedStrides)
	}

	if off < 0 {
		off = -off
		mat.N, a.N = a.N, mat.N
		mat.Uplo, a.Uplo = a.Uplo, mat.Uplo
	}
	if trianglesOverlap(off, mat.N, a.N, mat.Stride, mat.Uplo == blas.Upper, a.Uplo == blas.Upper) {
		panic(regionOverlap)
	}
	return false
}

func (v *Vector) checkOverlap(a blas64.Vector) bool {
	mat := v.mat
	if cap(mat.Data) == 0 || cap(a.Data) == 0 {
		return false
	}

	off := offset(mat.Data[:1], a.Data[:1])

	if off == 0 {
		// At least one element overlaps.
		if mat.Inc == a.Inc && len(mat.Data) == len(a.Data) {
			panic(regionIdentity)
		}
		panic(regionOverlap)
	}

	if off > 0 && len(mat.Data) <= off {
		// We know v is completely before a.
		return false
	}
	if off < 0 && len(a.Data) <= -off {
		// We know v is completely after a.
		return false
	}

	if mat.Inc != a.Inc {
		// Too hard, so assume the worst.
		panic(mismatchedStrides)
	}

	if mat.Inc == 1 || off&mat.Inc == 0 {
		panic(regionOverlap)
	}
	return false
}

// rectanglesOverlap returns whether the strided rectangles a and b overlap
// when b is offset by off elements after a but has at least one element before
// the end of a. off must be positive. a and b have aCols and bCols respectively.
//
// rectanglesOverlap works by shifting both matrices left such that the left
// column of a is at 0. The column indexes are flattened by obtaining the shifted
// relative left and right column positions modulo the common stride. This allows
// direct comparison of the column offsets when the matrix backing data slices
// are known to overlap.
func rectanglesOverlap(off, aCols, bCols, stride int) bool {
	if stride == 1 {
		// Unit stride means overlapping data
		// slices must overlap as matrices.
		return true
	}

	// Flatten the shifted matrix column positions
	// so a starts at 0, modulo the common stride.
	const aFrom = 0
	aTo := aCols
	// The mod stride operations here make the from
	// and to indexes comparable between a and b when
	// the data slices of a and b overlap.
	bFrom := off % stride
	bTo := (bFrom + bCols) % stride

	if bTo == 0 || bFrom < bTo {
		// b matrix is not wrapped: compare for
		// simple overlap.
		return bFrom < aTo
	}

	// b strictly wraps and so must overlap with a.
	return true
}

// trianglesOverlap returns whether the strided triangles a and b overlap
// when b is offset by off elements after a but has at least one element before
// the end of a. off must be positive. a and b are aSize×aSize and bSize×bSize
// respectively.
func trianglesOverlap(off, aSize, bSize, stride int, aUpper, bUpper bool) bool {
	if !rectanglesOverlap(off, aSize, bSize, stride) {
		// Fast return if bounding rectangles do not overlap.
		return false
	}

	// Find location of b relative to a.
	rowOffset := off / stride
	colOffset := off % stride
	if (off+bSize)%stride < colOffset {
		// We have wrapped, so readjust offsets.
		rowOffset++
		colOffset -= stride
	}

	if aUpper {
		// Check whether the upper left of b
		// is in the triangle of a
		if rowOffset >= 0 && rowOffset <= colOffset {
			return true
		}
		// Check whether the upper right of b
		// is in the triangle of a.
		return bUpper && rowOffset < colOffset+bSize
	}

	// Check whether the upper left of b
	// is in the triangle of a
	if colOffset >= 0 && rowOffset >= colOffset {
		return true
	}
	if bUpper {
		// Check whether the upper right corner of b
		// is in a or the upper row of b spans a row
		// of a.
		return rowOffset > colOffset+bSize || colOffset < 0
	}
	if colOffset < 0 {
		// Check whether the lower left of a
		// is in the triangle of b or below
		// the diagonal of a. This requires a
		// swap of reference origin.
		return -rowOffset+aSize > -colOffset
	}
	// Check whether the lower left of b
	// is in the triangle of a or below
	// the diagonal of a.
	return rowOffset+bSize > colOffset
}

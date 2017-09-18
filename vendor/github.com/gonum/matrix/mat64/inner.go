// Copyright ©2014 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat64

import (
	"github.com/gonum/blas"
	"github.com/gonum/internal/asm/f64"
	"github.com/gonum/matrix"
)

// Inner computes the generalized inner product
//   x^T A y
// between vectors x and y with matrix A. This is only a true inner product if
// A is symmetric positive definite, though the operation works for any matrix A.
//
// Inner panics if x.Len != m or y.Len != n when A is an m x n matrix.
func Inner(x *Vector, A Matrix, y *Vector) float64 {
	m, n := A.Dims()
	if x.Len() != m {
		panic(matrix.ErrShape)
	}
	if y.Len() != n {
		panic(matrix.ErrShape)
	}
	if m == 0 || n == 0 {
		return 0
	}

	var sum float64

	switch b := A.(type) {
	case RawSymmetricer:
		bmat := b.RawSymmetric()
		if bmat.Uplo != blas.Upper {
			// Panic as a string not a mat64.Error.
			panic(badSymTriangle)
		}
		for i := 0; i < x.Len(); i++ {
			xi := x.at(i)
			if xi != 0 {
				if y.mat.Inc == 1 {
					sum += xi * f64.DotUnitary(
						bmat.Data[i*bmat.Stride+i:i*bmat.Stride+n],
						y.mat.Data[i:],
					)
				} else {
					sum += xi * f64.DotInc(
						bmat.Data[i*bmat.Stride+i:i*bmat.Stride+n],
						y.mat.Data[i*y.mat.Inc:], uintptr(n-i),
						1, uintptr(y.mat.Inc),
						0, 0,
					)
				}
			}
			yi := y.at(i)
			if i != n-1 && yi != 0 {
				if x.mat.Inc == 1 {
					sum += yi * f64.DotUnitary(
						bmat.Data[i*bmat.Stride+i+1:i*bmat.Stride+n],
						x.mat.Data[i+1:],
					)
				} else {
					sum += yi * f64.DotInc(
						bmat.Data[i*bmat.Stride+i+1:i*bmat.Stride+n],
						x.mat.Data[(i+1)*x.mat.Inc:], uintptr(n-i-1),
						1, uintptr(x.mat.Inc),
						0, 0,
					)
				}
			}
		}
	case RawMatrixer:
		bmat := b.RawMatrix()
		for i := 0; i < x.Len(); i++ {
			xi := x.at(i)
			if xi != 0 {
				if y.mat.Inc == 1 {
					sum += xi * f64.DotUnitary(
						bmat.Data[i*bmat.Stride:i*bmat.Stride+n],
						y.mat.Data,
					)
				} else {
					sum += xi * f64.DotInc(
						bmat.Data[i*bmat.Stride:i*bmat.Stride+n],
						y.mat.Data, uintptr(n),
						1, uintptr(y.mat.Inc),
						0, 0,
					)
				}
			}
		}
	default:
		for i := 0; i < x.Len(); i++ {
			xi := x.at(i)
			for j := 0; j < y.Len(); j++ {
				sum += xi * A.At(i, j) * y.at(j)
			}
		}
	}
	return sum
}

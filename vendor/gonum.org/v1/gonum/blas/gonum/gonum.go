// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate ./single_precision.bash

package gonum

import "math"

type Implementation struct{}

// The following are panic strings used during parameter checks.
const (
	negativeN = "blas: n < 0"
	zeroIncX  = "blas: zero x index increment"
	zeroIncY  = "blas: zero y index increment"
	badLenX   = "blas: x index out of range"
	badLenY   = "blas: y index out of range"

	mLT0  = "blas: m < 0"
	nLT0  = "blas: n < 0"
	kLT0  = "blas: k < 0"
	kLLT0 = "blas: kL < 0"
	kULT0 = "blas: kU < 0"

	badUplo      = "blas: illegal triangle"
	badTranspose = "blas: illegal transpose"
	badDiag      = "blas: illegal diagonal"
	badSide      = "blas: illegal side"

	badLdA = "blas: index of a out of range"
	badLdB = "blas: index of b out of range"
	badLdC = "blas: index of c out of range"

	badX = "blas: x index out of range"
	badY = "blas: y index out of range"
)

// [SD]gemm behavior constants. These are kept here to keep them out of the
// way during single precision code genration.
const (
	blockSize   = 64 // b x b matrix
	minParBlock = 4  // minimum number of blocks needed to go parallel
	buffMul     = 4  // how big is the buffer relative to the number of workers
)

// subMul is a common type shared by [SD]gemm.
type subMul struct {
	i, j int // index of block
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func checkSMatrix(name byte, m, n int, a []float32, lda int) {
	if m < 0 {
		panic(mLT0)
	}
	if n < 0 {
		panic(nLT0)
	}
	if lda < n {
		panic("blas: illegal stride of " + string(name))
	}
	if len(a) < (m-1)*lda+n {
		panic("blas: index of " + string(name) + " out of range")
	}
}

func checkDMatrix(name byte, m, n int, a []float64, lda int) {
	if m < 0 {
		panic(mLT0)
	}
	if n < 0 {
		panic(nLT0)
	}
	if lda < n {
		panic("blas: illegal stride of " + string(name))
	}
	if len(a) < (m-1)*lda+n {
		panic("blas: index of " + string(name) + " out of range")
	}
}

func checkZMatrix(name byte, m, n int, a []complex128, lda int) {
	if m < 0 {
		panic(mLT0)
	}
	if n < 0 {
		panic(nLT0)
	}
	if lda < max(1, n) {
		panic("blas: illegal stride of " + string(name))
	}
	if len(a) < (m-1)*lda+n {
		panic("blas: insufficient " + string(name) + " matrix slice length")
	}
}

func checkZVector(name byte, n int, x []complex128, incX int) {
	if n < 0 {
		panic(nLT0)
	}
	if incX == 0 {
		panic(zeroIncX)
	}
	if (incX > 0 && (n-1)*incX >= len(x)) || (incX < 0 && (1-n)*incX >= len(x)) {
		panic("blas: insufficient " + string(name) + " vector slice length")
	}
}

// blocks returns the number of divisions of the dimension length with the given
// block size.
func blocks(dim, bsize int) int {
	return (dim + bsize - 1) / bsize
}

// dcabs1 returns |real(z)|+|imag(z)|.
func dcabs1(z complex128) float64 {
	return math.Abs(real(z)) + math.Abs(imag(z))
}

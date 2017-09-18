// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate ./single_precision.bash

package native

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

// [SD]gemm debugging constant.
const debug = false

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

// blocks returns the number of divisons of the dimension length with the given
// block size.
func blocks(dim, bsize int) int {
	return (dim + bsize - 1) / bsize
}

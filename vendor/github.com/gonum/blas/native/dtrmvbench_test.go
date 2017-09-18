// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.7

package native

import (
	"strconv"
	"testing"

	"github.com/gonum/blas"
	"github.com/gonum/blas/testblas"
)

func BenchmarkDtrmv(b *testing.B) {
	for _, n := range []int{testblas.MediumMat, testblas.LargeMat} {
		for _, incX := range []int{1, 5} {
			for _, uplo := range []blas.Uplo{blas.Upper, blas.Lower} {
				for _, trans := range []blas.Transpose{blas.NoTrans, blas.Trans} {
					for _, unit := range []blas.Diag{blas.NonUnit, blas.Unit} {
						var str string
						if n == testblas.MediumMat {
							str += "Med"
						} else if n == testblas.LargeMat {
							str += "Large"
						}
						str += "_Inc" + strconv.Itoa(incX)
						if uplo == blas.Upper {
							str += "_UP"
						} else {
							str += "_LO"
						}
						if trans == blas.NoTrans {
							str += "_NT"
						} else {
							str += "_TR"
						}
						if unit == blas.NonUnit {
							str += "_NU"
						} else {
							str += "_UN"
						}
						lda := n
						b.Run(str, func(b *testing.B) {
							testblas.DtrmvBenchmark(b, Implementation{}, n, lda, incX, uplo, trans, unit)
						})
					}
				}
			}
		}
	}
}

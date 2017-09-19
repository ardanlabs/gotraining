// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.7

package gonum

import (
	"fmt"
	"math/rand"
	"testing"

	"gonum.org/v1/gonum/blas"
)

var benchSinkZ []complex128

func BenchmarkZher(b *testing.B) {
	for _, uplo := range []blas.Uplo{blas.Upper, blas.Lower} {
		for _, n := range []int{10, 100, 1000, 10000} {
			for _, inc := range []int{1, 10, 1000} {
				benchmarkZher(b, uplo, n, inc)
			}
		}
	}
}

func benchmarkZher(b *testing.B, uplo blas.Uplo, n, inc int) {
	b.Run(fmt.Sprintf("Uplo%d-N%d-Inc%d", uplo, n, inc), func(b *testing.B) {
		rnd := rand.New(rand.NewSource(1))
		alpha := rnd.NormFloat64()
		x := make([]complex128, (n-1)*inc+1)
		for i := range x {
			x[i] = complex(rnd.NormFloat64(), rnd.NormFloat64())
		}
		a := make([]complex128, len(benchSinkZ))
		for i := range a {
			a[i] = complex(rnd.NormFloat64(), rnd.NormFloat64())
		}
		benchSinkZ = make([]complex128, n*n)
		copy(benchSinkZ, a)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			impl.Zher(uplo, n, alpha, x, inc, benchSinkZ, n)
			copy(benchSinkZ, a)
		}
	})
}

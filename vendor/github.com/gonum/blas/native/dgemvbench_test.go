// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package native

import (
	"testing"

	"github.com/gonum/blas/testblas"
)

func BenchmarkDgemvSmSmNoTransInc1(b *testing.B) {
	testblas.DgemvBenchmark(b, impl, NT, Sm, Sm, 1, 1)
}

func BenchmarkDgemvSmSmNoTransIncN(b *testing.B) {
	testblas.DgemvBenchmark(b, impl, NT, Sm, Sm, 2, 3)
}

func BenchmarkDgemvSmSmTransInc1(b *testing.B) {
	testblas.DgemvBenchmark(b, impl, T, Sm, Sm, 1, 1)
}

func BenchmarkDgemvSmSmTransIncN(b *testing.B) {
	testblas.DgemvBenchmark(b, impl, T, Sm, Sm, 2, 3)
}

func BenchmarkDgemvMedMedNoTransInc1(b *testing.B) {
	testblas.DgemvBenchmark(b, impl, NT, Med, Med, 1, 1)
}

func BenchmarkDgemvMedMedNoTransIncN(b *testing.B) {
	testblas.DgemvBenchmark(b, impl, NT, Med, Med, 2, 3)
}

func BenchmarkDgemvMedMedTransInc1(b *testing.B) {
	testblas.DgemvBenchmark(b, impl, T, Med, Med, 1, 1)
}

func BenchmarkDgemvMedMedTransIncN(b *testing.B) {
	testblas.DgemvBenchmark(b, impl, T, Med, Med, 2, 3)
}

func BenchmarkDgemvLgLgNoTransInc1(b *testing.B) {
	testblas.DgemvBenchmark(b, impl, NT, Lg, Lg, 1, 1)
}

func BenchmarkDgemvLgLgNoTransIncN(b *testing.B) {
	testblas.DgemvBenchmark(b, impl, NT, Lg, Lg, 2, 3)
}

func BenchmarkDgemvLgLgTransInc1(b *testing.B) {
	testblas.DgemvBenchmark(b, impl, T, Lg, Lg, 1, 1)
}

func BenchmarkDgemvLgLgTransIncN(b *testing.B) {
	testblas.DgemvBenchmark(b, impl, T, Lg, Lg, 2, 3)
}

func BenchmarkDgemvLgSmNoTransInc1(b *testing.B) {
	testblas.DgemvBenchmark(b, impl, NT, Lg, Sm, 1, 1)
}

func BenchmarkDgemvLgSmNoTransIncN(b *testing.B) {
	testblas.DgemvBenchmark(b, impl, NT, Lg, Sm, 2, 3)
}

func BenchmarkDgemvLgSmTransInc1(b *testing.B) {
	testblas.DgemvBenchmark(b, impl, T, Lg, Sm, 1, 1)
}

func BenchmarkDgemvLgSmTransIncN(b *testing.B) {
	testblas.DgemvBenchmark(b, impl, T, Lg, Sm, 2, 3)
}

func BenchmarkDgemvSmLgNoTransInc1(b *testing.B) {
	testblas.DgemvBenchmark(b, impl, NT, Sm, Lg, 1, 1)
}

func BenchmarkDgemvSmLgNoTransIncN(b *testing.B) {
	testblas.DgemvBenchmark(b, impl, NT, Sm, Lg, 2, 3)
}

func BenchmarkDgemvSmLgTransInc1(b *testing.B) {
	testblas.DgemvBenchmark(b, impl, T, Sm, Lg, 1, 1)
}

func BenchmarkDgemvSmLgTransIncN(b *testing.B) {
	testblas.DgemvBenchmark(b, impl, T, Sm, Lg, 2, 3)
}

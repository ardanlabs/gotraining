// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package native

import (
	"testing"

	"github.com/gonum/blas/testblas"
)

func BenchmarkDgerSmSmInc1(b *testing.B) {
	testblas.DgerBenchmark(b, impl, Sm, Sm, 1, 1)
}

func BenchmarkDgerSmSmIncN(b *testing.B) {
	testblas.DgerBenchmark(b, impl, Sm, Sm, 2, 3)
}

func BenchmarkDgerMedMedInc1(b *testing.B) {
	testblas.DgerBenchmark(b, impl, Med, Med, 1, 1)
}

func BenchmarkDgerMedMedIncN(b *testing.B) {
	testblas.DgerBenchmark(b, impl, Med, Med, 2, 3)
}

func BenchmarkDgerLgLgInc1(b *testing.B) {
	testblas.DgerBenchmark(b, impl, Lg, Lg, 1, 1)
}

func BenchmarkDgerLgLgIncN(b *testing.B) {
	testblas.DgerBenchmark(b, impl, Lg, Lg, 2, 3)
}

func BenchmarkDgerLgSmInc1(b *testing.B) {
	testblas.DgerBenchmark(b, impl, Lg, Sm, 1, 1)
}

func BenchmarkDgerLgSmIncN(b *testing.B) {
	testblas.DgerBenchmark(b, impl, Lg, Sm, 2, 3)
}

func BenchmarkDgerSmLgInc1(b *testing.B) {
	testblas.DgerBenchmark(b, impl, Sm, Lg, 1, 1)
}

func BenchmarkDgerSmLgIncN(b *testing.B) {
	testblas.DgerBenchmark(b, impl, Sm, Lg, 2, 3)
}

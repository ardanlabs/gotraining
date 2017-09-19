// Copyright ©2014 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package native

import (
	"testing"

	"github.com/gonum/blas/testblas"
)

func BenchmarkDgemmSmSmSm(b *testing.B) {
	testblas.DgemmBenchmark(b, impl, Sm, Sm, Sm, NT, NT)
}

func BenchmarkDgemmMedMedMed(b *testing.B) {
	testblas.DgemmBenchmark(b, impl, Med, Med, Med, NT, NT)
}

func BenchmarkDgemmMedLgMed(b *testing.B) {
	testblas.DgemmBenchmark(b, impl, Med, Lg, Med, NT, NT)
}

func BenchmarkDgemmLgLgLg(b *testing.B) {
	testblas.DgemmBenchmark(b, impl, Lg, Lg, Lg, NT, NT)
}

func BenchmarkDgemmLgSmLg(b *testing.B) {
	testblas.DgemmBenchmark(b, impl, Lg, Sm, Lg, NT, NT)
}

func BenchmarkDgemmLgLgSm(b *testing.B) {
	testblas.DgemmBenchmark(b, impl, Lg, Lg, Sm, NT, NT)
}

func BenchmarkDgemmHgHgSm(b *testing.B) {
	testblas.DgemmBenchmark(b, impl, Hg, Hg, Sm, NT, NT)
}

func BenchmarkDgemmMedMedMedTNT(b *testing.B) {
	testblas.DgemmBenchmark(b, impl, Med, Med, Med, T, NT)
}

func BenchmarkDgemmMedMedMedNTT(b *testing.B) {
	testblas.DgemmBenchmark(b, impl, Med, Med, Med, NT, T)
}

func BenchmarkDgemmMedMedMedTT(b *testing.B) {
	testblas.DgemmBenchmark(b, impl, Med, Med, Med, T, T)
}

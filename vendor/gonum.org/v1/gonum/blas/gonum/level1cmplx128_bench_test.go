// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gonum

import (
	"math/rand"
	"testing"
)

func benchmarkZdscal(b *testing.B, n, inc int) {
	rnd := rand.New(rand.NewSource(1))
	alpha := rnd.NormFloat64()
	x := make([]complex128, (n-1)*inc+1)
	for i := range x {
		x[i] = complex(rnd.NormFloat64(), rnd.NormFloat64())
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		impl.Zdscal(n, alpha, x, inc)
	}
}

func BenchmarkZdscalN10Inc1(b *testing.B)     { benchmarkZdscal(b, 10, 1) }
func BenchmarkZdscalN100Inc1(b *testing.B)    { benchmarkZdscal(b, 100, 1) }
func BenchmarkZdscalN1000Inc1(b *testing.B)   { benchmarkZdscal(b, 1000, 1) }
func BenchmarkZdscalN10000Inc1(b *testing.B)  { benchmarkZdscal(b, 10000, 1) }
func BenchmarkZdscalN100000Inc1(b *testing.B) { benchmarkZdscal(b, 100000, 1) }

func BenchmarkZdscalN10Inc10(b *testing.B)     { benchmarkZdscal(b, 10, 10) }
func BenchmarkZdscalN100Inc10(b *testing.B)    { benchmarkZdscal(b, 100, 10) }
func BenchmarkZdscalN1000Inc10(b *testing.B)   { benchmarkZdscal(b, 1000, 10) }
func BenchmarkZdscalN10000Inc10(b *testing.B)  { benchmarkZdscal(b, 10000, 10) }
func BenchmarkZdscalN100000Inc10(b *testing.B) { benchmarkZdscal(b, 100000, 10) }

func BenchmarkZdscalN10Inc1000(b *testing.B)     { benchmarkZdscal(b, 10, 1000) }
func BenchmarkZdscalN100Inc1000(b *testing.B)    { benchmarkZdscal(b, 100, 1000) }
func BenchmarkZdscalN1000Inc1000(b *testing.B)   { benchmarkZdscal(b, 1000, 1000) }
func BenchmarkZdscalN10000Inc1000(b *testing.B)  { benchmarkZdscal(b, 10000, 1000) }
func BenchmarkZdscalN100000Inc1000(b *testing.B) { benchmarkZdscal(b, 100000, 1000) }

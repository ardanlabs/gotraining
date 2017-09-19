// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.5,!go1.7

package c128

import "testing"

func benchdotu(b *testing.B, n int64, fn func(x, y []complex128) complex128) {
	x, y := x[:n], y[:n]
	b.SetBytes(256 * n)
	for i := 0; i < b.N; i++ {
		benchSink = fn(x, y)
	}
}

func BenchmarkDotcUnitary1(t *testing.B)     { benchdotu(t, 1, DotcUnitary) }
func BenchmarkDotcUnitary2(t *testing.B)     { benchdotu(t, 2, DotcUnitary) }
func BenchmarkDotcUnitary3(t *testing.B)     { benchdotu(t, 3, DotcUnitary) }
func BenchmarkDotcUnitary4(t *testing.B)     { benchdotu(t, 4, DotcUnitary) }
func BenchmarkDotcUnitary5(t *testing.B)     { benchdotu(t, 5, DotcUnitary) }
func BenchmarkDotcUnitary10(t *testing.B)    { benchdotu(t, 10, DotcUnitary) }
func BenchmarkDotcUnitary100(t *testing.B)   { benchdotu(t, 100, DotcUnitary) }
func BenchmarkDotcUnitary1000(t *testing.B)  { benchdotu(t, 1000, DotcUnitary) }
func BenchmarkDotcUnitary5000(t *testing.B)  { benchdotu(t, 5000, DotcUnitary) }
func BenchmarkDotcUnitary10000(t *testing.B) { benchdotu(t, 10000, DotcUnitary) }
func BenchmarkDotcUnitary50000(t *testing.B) { benchdotu(t, 50000, DotcUnitary) }

func BenchmarkDotuUnitary1(t *testing.B)     { benchdotu(t, 1, DotuUnitary) }
func BenchmarkDotuUnitary2(t *testing.B)     { benchdotu(t, 2, DotuUnitary) }
func BenchmarkDotuUnitary3(t *testing.B)     { benchdotu(t, 3, DotuUnitary) }
func BenchmarkDotuUnitary4(t *testing.B)     { benchdotu(t, 4, DotuUnitary) }
func BenchmarkDotuUnitary5(t *testing.B)     { benchdotu(t, 5, DotuUnitary) }
func BenchmarkDotuUnitary10(t *testing.B)    { benchdotu(t, 10, DotuUnitary) }
func BenchmarkDotuUnitary100(t *testing.B)   { benchdotu(t, 100, DotuUnitary) }
func BenchmarkDotuUnitary1000(t *testing.B)  { benchdotu(t, 1000, DotuUnitary) }
func BenchmarkDotuUnitary5000(t *testing.B)  { benchdotu(t, 5000, DotuUnitary) }
func BenchmarkDotuUnitary10000(t *testing.B) { benchdotu(t, 10000, DotuUnitary) }

func benchdoti(b *testing.B, ln, inc int, fn func(x, y []complex128, n, incX, incY, ix, iy uintptr) complex128) {
	b.SetBytes(int64(256 * ln))
	var idx int
	if inc < 0 {
		idx = (-ln + 1) * inc
	}
	for i := 0; i < b.N; i++ {
		benchSink = fn(x, y, uintptr(ln), uintptr(inc), uintptr(inc), uintptr(idx), uintptr(idx))
	}
}

func BenchmarkDotcInc_1_inc1(t *testing.B) { benchdoti(t, 1, 1, DotcInc) }

func BenchmarkDotcInc_2_inc1(t *testing.B)  { benchdoti(t, 2, 1, DotcInc) }
func BenchmarkDotcInc_2_inc2(t *testing.B)  { benchdoti(t, 2, 2, DotcInc) }
func BenchmarkDotcInc_2_inc4(t *testing.B)  { benchdoti(t, 2, 4, DotcInc) }
func BenchmarkDotcInc_2_inc10(t *testing.B) { benchdoti(t, 2, 10, DotcInc) }

func BenchmarkDotcInc_3_inc1(t *testing.B)  { benchdoti(t, 3, 1, DotcInc) }
func BenchmarkDotcInc_3_inc2(t *testing.B)  { benchdoti(t, 3, 2, DotcInc) }
func BenchmarkDotcInc_3_inc4(t *testing.B)  { benchdoti(t, 3, 4, DotcInc) }
func BenchmarkDotcInc_3_inc10(t *testing.B) { benchdoti(t, 3, 10, DotcInc) }

func BenchmarkDotcInc_4_inc1(t *testing.B)  { benchdoti(t, 4, 1, DotcInc) }
func BenchmarkDotcInc_4_inc2(t *testing.B)  { benchdoti(t, 4, 2, DotcInc) }
func BenchmarkDotcInc_4_inc4(t *testing.B)  { benchdoti(t, 4, 4, DotcInc) }
func BenchmarkDotcInc_4_inc10(t *testing.B) { benchdoti(t, 4, 10, DotcInc) }

func BenchmarkDotcInc_10_inc1(t *testing.B)  { benchdoti(t, 10, 1, DotcInc) }
func BenchmarkDotcInc_10_inc2(t *testing.B)  { benchdoti(t, 10, 2, DotcInc) }
func BenchmarkDotcInc_10_inc4(t *testing.B)  { benchdoti(t, 10, 4, DotcInc) }
func BenchmarkDotcInc_10_inc10(t *testing.B) { benchdoti(t, 10, 10, DotcInc) }

func BenchmarkDotcInc_1000_inc1(t *testing.B)  { benchdoti(t, 1000, 1, DotcInc) }
func BenchmarkDotcInc_1000_inc2(t *testing.B)  { benchdoti(t, 1000, 2, DotcInc) }
func BenchmarkDotcInc_1000_inc4(t *testing.B)  { benchdoti(t, 1000, 4, DotcInc) }
func BenchmarkDotcInc_1000_inc10(t *testing.B) { benchdoti(t, 1000, 10, DotcInc) }

func BenchmarkDotcInc_100000_inc1(t *testing.B)  { benchdoti(t, 100000, 1, DotcInc) }
func BenchmarkDotcInc_100000_inc2(t *testing.B)  { benchdoti(t, 100000, 2, DotcInc) }
func BenchmarkDotcInc_100000_inc4(t *testing.B)  { benchdoti(t, 100000, 4, DotcInc) }
func BenchmarkDotcInc_100000_inc10(t *testing.B) { benchdoti(t, 100000, 10, DotcInc) }

func BenchmarkDotcInc_100000_incM1(t *testing.B) { benchdoti(t, 100000, -1, DotcInc) }
func BenchmarkDotcInc_100000_incM2(t *testing.B) { benchdoti(t, 100000, -2, DotcInc) }
func BenchmarkDotcInc_100000_incM4(t *testing.B) { benchdoti(t, 100000, -4, DotcInc) }

func BenchmarkDotuInc_1_inc1(t *testing.B) { benchdoti(t, 1, 1, DotuInc) }

func BenchmarkDotuInc_2_inc1(t *testing.B)  { benchdoti(t, 2, 1, DotuInc) }
func BenchmarkDotuInc_2_inc2(t *testing.B)  { benchdoti(t, 2, 2, DotuInc) }
func BenchmarkDotuInc_2_inc4(t *testing.B)  { benchdoti(t, 2, 4, DotuInc) }
func BenchmarkDotuInc_2_inc10(t *testing.B) { benchdoti(t, 2, 10, DotuInc) }

func BenchmarkDotuInc_3_inc1(t *testing.B)  { benchdoti(t, 3, 1, DotuInc) }
func BenchmarkDotuInc_3_inc2(t *testing.B)  { benchdoti(t, 3, 2, DotuInc) }
func BenchmarkDotuInc_3_inc4(t *testing.B)  { benchdoti(t, 3, 4, DotuInc) }
func BenchmarkDotuInc_3_inc10(t *testing.B) { benchdoti(t, 3, 10, DotuInc) }

func BenchmarkDotuInc_4_inc1(t *testing.B)  { benchdoti(t, 4, 1, DotuInc) }
func BenchmarkDotuInc_4_inc2(t *testing.B)  { benchdoti(t, 4, 2, DotuInc) }
func BenchmarkDotuInc_4_inc4(t *testing.B)  { benchdoti(t, 4, 4, DotuInc) }
func BenchmarkDotuInc_4_inc10(t *testing.B) { benchdoti(t, 4, 10, DotuInc) }

func BenchmarkDotuInc_10_inc1(t *testing.B)  { benchdoti(t, 10, 1, DotuInc) }
func BenchmarkDotuInc_10_inc2(t *testing.B)  { benchdoti(t, 10, 2, DotuInc) }
func BenchmarkDotuInc_10_inc4(t *testing.B)  { benchdoti(t, 10, 4, DotuInc) }
func BenchmarkDotuInc_10_inc10(t *testing.B) { benchdoti(t, 10, 10, DotuInc) }

func BenchmarkDotuInc_1000_inc1(t *testing.B)  { benchdoti(t, 1000, 1, DotuInc) }
func BenchmarkDotuInc_1000_inc2(t *testing.B)  { benchdoti(t, 1000, 2, DotuInc) }
func BenchmarkDotuInc_1000_inc4(t *testing.B)  { benchdoti(t, 1000, 4, DotuInc) }
func BenchmarkDotuInc_1000_inc10(t *testing.B) { benchdoti(t, 1000, 10, DotuInc) }

func BenchmarkDotuInc_100000_inc1(t *testing.B)  { benchdoti(t, 100000, 1, DotuInc) }
func BenchmarkDotuInc_100000_inc2(t *testing.B)  { benchdoti(t, 100000, 2, DotuInc) }
func BenchmarkDotuInc_100000_inc4(t *testing.B)  { benchdoti(t, 100000, 4, DotuInc) }
func BenchmarkDotuInc_100000_inc10(t *testing.B) { benchdoti(t, 100000, 10, DotuInc) }

func BenchmarkDotuInc_100000_incM1(t *testing.B) { benchdoti(t, 100000, -1, DotuInc) }
func BenchmarkDotuInc_100000_incM2(t *testing.B) { benchdoti(t, 100000, -2, DotuInc) }
func BenchmarkDotuInc_100000_incM4(t *testing.B) { benchdoti(t, 100000, -4, DotuInc) }

// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !go1.7

package f64

import (
	"math/rand"
	"testing"
)

var (
	a = float64(2)
	x = make([]float64, 1000000)
	y = make([]float64, 1000000)
	z = make([]float64, 1000000)
)

func init() {
	for n := range x {
		x[n] = float64(n)
		y[n] = float64(n)
	}
}

func benchaxpyu(t *testing.B, n int, f func(a float64, x, y []float64)) {
	x, y := x[:n], y[:n]

	for i := 0; i < t.N; i++ {
		f(a, x, y)
	}
}

func naiveaxpyu(a float64, x, y []float64) {
	for i, v := range x {
		y[i] += a * v
	}
}

func BenchmarkF64AxpyUnitary1(t *testing.B)     { benchaxpyu(t, 1, AxpyUnitary) }
func BenchmarkF64AxpyUnitary2(t *testing.B)     { benchaxpyu(t, 2, AxpyUnitary) }
func BenchmarkF64AxpyUnitary3(t *testing.B)     { benchaxpyu(t, 3, AxpyUnitary) }
func BenchmarkF64AxpyUnitary4(t *testing.B)     { benchaxpyu(t, 4, AxpyUnitary) }
func BenchmarkF64AxpyUnitary5(t *testing.B)     { benchaxpyu(t, 5, AxpyUnitary) }
func BenchmarkF64AxpyUnitary10(t *testing.B)    { benchaxpyu(t, 10, AxpyUnitary) }
func BenchmarkF64AxpyUnitary100(t *testing.B)   { benchaxpyu(t, 100, AxpyUnitary) }
func BenchmarkF64AxpyUnitary1000(t *testing.B)  { benchaxpyu(t, 1000, AxpyUnitary) }
func BenchmarkF64AxpyUnitary5000(t *testing.B)  { benchaxpyu(t, 5000, AxpyUnitary) }
func BenchmarkF64AxpyUnitary10000(t *testing.B) { benchaxpyu(t, 10000, AxpyUnitary) }
func BenchmarkF64AxpyUnitary50000(t *testing.B) { benchaxpyu(t, 50000, AxpyUnitary) }

func BenchmarkLF64AxpyUnitary1(t *testing.B)     { benchaxpyu(t, 1, naiveaxpyu) }
func BenchmarkLF64AxpyUnitary2(t *testing.B)     { benchaxpyu(t, 2, naiveaxpyu) }
func BenchmarkLF64AxpyUnitary3(t *testing.B)     { benchaxpyu(t, 3, naiveaxpyu) }
func BenchmarkLF64AxpyUnitary4(t *testing.B)     { benchaxpyu(t, 4, naiveaxpyu) }
func BenchmarkLF64AxpyUnitary5(t *testing.B)     { benchaxpyu(t, 5, naiveaxpyu) }
func BenchmarkLF64AxpyUnitary10(t *testing.B)    { benchaxpyu(t, 10, naiveaxpyu) }
func BenchmarkLF64AxpyUnitary100(t *testing.B)   { benchaxpyu(t, 100, naiveaxpyu) }
func BenchmarkLF64AxpyUnitary1000(t *testing.B)  { benchaxpyu(t, 1000, naiveaxpyu) }
func BenchmarkLF64AxpyUnitary5000(t *testing.B)  { benchaxpyu(t, 5000, naiveaxpyu) }
func BenchmarkLF64AxpyUnitary10000(t *testing.B) { benchaxpyu(t, 10000, naiveaxpyu) }
func BenchmarkLF64AxpyUnitary50000(t *testing.B) { benchaxpyu(t, 50000, naiveaxpyu) }

func benchaxpyut(t *testing.B, n int, f func(d []float64, a float64, x, y []float64)) {
	x, y, z := x[:n], y[:n], z[:n]

	for i := 0; i < t.N; i++ {
		f(z, a, x, y)
	}
}

func naiveaxpyut(d []float64, a float64, x, y []float64) {
	for i, v := range x {
		d[i] = y[i] + a*v
	}
}

func BenchmarkF64AxpyUnitaryTo1(t *testing.B)     { benchaxpyut(t, 1, AxpyUnitaryTo) }
func BenchmarkF64AxpyUnitaryTo2(t *testing.B)     { benchaxpyut(t, 2, AxpyUnitaryTo) }
func BenchmarkF64AxpyUnitaryTo3(t *testing.B)     { benchaxpyut(t, 3, AxpyUnitaryTo) }
func BenchmarkF64AxpyUnitaryTo4(t *testing.B)     { benchaxpyut(t, 4, AxpyUnitaryTo) }
func BenchmarkF64AxpyUnitaryTo5(t *testing.B)     { benchaxpyut(t, 5, AxpyUnitaryTo) }
func BenchmarkF64AxpyUnitaryTo10(t *testing.B)    { benchaxpyut(t, 10, AxpyUnitaryTo) }
func BenchmarkF64AxpyUnitaryTo100(t *testing.B)   { benchaxpyut(t, 100, AxpyUnitaryTo) }
func BenchmarkF64AxpyUnitaryTo1000(t *testing.B)  { benchaxpyut(t, 1000, AxpyUnitaryTo) }
func BenchmarkF64AxpyUnitaryTo5000(t *testing.B)  { benchaxpyut(t, 5000, AxpyUnitaryTo) }
func BenchmarkF64AxpyUnitaryTo10000(t *testing.B) { benchaxpyut(t, 10000, AxpyUnitaryTo) }
func BenchmarkF64AxpyUnitaryTo50000(t *testing.B) { benchaxpyut(t, 50000, AxpyUnitaryTo) }

func BenchmarkLF64AxpyUnitaryTo1(t *testing.B)     { benchaxpyut(t, 1, naiveaxpyut) }
func BenchmarkLF64AxpyUnitaryTo2(t *testing.B)     { benchaxpyut(t, 2, naiveaxpyut) }
func BenchmarkLF64AxpyUnitaryTo3(t *testing.B)     { benchaxpyut(t, 3, naiveaxpyut) }
func BenchmarkLF64AxpyUnitaryTo4(t *testing.B)     { benchaxpyut(t, 4, naiveaxpyut) }
func BenchmarkLF64AxpyUnitaryTo5(t *testing.B)     { benchaxpyut(t, 5, naiveaxpyut) }
func BenchmarkLF64AxpyUnitaryTo10(t *testing.B)    { benchaxpyut(t, 10, naiveaxpyut) }
func BenchmarkLF64AxpyUnitaryTo100(t *testing.B)   { benchaxpyut(t, 100, naiveaxpyut) }
func BenchmarkLF64AxpyUnitaryTo1000(t *testing.B)  { benchaxpyut(t, 1000, naiveaxpyut) }
func BenchmarkLF64AxpyUnitaryTo5000(t *testing.B)  { benchaxpyut(t, 5000, naiveaxpyut) }
func BenchmarkLF64AxpyUnitaryTo10000(t *testing.B) { benchaxpyut(t, 10000, naiveaxpyut) }
func BenchmarkLF64AxpyUnitaryTo50000(t *testing.B) { benchaxpyut(t, 50000, naiveaxpyut) }

func benchaxpyinc(t *testing.B, ln, t_inc int, f func(alpha float64, x, y []float64, n, incX, incY, ix, iy uintptr)) {
	n, inc := uintptr(ln), uintptr(t_inc)
	var idx int
	if t_inc < 0 {
		idx = (-ln + 1) * t_inc
	}

	for i := 0; i < t.N; i++ {
		f(1, x, y, n, inc, inc, uintptr(idx), uintptr(idx))
	}
}

func naiveaxpyinc(alpha float64, x, y []float64, n, incX, incY, ix, iy uintptr) {
	for i := 0; i < int(n); i++ {
		y[iy] += alpha * x[ix]
		ix += incX
		iy += incY
	}
}

func BenchmarkF64AxpyIncN1Inc1(b *testing.B) { benchaxpyinc(b, 1, 1, AxpyInc) }

func BenchmarkF64AxpyIncN2Inc1(b *testing.B)  { benchaxpyinc(b, 2, 1, AxpyInc) }
func BenchmarkF64AxpyIncN2Inc2(b *testing.B)  { benchaxpyinc(b, 2, 2, AxpyInc) }
func BenchmarkF64AxpyIncN2Inc4(b *testing.B)  { benchaxpyinc(b, 2, 4, AxpyInc) }
func BenchmarkF64AxpyIncN2Inc10(b *testing.B) { benchaxpyinc(b, 2, 10, AxpyInc) }

func BenchmarkF64AxpyIncN3Inc1(b *testing.B)  { benchaxpyinc(b, 3, 1, AxpyInc) }
func BenchmarkF64AxpyIncN3Inc2(b *testing.B)  { benchaxpyinc(b, 3, 2, AxpyInc) }
func BenchmarkF64AxpyIncN3Inc4(b *testing.B)  { benchaxpyinc(b, 3, 4, AxpyInc) }
func BenchmarkF64AxpyIncN3Inc10(b *testing.B) { benchaxpyinc(b, 3, 10, AxpyInc) }

func BenchmarkF64AxpyIncN4Inc1(b *testing.B)  { benchaxpyinc(b, 4, 1, AxpyInc) }
func BenchmarkF64AxpyIncN4Inc2(b *testing.B)  { benchaxpyinc(b, 4, 2, AxpyInc) }
func BenchmarkF64AxpyIncN4Inc4(b *testing.B)  { benchaxpyinc(b, 4, 4, AxpyInc) }
func BenchmarkF64AxpyIncN4Inc10(b *testing.B) { benchaxpyinc(b, 4, 10, AxpyInc) }

func BenchmarkF64AxpyIncN10Inc1(b *testing.B)  { benchaxpyinc(b, 10, 1, AxpyInc) }
func BenchmarkF64AxpyIncN10Inc2(b *testing.B)  { benchaxpyinc(b, 10, 2, AxpyInc) }
func BenchmarkF64AxpyIncN10Inc4(b *testing.B)  { benchaxpyinc(b, 10, 4, AxpyInc) }
func BenchmarkF64AxpyIncN10Inc10(b *testing.B) { benchaxpyinc(b, 10, 10, AxpyInc) }

func BenchmarkF64AxpyIncN1000Inc1(b *testing.B)  { benchaxpyinc(b, 1000, 1, AxpyInc) }
func BenchmarkF64AxpyIncN1000Inc2(b *testing.B)  { benchaxpyinc(b, 1000, 2, AxpyInc) }
func BenchmarkF64AxpyIncN1000Inc4(b *testing.B)  { benchaxpyinc(b, 1000, 4, AxpyInc) }
func BenchmarkF64AxpyIncN1000Inc10(b *testing.B) { benchaxpyinc(b, 1000, 10, AxpyInc) }

func BenchmarkF64AxpyIncN100000Inc1(b *testing.B)  { benchaxpyinc(b, 100000, 1, AxpyInc) }
func BenchmarkF64AxpyIncN100000Inc2(b *testing.B)  { benchaxpyinc(b, 100000, 2, AxpyInc) }
func BenchmarkF64AxpyIncN100000Inc4(b *testing.B)  { benchaxpyinc(b, 100000, 4, AxpyInc) }
func BenchmarkF64AxpyIncN100000Inc10(b *testing.B) { benchaxpyinc(b, 100000, 10, AxpyInc) }

func BenchmarkF64AxpyIncN100000IncM1(b *testing.B)  { benchaxpyinc(b, 100000, -1, AxpyInc) }
func BenchmarkF64AxpyIncN100000IncM2(b *testing.B)  { benchaxpyinc(b, 100000, -2, AxpyInc) }
func BenchmarkF64AxpyIncN100000IncM4(b *testing.B)  { benchaxpyinc(b, 100000, -4, AxpyInc) }
func BenchmarkF64AxpyIncN100000IncM10(b *testing.B) { benchaxpyinc(b, 100000, -10, AxpyInc) }

func BenchmarkLF64AxpyIncN1Inc1(b *testing.B) { benchaxpyinc(b, 1, 1, naiveaxpyinc) }

func BenchmarkLF64AxpyIncN2Inc1(b *testing.B)  { benchaxpyinc(b, 2, 1, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN2Inc2(b *testing.B)  { benchaxpyinc(b, 2, 2, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN2Inc4(b *testing.B)  { benchaxpyinc(b, 2, 4, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN2Inc10(b *testing.B) { benchaxpyinc(b, 2, 10, naiveaxpyinc) }

func BenchmarkLF64AxpyIncN3Inc1(b *testing.B)  { benchaxpyinc(b, 3, 1, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN3Inc2(b *testing.B)  { benchaxpyinc(b, 3, 2, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN3Inc4(b *testing.B)  { benchaxpyinc(b, 3, 4, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN3Inc10(b *testing.B) { benchaxpyinc(b, 3, 10, naiveaxpyinc) }

func BenchmarkLF64AxpyIncN4Inc1(b *testing.B)  { benchaxpyinc(b, 4, 1, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN4Inc2(b *testing.B)  { benchaxpyinc(b, 4, 2, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN4Inc4(b *testing.B)  { benchaxpyinc(b, 4, 4, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN4Inc10(b *testing.B) { benchaxpyinc(b, 4, 10, naiveaxpyinc) }

func BenchmarkLF64AxpyIncN10Inc1(b *testing.B)  { benchaxpyinc(b, 10, 1, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN10Inc2(b *testing.B)  { benchaxpyinc(b, 10, 2, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN10Inc4(b *testing.B)  { benchaxpyinc(b, 10, 4, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN10Inc10(b *testing.B) { benchaxpyinc(b, 10, 10, naiveaxpyinc) }

func BenchmarkLF64AxpyIncN1000Inc1(b *testing.B)  { benchaxpyinc(b, 1000, 1, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN1000Inc2(b *testing.B)  { benchaxpyinc(b, 1000, 2, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN1000Inc4(b *testing.B)  { benchaxpyinc(b, 1000, 4, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN1000Inc10(b *testing.B) { benchaxpyinc(b, 1000, 10, naiveaxpyinc) }

func BenchmarkLF64AxpyIncN100000Inc1(b *testing.B)  { benchaxpyinc(b, 100000, 1, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN100000Inc2(b *testing.B)  { benchaxpyinc(b, 100000, 2, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN100000Inc4(b *testing.B)  { benchaxpyinc(b, 100000, 4, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN100000Inc10(b *testing.B) { benchaxpyinc(b, 100000, 10, naiveaxpyinc) }

func BenchmarkLF64AxpyIncN100000IncM1(b *testing.B)  { benchaxpyinc(b, 100000, -1, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN100000IncM2(b *testing.B)  { benchaxpyinc(b, 100000, -2, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN100000IncM4(b *testing.B)  { benchaxpyinc(b, 100000, -4, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN100000IncM10(b *testing.B) { benchaxpyinc(b, 100000, -10, naiveaxpyinc) }

func benchaxpyincto(t *testing.B, ln, t_inc int, f func(dst []float64, incDst, idst uintptr, alpha float64, x, y []float64, n, incX, incY, ix, iy uintptr)) {
	n, inc := uintptr(ln), uintptr(t_inc)
	var idx int
	if t_inc < 0 {
		idx = (-ln + 1) * t_inc
	}

	for i := 0; i < t.N; i++ {
		f(z, inc, uintptr(idx), 1, x, y, n, inc, inc, uintptr(idx), uintptr(idx))
	}
}

func naiveaxpyincto(dst []float64, incDst, idst uintptr, alpha float64, x, y []float64, n, incX, incY, ix, iy uintptr) {
	for i := 0; i < int(n); i++ {
		dst[idst] = alpha*x[ix] + y[iy]
		ix += incX
		iy += incY
		idst += incDst
	}
}

func BenchmarkF64AxpyIncToN1Inc1(b *testing.B) { benchaxpyincto(b, 1, 1, AxpyIncTo) }

func BenchmarkF64AxpyIncToN2Inc1(b *testing.B)  { benchaxpyincto(b, 2, 1, AxpyIncTo) }
func BenchmarkF64AxpyIncToN2Inc2(b *testing.B)  { benchaxpyincto(b, 2, 2, AxpyIncTo) }
func BenchmarkF64AxpyIncToN2Inc4(b *testing.B)  { benchaxpyincto(b, 2, 4, AxpyIncTo) }
func BenchmarkF64AxpyIncToN2Inc10(b *testing.B) { benchaxpyincto(b, 2, 10, AxpyIncTo) }

func BenchmarkF64AxpyIncToN3Inc1(b *testing.B)  { benchaxpyincto(b, 3, 1, AxpyIncTo) }
func BenchmarkF64AxpyIncToN3Inc2(b *testing.B)  { benchaxpyincto(b, 3, 2, AxpyIncTo) }
func BenchmarkF64AxpyIncToN3Inc4(b *testing.B)  { benchaxpyincto(b, 3, 4, AxpyIncTo) }
func BenchmarkF64AxpyIncToN3Inc10(b *testing.B) { benchaxpyincto(b, 3, 10, AxpyIncTo) }

func BenchmarkF64AxpyIncToN4Inc1(b *testing.B)  { benchaxpyincto(b, 4, 1, AxpyIncTo) }
func BenchmarkF64AxpyIncToN4Inc2(b *testing.B)  { benchaxpyincto(b, 4, 2, AxpyIncTo) }
func BenchmarkF64AxpyIncToN4Inc4(b *testing.B)  { benchaxpyincto(b, 4, 4, AxpyIncTo) }
func BenchmarkF64AxpyIncToN4Inc10(b *testing.B) { benchaxpyincto(b, 4, 10, AxpyIncTo) }

func BenchmarkF64AxpyIncToN10Inc1(b *testing.B)  { benchaxpyincto(b, 10, 1, AxpyIncTo) }
func BenchmarkF64AxpyIncToN10Inc2(b *testing.B)  { benchaxpyincto(b, 10, 2, AxpyIncTo) }
func BenchmarkF64AxpyIncToN10Inc4(b *testing.B)  { benchaxpyincto(b, 10, 4, AxpyIncTo) }
func BenchmarkF64AxpyIncToN10Inc10(b *testing.B) { benchaxpyincto(b, 10, 10, AxpyIncTo) }

func BenchmarkF64AxpyIncToN1000Inc1(b *testing.B)  { benchaxpyincto(b, 1000, 1, AxpyIncTo) }
func BenchmarkF64AxpyIncToN1000Inc2(b *testing.B)  { benchaxpyincto(b, 1000, 2, AxpyIncTo) }
func BenchmarkF64AxpyIncToN1000Inc4(b *testing.B)  { benchaxpyincto(b, 1000, 4, AxpyIncTo) }
func BenchmarkF64AxpyIncToN1000Inc10(b *testing.B) { benchaxpyincto(b, 1000, 10, AxpyIncTo) }

func BenchmarkF64AxpyIncToN100000Inc1(b *testing.B)  { benchaxpyincto(b, 100000, 1, AxpyIncTo) }
func BenchmarkF64AxpyIncToN100000Inc2(b *testing.B)  { benchaxpyincto(b, 100000, 2, AxpyIncTo) }
func BenchmarkF64AxpyIncToN100000Inc4(b *testing.B)  { benchaxpyincto(b, 100000, 4, AxpyIncTo) }
func BenchmarkF64AxpyIncToN100000Inc10(b *testing.B) { benchaxpyincto(b, 100000, 10, AxpyIncTo) }

func BenchmarkF64AxpyIncToN100000IncM1(b *testing.B)  { benchaxpyincto(b, 100000, -1, AxpyIncTo) }
func BenchmarkF64AxpyIncToN100000IncM2(b *testing.B)  { benchaxpyincto(b, 100000, -2, AxpyIncTo) }
func BenchmarkF64AxpyIncToN100000IncM4(b *testing.B)  { benchaxpyincto(b, 100000, -4, AxpyIncTo) }
func BenchmarkF64AxpyIncToN100000IncM10(b *testing.B) { benchaxpyincto(b, 100000, -10, AxpyIncTo) }

func BenchmarkLF64AxpyIncToN1Inc1(b *testing.B) { benchaxpyincto(b, 1, 1, naiveaxpyincto) }

func BenchmarkLF64AxpyIncToN2Inc1(b *testing.B)  { benchaxpyincto(b, 2, 1, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN2Inc2(b *testing.B)  { benchaxpyincto(b, 2, 2, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN2Inc4(b *testing.B)  { benchaxpyincto(b, 2, 4, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN2Inc10(b *testing.B) { benchaxpyincto(b, 2, 10, naiveaxpyincto) }

func BenchmarkLF64AxpyIncToN3Inc1(b *testing.B)  { benchaxpyincto(b, 3, 1, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN3Inc2(b *testing.B)  { benchaxpyincto(b, 3, 2, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN3Inc4(b *testing.B)  { benchaxpyincto(b, 3, 4, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN3Inc10(b *testing.B) { benchaxpyincto(b, 3, 10, naiveaxpyincto) }

func BenchmarkLF64AxpyIncToN4Inc1(b *testing.B)  { benchaxpyincto(b, 4, 1, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN4Inc2(b *testing.B)  { benchaxpyincto(b, 4, 2, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN4Inc4(b *testing.B)  { benchaxpyincto(b, 4, 4, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN4Inc10(b *testing.B) { benchaxpyincto(b, 4, 10, naiveaxpyincto) }

func BenchmarkLF64AxpyIncToN10Inc1(b *testing.B)  { benchaxpyincto(b, 10, 1, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN10Inc2(b *testing.B)  { benchaxpyincto(b, 10, 2, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN10Inc4(b *testing.B)  { benchaxpyincto(b, 10, 4, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN10Inc10(b *testing.B) { benchaxpyincto(b, 10, 10, naiveaxpyincto) }

func BenchmarkLF64AxpyIncToN1000Inc1(b *testing.B)  { benchaxpyincto(b, 1000, 1, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN1000Inc2(b *testing.B)  { benchaxpyincto(b, 1000, 2, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN1000Inc4(b *testing.B)  { benchaxpyincto(b, 1000, 4, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN1000Inc10(b *testing.B) { benchaxpyincto(b, 1000, 10, naiveaxpyincto) }

func BenchmarkLF64AxpyIncToN100000Inc1(b *testing.B)  { benchaxpyincto(b, 100000, 1, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN100000Inc2(b *testing.B)  { benchaxpyincto(b, 100000, 2, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN100000Inc4(b *testing.B)  { benchaxpyincto(b, 100000, 4, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN100000Inc10(b *testing.B) { benchaxpyincto(b, 100000, 10, naiveaxpyincto) }

func BenchmarkLF64AxpyIncToN100000IncM1(b *testing.B)  { benchaxpyincto(b, 100000, -1, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN100000IncM2(b *testing.B)  { benchaxpyincto(b, 100000, -2, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN100000IncM4(b *testing.B)  { benchaxpyincto(b, 100000, -4, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN100000IncM10(b *testing.B) { benchaxpyincto(b, 100000, -10, naiveaxpyincto) }

// Scal* benchmarks
func BenchmarkDscalUnitaryN1(b *testing.B)      { benchmarkDscalUnitary(b, 1) }
func BenchmarkDscalUnitaryN2(b *testing.B)      { benchmarkDscalUnitary(b, 2) }
func BenchmarkDscalUnitaryN3(b *testing.B)      { benchmarkDscalUnitary(b, 3) }
func BenchmarkDscalUnitaryN4(b *testing.B)      { benchmarkDscalUnitary(b, 4) }
func BenchmarkDscalUnitaryN10(b *testing.B)     { benchmarkDscalUnitary(b, 10) }
func BenchmarkDscalUnitaryN100(b *testing.B)    { benchmarkDscalUnitary(b, 100) }
func BenchmarkDscalUnitaryN1000(b *testing.B)   { benchmarkDscalUnitary(b, 1000) }
func BenchmarkDscalUnitaryN10000(b *testing.B)  { benchmarkDscalUnitary(b, 10000) }
func BenchmarkDscalUnitaryN100000(b *testing.B) { benchmarkDscalUnitary(b, 100000) }

func benchmarkDscalUnitary(b *testing.B, n int) {
	x := randomSlice(n, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i += 2 {
		ScalUnitary(2, x)
		ScalUnitary(0.5, x)
	}
	benchSink = x
}

func BenchmarkDscalUnitaryToN1(b *testing.B)      { benchmarkDscalUnitaryTo(b, 1) }
func BenchmarkDscalUnitaryToN2(b *testing.B)      { benchmarkDscalUnitaryTo(b, 2) }
func BenchmarkDscalUnitaryToN3(b *testing.B)      { benchmarkDscalUnitaryTo(b, 3) }
func BenchmarkDscalUnitaryToN4(b *testing.B)      { benchmarkDscalUnitaryTo(b, 4) }
func BenchmarkDscalUnitaryToN10(b *testing.B)     { benchmarkDscalUnitaryTo(b, 10) }
func BenchmarkDscalUnitaryToN100(b *testing.B)    { benchmarkDscalUnitaryTo(b, 100) }
func BenchmarkDscalUnitaryToN1000(b *testing.B)   { benchmarkDscalUnitaryTo(b, 1000) }
func BenchmarkDscalUnitaryToN10000(b *testing.B)  { benchmarkDscalUnitaryTo(b, 10000) }
func BenchmarkDscalUnitaryToN100000(b *testing.B) { benchmarkDscalUnitaryTo(b, 100000) }

func benchmarkDscalUnitaryTo(b *testing.B, n int) {
	x := randomSlice(n, 1)
	dst := randomSlice(n, 1)
	a := rand.Float64()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ScalUnitaryTo(dst, a, x)
	}
	benchSink = dst
}

func BenchmarkDscalUnitaryToXN1(b *testing.B)      { benchmarkDscalUnitaryToX(b, 1) }
func BenchmarkDscalUnitaryToXN2(b *testing.B)      { benchmarkDscalUnitaryToX(b, 2) }
func BenchmarkDscalUnitaryToXN3(b *testing.B)      { benchmarkDscalUnitaryToX(b, 3) }
func BenchmarkDscalUnitaryToXN4(b *testing.B)      { benchmarkDscalUnitaryToX(b, 4) }
func BenchmarkDscalUnitaryToXN10(b *testing.B)     { benchmarkDscalUnitaryToX(b, 10) }
func BenchmarkDscalUnitaryToXN100(b *testing.B)    { benchmarkDscalUnitaryToX(b, 100) }
func BenchmarkDscalUnitaryToXN1000(b *testing.B)   { benchmarkDscalUnitaryToX(b, 1000) }
func BenchmarkDscalUnitaryToXN10000(b *testing.B)  { benchmarkDscalUnitaryToX(b, 10000) }
func BenchmarkDscalUnitaryToXN100000(b *testing.B) { benchmarkDscalUnitaryToX(b, 100000) }

func benchmarkDscalUnitaryToX(b *testing.B, n int) {
	x := randomSlice(n, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i += 2 {
		ScalUnitaryTo(x, 2, x)
		ScalUnitaryTo(x, 0.5, x)
	}
	benchSink = x
}

func BenchmarkDscalIncN1Inc1(b *testing.B) { benchmarkDscalInc(b, 1, 1) }

func BenchmarkDscalIncN2Inc1(b *testing.B)  { benchmarkDscalInc(b, 2, 1) }
func BenchmarkDscalIncN2Inc2(b *testing.B)  { benchmarkDscalInc(b, 2, 2) }
func BenchmarkDscalIncN2Inc4(b *testing.B)  { benchmarkDscalInc(b, 2, 4) }
func BenchmarkDscalIncN2Inc10(b *testing.B) { benchmarkDscalInc(b, 2, 10) }

func BenchmarkDscalIncN3Inc1(b *testing.B)  { benchmarkDscalInc(b, 3, 1) }
func BenchmarkDscalIncN3Inc2(b *testing.B)  { benchmarkDscalInc(b, 3, 2) }
func BenchmarkDscalIncN3Inc4(b *testing.B)  { benchmarkDscalInc(b, 3, 4) }
func BenchmarkDscalIncN3Inc10(b *testing.B) { benchmarkDscalInc(b, 3, 10) }

func BenchmarkDscalIncN4Inc1(b *testing.B)  { benchmarkDscalInc(b, 4, 1) }
func BenchmarkDscalIncN4Inc2(b *testing.B)  { benchmarkDscalInc(b, 4, 2) }
func BenchmarkDscalIncN4Inc4(b *testing.B)  { benchmarkDscalInc(b, 4, 4) }
func BenchmarkDscalIncN4Inc10(b *testing.B) { benchmarkDscalInc(b, 4, 10) }

func BenchmarkDscalIncN10Inc1(b *testing.B)  { benchmarkDscalInc(b, 10, 1) }
func BenchmarkDscalIncN10Inc2(b *testing.B)  { benchmarkDscalInc(b, 10, 2) }
func BenchmarkDscalIncN10Inc4(b *testing.B)  { benchmarkDscalInc(b, 10, 4) }
func BenchmarkDscalIncN10Inc10(b *testing.B) { benchmarkDscalInc(b, 10, 10) }

func BenchmarkDscalIncN1000Inc1(b *testing.B)  { benchmarkDscalInc(b, 1000, 1) }
func BenchmarkDscalIncN1000Inc2(b *testing.B)  { benchmarkDscalInc(b, 1000, 2) }
func BenchmarkDscalIncN1000Inc4(b *testing.B)  { benchmarkDscalInc(b, 1000, 4) }
func BenchmarkDscalIncN1000Inc10(b *testing.B) { benchmarkDscalInc(b, 1000, 10) }

func BenchmarkDscalIncN100000Inc1(b *testing.B)  { benchmarkDscalInc(b, 100000, 1) }
func BenchmarkDscalIncN100000Inc2(b *testing.B)  { benchmarkDscalInc(b, 100000, 2) }
func BenchmarkDscalIncN100000Inc4(b *testing.B)  { benchmarkDscalInc(b, 100000, 4) }
func BenchmarkDscalIncN100000Inc10(b *testing.B) { benchmarkDscalInc(b, 100000, 10) }

func benchmarkDscalInc(b *testing.B, n, inc int) {
	x := randomSlice(n, inc)
	b.ResetTimer()
	for i := 0; i < b.N; i += 2 {
		ScalInc(2, x, uintptr(n), uintptr(inc))
		ScalInc(0.5, x, uintptr(n), uintptr(inc))
	}
	benchSink = x
}

func BenchmarkDscalIncToN1Inc1(b *testing.B) { benchmarkDscalIncTo(b, 1, 1) }

func BenchmarkDscalIncToN2Inc1(b *testing.B)  { benchmarkDscalIncTo(b, 2, 1) }
func BenchmarkDscalIncToN2Inc2(b *testing.B)  { benchmarkDscalIncTo(b, 2, 2) }
func BenchmarkDscalIncToN2Inc4(b *testing.B)  { benchmarkDscalIncTo(b, 2, 4) }
func BenchmarkDscalIncToN2Inc10(b *testing.B) { benchmarkDscalIncTo(b, 2, 10) }

func BenchmarkDscalIncToN3Inc1(b *testing.B)  { benchmarkDscalIncTo(b, 3, 1) }
func BenchmarkDscalIncToN3Inc2(b *testing.B)  { benchmarkDscalIncTo(b, 3, 2) }
func BenchmarkDscalIncToN3Inc4(b *testing.B)  { benchmarkDscalIncTo(b, 3, 4) }
func BenchmarkDscalIncToN3Inc10(b *testing.B) { benchmarkDscalIncTo(b, 3, 10) }

func BenchmarkDscalIncToN4Inc1(b *testing.B)  { benchmarkDscalIncTo(b, 4, 1) }
func BenchmarkDscalIncToN4Inc2(b *testing.B)  { benchmarkDscalIncTo(b, 4, 2) }
func BenchmarkDscalIncToN4Inc4(b *testing.B)  { benchmarkDscalIncTo(b, 4, 4) }
func BenchmarkDscalIncToN4Inc10(b *testing.B) { benchmarkDscalIncTo(b, 4, 10) }

func BenchmarkDscalIncToN10Inc1(b *testing.B)  { benchmarkDscalIncTo(b, 10, 1) }
func BenchmarkDscalIncToN10Inc2(b *testing.B)  { benchmarkDscalIncTo(b, 10, 2) }
func BenchmarkDscalIncToN10Inc4(b *testing.B)  { benchmarkDscalIncTo(b, 10, 4) }
func BenchmarkDscalIncToN10Inc10(b *testing.B) { benchmarkDscalIncTo(b, 10, 10) }

func BenchmarkDscalIncToN1000Inc1(b *testing.B)  { benchmarkDscalIncTo(b, 1000, 1) }
func BenchmarkDscalIncToN1000Inc2(b *testing.B)  { benchmarkDscalIncTo(b, 1000, 2) }
func BenchmarkDscalIncToN1000Inc4(b *testing.B)  { benchmarkDscalIncTo(b, 1000, 4) }
func BenchmarkDscalIncToN1000Inc10(b *testing.B) { benchmarkDscalIncTo(b, 1000, 10) }

func BenchmarkDscalIncToN100000Inc1(b *testing.B)  { benchmarkDscalIncTo(b, 100000, 1) }
func BenchmarkDscalIncToN100000Inc2(b *testing.B)  { benchmarkDscalIncTo(b, 100000, 2) }
func BenchmarkDscalIncToN100000Inc4(b *testing.B)  { benchmarkDscalIncTo(b, 100000, 4) }
func BenchmarkDscalIncToN100000Inc10(b *testing.B) { benchmarkDscalIncTo(b, 100000, 10) }

func benchmarkDscalIncTo(b *testing.B, n, inc int) {
	x := randomSlice(n, inc)
	dst := randomSlice(n, inc)
	a := rand.Float64()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ScalIncTo(dst, uintptr(inc), a, x, uintptr(n), uintptr(inc))
	}
	benchSink = dst
}

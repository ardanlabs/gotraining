// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package f32

import "testing"

var (
	a = float32(2)
	x = make([]float32, 1000000)
	y = make([]float32, 1000000)
	z = make([]float32, 1000000)
)

func init() {
	for n := range x {
		x[n] = float32(n)
		y[n] = float32(n)
	}
}

func benchaxpyu(t *testing.B, n int, f func(a float32, x, y []float32)) {
	x, y := x[:n], y[:n]
	for i := 0; i < t.N; i++ {
		f(a, x, y)
	}
}

func naiveaxpyu(a float32, x, y []float32) {
	for i, v := range x {
		y[i] += a * v
	}
}

func BenchmarkF32AxpyUnitary1(t *testing.B)     { benchaxpyu(t, 1, AxpyUnitary) }
func BenchmarkF32AxpyUnitary2(t *testing.B)     { benchaxpyu(t, 2, AxpyUnitary) }
func BenchmarkF32AxpyUnitary3(t *testing.B)     { benchaxpyu(t, 3, AxpyUnitary) }
func BenchmarkF32AxpyUnitary4(t *testing.B)     { benchaxpyu(t, 4, AxpyUnitary) }
func BenchmarkF32AxpyUnitary5(t *testing.B)     { benchaxpyu(t, 5, AxpyUnitary) }
func BenchmarkF32AxpyUnitary10(t *testing.B)    { benchaxpyu(t, 10, AxpyUnitary) }
func BenchmarkF32AxpyUnitary100(t *testing.B)   { benchaxpyu(t, 100, AxpyUnitary) }
func BenchmarkF32AxpyUnitary1000(t *testing.B)  { benchaxpyu(t, 1000, AxpyUnitary) }
func BenchmarkF32AxpyUnitary5000(t *testing.B)  { benchaxpyu(t, 5000, AxpyUnitary) }
func BenchmarkF32AxpyUnitary10000(t *testing.B) { benchaxpyu(t, 10000, AxpyUnitary) }
func BenchmarkF32AxpyUnitary50000(t *testing.B) { benchaxpyu(t, 50000, AxpyUnitary) }

func BenchmarkLF32AxpyUnitary1(t *testing.B)     { benchaxpyu(t, 1, naiveaxpyu) }
func BenchmarkLF32AxpyUnitary2(t *testing.B)     { benchaxpyu(t, 2, naiveaxpyu) }
func BenchmarkLF32AxpyUnitary3(t *testing.B)     { benchaxpyu(t, 3, naiveaxpyu) }
func BenchmarkLF32AxpyUnitary4(t *testing.B)     { benchaxpyu(t, 4, naiveaxpyu) }
func BenchmarkLF32AxpyUnitary5(t *testing.B)     { benchaxpyu(t, 5, naiveaxpyu) }
func BenchmarkLF32AxpyUnitary10(t *testing.B)    { benchaxpyu(t, 10, naiveaxpyu) }
func BenchmarkLF32AxpyUnitary100(t *testing.B)   { benchaxpyu(t, 100, naiveaxpyu) }
func BenchmarkLF32AxpyUnitary1000(t *testing.B)  { benchaxpyu(t, 1000, naiveaxpyu) }
func BenchmarkLF32AxpyUnitary5000(t *testing.B)  { benchaxpyu(t, 5000, naiveaxpyu) }
func BenchmarkLF32AxpyUnitary10000(t *testing.B) { benchaxpyu(t, 10000, naiveaxpyu) }
func BenchmarkLF32AxpyUnitary50000(t *testing.B) { benchaxpyu(t, 50000, naiveaxpyu) }

func benchaxpyut(t *testing.B, n int, f func(d []float32, a float32, x, y []float32)) {
	x, y, z := x[:n], y[:n], z[:n]
	for i := 0; i < t.N; i++ {
		f(z, a, x, y)
	}
}

func naiveaxpyut(d []float32, a float32, x, y []float32) {
	for i, v := range x {
		d[i] = y[i] + a*v
	}
}

func BenchmarkF32AxpyUnitaryTo1(t *testing.B)     { benchaxpyut(t, 1, AxpyUnitaryTo) }
func BenchmarkF32AxpyUnitaryTo2(t *testing.B)     { benchaxpyut(t, 2, AxpyUnitaryTo) }
func BenchmarkF32AxpyUnitaryTo3(t *testing.B)     { benchaxpyut(t, 3, AxpyUnitaryTo) }
func BenchmarkF32AxpyUnitaryTo4(t *testing.B)     { benchaxpyut(t, 4, AxpyUnitaryTo) }
func BenchmarkF32AxpyUnitaryTo5(t *testing.B)     { benchaxpyut(t, 5, AxpyUnitaryTo) }
func BenchmarkF32AxpyUnitaryTo10(t *testing.B)    { benchaxpyut(t, 10, AxpyUnitaryTo) }
func BenchmarkF32AxpyUnitaryTo100(t *testing.B)   { benchaxpyut(t, 100, AxpyUnitaryTo) }
func BenchmarkF32AxpyUnitaryTo1000(t *testing.B)  { benchaxpyut(t, 1000, AxpyUnitaryTo) }
func BenchmarkF32AxpyUnitaryTo5000(t *testing.B)  { benchaxpyut(t, 5000, AxpyUnitaryTo) }
func BenchmarkF32AxpyUnitaryTo10000(t *testing.B) { benchaxpyut(t, 10000, AxpyUnitaryTo) }
func BenchmarkF32AxpyUnitaryTo50000(t *testing.B) { benchaxpyut(t, 50000, AxpyUnitaryTo) }

func BenchmarkLF32AxpyUnitaryTo1(t *testing.B)     { benchaxpyut(t, 1, naiveaxpyut) }
func BenchmarkLF32AxpyUnitaryTo2(t *testing.B)     { benchaxpyut(t, 2, naiveaxpyut) }
func BenchmarkLF32AxpyUnitaryTo3(t *testing.B)     { benchaxpyut(t, 3, naiveaxpyut) }
func BenchmarkLF32AxpyUnitaryTo4(t *testing.B)     { benchaxpyut(t, 4, naiveaxpyut) }
func BenchmarkLF32AxpyUnitaryTo5(t *testing.B)     { benchaxpyut(t, 5, naiveaxpyut) }
func BenchmarkLF32AxpyUnitaryTo10(t *testing.B)    { benchaxpyut(t, 10, naiveaxpyut) }
func BenchmarkLF32AxpyUnitaryTo100(t *testing.B)   { benchaxpyut(t, 100, naiveaxpyut) }
func BenchmarkLF32AxpyUnitaryTo1000(t *testing.B)  { benchaxpyut(t, 1000, naiveaxpyut) }
func BenchmarkLF32AxpyUnitaryTo5000(t *testing.B)  { benchaxpyut(t, 5000, naiveaxpyut) }
func BenchmarkLF32AxpyUnitaryTo10000(t *testing.B) { benchaxpyut(t, 10000, naiveaxpyut) }
func BenchmarkLF32AxpyUnitaryTo50000(t *testing.B) { benchaxpyut(t, 50000, naiveaxpyut) }

func benchaxpyinc(t *testing.B, ln, t_inc int, f func(alpha float32, x, y []float32, n, incX, incY, ix, iy uintptr)) {
	n, inc := uintptr(ln), uintptr(t_inc)
	var idx int
	if t_inc < 0 {
		idx = (-ln + 1) * t_inc
	}
	for i := 0; i < t.N; i++ {
		f(1, x, y, n, inc, inc, uintptr(idx), uintptr(idx))
	}
}

func naiveaxpyinc(alpha float32, x, y []float32, n, incX, incY, ix, iy uintptr) {
	for i := 0; i < int(n); i++ {
		y[iy] += alpha * x[ix]
		ix += incX
		iy += incY
	}
}

func BenchmarkF32AxpyIncN1Inc1(b *testing.B) { benchaxpyinc(b, 1, 1, AxpyInc) }

func BenchmarkF32AxpyIncN2Inc1(b *testing.B)  { benchaxpyinc(b, 2, 1, AxpyInc) }
func BenchmarkF32AxpyIncN2Inc2(b *testing.B)  { benchaxpyinc(b, 2, 2, AxpyInc) }
func BenchmarkF32AxpyIncN2Inc4(b *testing.B)  { benchaxpyinc(b, 2, 4, AxpyInc) }
func BenchmarkF32AxpyIncN2Inc10(b *testing.B) { benchaxpyinc(b, 2, 10, AxpyInc) }

func BenchmarkF32AxpyIncN3Inc1(b *testing.B)  { benchaxpyinc(b, 3, 1, AxpyInc) }
func BenchmarkF32AxpyIncN3Inc2(b *testing.B)  { benchaxpyinc(b, 3, 2, AxpyInc) }
func BenchmarkF32AxpyIncN3Inc4(b *testing.B)  { benchaxpyinc(b, 3, 4, AxpyInc) }
func BenchmarkF32AxpyIncN3Inc10(b *testing.B) { benchaxpyinc(b, 3, 10, AxpyInc) }

func BenchmarkF32AxpyIncN4Inc1(b *testing.B)  { benchaxpyinc(b, 4, 1, AxpyInc) }
func BenchmarkF32AxpyIncN4Inc2(b *testing.B)  { benchaxpyinc(b, 4, 2, AxpyInc) }
func BenchmarkF32AxpyIncN4Inc4(b *testing.B)  { benchaxpyinc(b, 4, 4, AxpyInc) }
func BenchmarkF32AxpyIncN4Inc10(b *testing.B) { benchaxpyinc(b, 4, 10, AxpyInc) }

func BenchmarkF32AxpyIncN10Inc1(b *testing.B)  { benchaxpyinc(b, 10, 1, AxpyInc) }
func BenchmarkF32AxpyIncN10Inc2(b *testing.B)  { benchaxpyinc(b, 10, 2, AxpyInc) }
func BenchmarkF32AxpyIncN10Inc4(b *testing.B)  { benchaxpyinc(b, 10, 4, AxpyInc) }
func BenchmarkF32AxpyIncN10Inc10(b *testing.B) { benchaxpyinc(b, 10, 10, AxpyInc) }

func BenchmarkF32AxpyIncN1000Inc1(b *testing.B)  { benchaxpyinc(b, 1000, 1, AxpyInc) }
func BenchmarkF32AxpyIncN1000Inc2(b *testing.B)  { benchaxpyinc(b, 1000, 2, AxpyInc) }
func BenchmarkF32AxpyIncN1000Inc4(b *testing.B)  { benchaxpyinc(b, 1000, 4, AxpyInc) }
func BenchmarkF32AxpyIncN1000Inc10(b *testing.B) { benchaxpyinc(b, 1000, 10, AxpyInc) }

func BenchmarkF32AxpyIncN100000Inc1(b *testing.B)  { benchaxpyinc(b, 100000, 1, AxpyInc) }
func BenchmarkF32AxpyIncN100000Inc2(b *testing.B)  { benchaxpyinc(b, 100000, 2, AxpyInc) }
func BenchmarkF32AxpyIncN100000Inc4(b *testing.B)  { benchaxpyinc(b, 100000, 4, AxpyInc) }
func BenchmarkF32AxpyIncN100000Inc10(b *testing.B) { benchaxpyinc(b, 100000, 10, AxpyInc) }

func BenchmarkF32AxpyIncN100000IncM1(b *testing.B)  { benchaxpyinc(b, 100000, -1, AxpyInc) }
func BenchmarkF32AxpyIncN100000IncM2(b *testing.B)  { benchaxpyinc(b, 100000, -2, AxpyInc) }
func BenchmarkF32AxpyIncN100000IncM4(b *testing.B)  { benchaxpyinc(b, 100000, -4, AxpyInc) }
func BenchmarkF32AxpyIncN100000IncM10(b *testing.B) { benchaxpyinc(b, 100000, -10, AxpyInc) }

func BenchmarkLF32AxpyIncN1Inc1(b *testing.B) { benchaxpyinc(b, 1, 1, naiveaxpyinc) }

func BenchmarkLF32AxpyIncN2Inc1(b *testing.B)  { benchaxpyinc(b, 2, 1, naiveaxpyinc) }
func BenchmarkLF32AxpyIncN2Inc2(b *testing.B)  { benchaxpyinc(b, 2, 2, naiveaxpyinc) }
func BenchmarkLF32AxpyIncN2Inc4(b *testing.B)  { benchaxpyinc(b, 2, 4, naiveaxpyinc) }
func BenchmarkLF32AxpyIncN2Inc10(b *testing.B) { benchaxpyinc(b, 2, 10, naiveaxpyinc) }

func BenchmarkLF32AxpyIncN3Inc1(b *testing.B)  { benchaxpyinc(b, 3, 1, naiveaxpyinc) }
func BenchmarkLF32AxpyIncN3Inc2(b *testing.B)  { benchaxpyinc(b, 3, 2, naiveaxpyinc) }
func BenchmarkLF32AxpyIncN3Inc4(b *testing.B)  { benchaxpyinc(b, 3, 4, naiveaxpyinc) }
func BenchmarkLF32AxpyIncN3Inc10(b *testing.B) { benchaxpyinc(b, 3, 10, naiveaxpyinc) }

func BenchmarkLF32AxpyIncN4Inc1(b *testing.B)  { benchaxpyinc(b, 4, 1, naiveaxpyinc) }
func BenchmarkLF32AxpyIncN4Inc2(b *testing.B)  { benchaxpyinc(b, 4, 2, naiveaxpyinc) }
func BenchmarkLF32AxpyIncN4Inc4(b *testing.B)  { benchaxpyinc(b, 4, 4, naiveaxpyinc) }
func BenchmarkLF32AxpyIncN4Inc10(b *testing.B) { benchaxpyinc(b, 4, 10, naiveaxpyinc) }

func BenchmarkLF32AxpyIncN10Inc1(b *testing.B)  { benchaxpyinc(b, 10, 1, naiveaxpyinc) }
func BenchmarkLF32AxpyIncN10Inc2(b *testing.B)  { benchaxpyinc(b, 10, 2, naiveaxpyinc) }
func BenchmarkLF32AxpyIncN10Inc4(b *testing.B)  { benchaxpyinc(b, 10, 4, naiveaxpyinc) }
func BenchmarkLF32AxpyIncN10Inc10(b *testing.B) { benchaxpyinc(b, 10, 10, naiveaxpyinc) }

func BenchmarkLF32AxpyIncN1000Inc1(b *testing.B)  { benchaxpyinc(b, 1000, 1, naiveaxpyinc) }
func BenchmarkLF32AxpyIncN1000Inc2(b *testing.B)  { benchaxpyinc(b, 1000, 2, naiveaxpyinc) }
func BenchmarkLF32AxpyIncN1000Inc4(b *testing.B)  { benchaxpyinc(b, 1000, 4, naiveaxpyinc) }
func BenchmarkLF32AxpyIncN1000Inc10(b *testing.B) { benchaxpyinc(b, 1000, 10, naiveaxpyinc) }

func BenchmarkLF32AxpyIncN100000Inc1(b *testing.B)  { benchaxpyinc(b, 100000, 1, naiveaxpyinc) }
func BenchmarkLF32AxpyIncN100000Inc2(b *testing.B)  { benchaxpyinc(b, 100000, 2, naiveaxpyinc) }
func BenchmarkLF32AxpyIncN100000Inc4(b *testing.B)  { benchaxpyinc(b, 100000, 4, naiveaxpyinc) }
func BenchmarkLF32AxpyIncN100000Inc10(b *testing.B) { benchaxpyinc(b, 100000, 10, naiveaxpyinc) }

func BenchmarkLF32AxpyIncN100000IncM1(b *testing.B)  { benchaxpyinc(b, 100000, -1, naiveaxpyinc) }
func BenchmarkLF32AxpyIncN100000IncM2(b *testing.B)  { benchaxpyinc(b, 100000, -2, naiveaxpyinc) }
func BenchmarkLF32AxpyIncN100000IncM4(b *testing.B)  { benchaxpyinc(b, 100000, -4, naiveaxpyinc) }
func BenchmarkLF32AxpyIncN100000IncM10(b *testing.B) { benchaxpyinc(b, 100000, -10, naiveaxpyinc) }

func benchaxpyincto(t *testing.B, ln, t_inc int, f func(dst []float32, incDst, idst uintptr, alpha float32, x, y []float32, n, incX, incY, ix, iy uintptr)) {
	n, inc := uintptr(ln), uintptr(t_inc)
	var idx int
	if t_inc < 0 {
		idx = (-ln + 1) * t_inc
	}
	for i := 0; i < t.N; i++ {
		f(z, inc, uintptr(idx), 1, x, y, n, inc, inc, uintptr(idx), uintptr(idx))
	}
}

func naiveaxpyincto(dst []float32, incDst, idst uintptr, alpha float32, x, y []float32, n, incX, incY, ix, iy uintptr) {
	for i := 0; i < int(n); i++ {
		dst[idst] = alpha*x[ix] + y[iy]
		ix += incX
		iy += incY
		idst += incDst
	}
}

func BenchmarkF32AxpyIncToN1Inc1(b *testing.B) { benchaxpyincto(b, 1, 1, AxpyIncTo) }

func BenchmarkF32AxpyIncToN2Inc1(b *testing.B)  { benchaxpyincto(b, 2, 1, AxpyIncTo) }
func BenchmarkF32AxpyIncToN2Inc2(b *testing.B)  { benchaxpyincto(b, 2, 2, AxpyIncTo) }
func BenchmarkF32AxpyIncToN2Inc4(b *testing.B)  { benchaxpyincto(b, 2, 4, AxpyIncTo) }
func BenchmarkF32AxpyIncToN2Inc10(b *testing.B) { benchaxpyincto(b, 2, 10, AxpyIncTo) }

func BenchmarkF32AxpyIncToN3Inc1(b *testing.B)  { benchaxpyincto(b, 3, 1, AxpyIncTo) }
func BenchmarkF32AxpyIncToN3Inc2(b *testing.B)  { benchaxpyincto(b, 3, 2, AxpyIncTo) }
func BenchmarkF32AxpyIncToN3Inc4(b *testing.B)  { benchaxpyincto(b, 3, 4, AxpyIncTo) }
func BenchmarkF32AxpyIncToN3Inc10(b *testing.B) { benchaxpyincto(b, 3, 10, AxpyIncTo) }

func BenchmarkF32AxpyIncToN4Inc1(b *testing.B)  { benchaxpyincto(b, 4, 1, AxpyIncTo) }
func BenchmarkF32AxpyIncToN4Inc2(b *testing.B)  { benchaxpyincto(b, 4, 2, AxpyIncTo) }
func BenchmarkF32AxpyIncToN4Inc4(b *testing.B)  { benchaxpyincto(b, 4, 4, AxpyIncTo) }
func BenchmarkF32AxpyIncToN4Inc10(b *testing.B) { benchaxpyincto(b, 4, 10, AxpyIncTo) }

func BenchmarkF32AxpyIncToN10Inc1(b *testing.B)  { benchaxpyincto(b, 10, 1, AxpyIncTo) }
func BenchmarkF32AxpyIncToN10Inc2(b *testing.B)  { benchaxpyincto(b, 10, 2, AxpyIncTo) }
func BenchmarkF32AxpyIncToN10Inc4(b *testing.B)  { benchaxpyincto(b, 10, 4, AxpyIncTo) }
func BenchmarkF32AxpyIncToN10Inc10(b *testing.B) { benchaxpyincto(b, 10, 10, AxpyIncTo) }

func BenchmarkF32AxpyIncToN1000Inc1(b *testing.B)  { benchaxpyincto(b, 1000, 1, AxpyIncTo) }
func BenchmarkF32AxpyIncToN1000Inc2(b *testing.B)  { benchaxpyincto(b, 1000, 2, AxpyIncTo) }
func BenchmarkF32AxpyIncToN1000Inc4(b *testing.B)  { benchaxpyincto(b, 1000, 4, AxpyIncTo) }
func BenchmarkF32AxpyIncToN1000Inc10(b *testing.B) { benchaxpyincto(b, 1000, 10, AxpyIncTo) }

func BenchmarkF32AxpyIncToN100000Inc1(b *testing.B)  { benchaxpyincto(b, 100000, 1, AxpyIncTo) }
func BenchmarkF32AxpyIncToN100000Inc2(b *testing.B)  { benchaxpyincto(b, 100000, 2, AxpyIncTo) }
func BenchmarkF32AxpyIncToN100000Inc4(b *testing.B)  { benchaxpyincto(b, 100000, 4, AxpyIncTo) }
func BenchmarkF32AxpyIncToN100000Inc10(b *testing.B) { benchaxpyincto(b, 100000, 10, AxpyIncTo) }

func BenchmarkF32AxpyIncToN100000IncM1(b *testing.B)  { benchaxpyincto(b, 100000, -1, AxpyIncTo) }
func BenchmarkF32AxpyIncToN100000IncM2(b *testing.B)  { benchaxpyincto(b, 100000, -2, AxpyIncTo) }
func BenchmarkF32AxpyIncToN100000IncM4(b *testing.B)  { benchaxpyincto(b, 100000, -4, AxpyIncTo) }
func BenchmarkF32AxpyIncToN100000IncM10(b *testing.B) { benchaxpyincto(b, 100000, -10, AxpyIncTo) }

func BenchmarkLF32AxpyIncToN1Inc1(b *testing.B) { benchaxpyincto(b, 1, 1, naiveaxpyincto) }

func BenchmarkLF32AxpyIncToN2Inc1(b *testing.B)  { benchaxpyincto(b, 2, 1, naiveaxpyincto) }
func BenchmarkLF32AxpyIncToN2Inc2(b *testing.B)  { benchaxpyincto(b, 2, 2, naiveaxpyincto) }
func BenchmarkLF32AxpyIncToN2Inc4(b *testing.B)  { benchaxpyincto(b, 2, 4, naiveaxpyincto) }
func BenchmarkLF32AxpyIncToN2Inc10(b *testing.B) { benchaxpyincto(b, 2, 10, naiveaxpyincto) }

func BenchmarkLF32AxpyIncToN3Inc1(b *testing.B)  { benchaxpyincto(b, 3, 1, naiveaxpyincto) }
func BenchmarkLF32AxpyIncToN3Inc2(b *testing.B)  { benchaxpyincto(b, 3, 2, naiveaxpyincto) }
func BenchmarkLF32AxpyIncToN3Inc4(b *testing.B)  { benchaxpyincto(b, 3, 4, naiveaxpyincto) }
func BenchmarkLF32AxpyIncToN3Inc10(b *testing.B) { benchaxpyincto(b, 3, 10, naiveaxpyincto) }

func BenchmarkLF32AxpyIncToN4Inc1(b *testing.B)  { benchaxpyincto(b, 4, 1, naiveaxpyincto) }
func BenchmarkLF32AxpyIncToN4Inc2(b *testing.B)  { benchaxpyincto(b, 4, 2, naiveaxpyincto) }
func BenchmarkLF32AxpyIncToN4Inc4(b *testing.B)  { benchaxpyincto(b, 4, 4, naiveaxpyincto) }
func BenchmarkLF32AxpyIncToN4Inc10(b *testing.B) { benchaxpyincto(b, 4, 10, naiveaxpyincto) }

func BenchmarkLF32AxpyIncToN10Inc1(b *testing.B)  { benchaxpyincto(b, 10, 1, naiveaxpyincto) }
func BenchmarkLF32AxpyIncToN10Inc2(b *testing.B)  { benchaxpyincto(b, 10, 2, naiveaxpyincto) }
func BenchmarkLF32AxpyIncToN10Inc4(b *testing.B)  { benchaxpyincto(b, 10, 4, naiveaxpyincto) }
func BenchmarkLF32AxpyIncToN10Inc10(b *testing.B) { benchaxpyincto(b, 10, 10, naiveaxpyincto) }

func BenchmarkLF32AxpyIncToN1000Inc1(b *testing.B)  { benchaxpyincto(b, 1000, 1, naiveaxpyincto) }
func BenchmarkLF32AxpyIncToN1000Inc2(b *testing.B)  { benchaxpyincto(b, 1000, 2, naiveaxpyincto) }
func BenchmarkLF32AxpyIncToN1000Inc4(b *testing.B)  { benchaxpyincto(b, 1000, 4, naiveaxpyincto) }
func BenchmarkLF32AxpyIncToN1000Inc10(b *testing.B) { benchaxpyincto(b, 1000, 10, naiveaxpyincto) }

func BenchmarkLF32AxpyIncToN100000Inc1(b *testing.B)  { benchaxpyincto(b, 100000, 1, naiveaxpyincto) }
func BenchmarkLF32AxpyIncToN100000Inc2(b *testing.B)  { benchaxpyincto(b, 100000, 2, naiveaxpyincto) }
func BenchmarkLF32AxpyIncToN100000Inc4(b *testing.B)  { benchaxpyincto(b, 100000, 4, naiveaxpyincto) }
func BenchmarkLF32AxpyIncToN100000Inc10(b *testing.B) { benchaxpyincto(b, 100000, 10, naiveaxpyincto) }

func BenchmarkLF32AxpyIncToN100000IncM1(b *testing.B)  { benchaxpyincto(b, 100000, -1, naiveaxpyincto) }
func BenchmarkLF32AxpyIncToN100000IncM2(b *testing.B)  { benchaxpyincto(b, 100000, -2, naiveaxpyincto) }
func BenchmarkLF32AxpyIncToN100000IncM4(b *testing.B)  { benchaxpyincto(b, 100000, -4, naiveaxpyincto) }
func BenchmarkLF32AxpyIncToN100000IncM10(b *testing.B) { benchaxpyincto(b, 100000, -10, naiveaxpyincto) }

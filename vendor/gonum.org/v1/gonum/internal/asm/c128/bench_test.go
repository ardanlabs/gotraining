// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package c128

import "testing"

var (
	a = complex128(2 + 2i)
	x = make([]complex128, 1000000)
	y = make([]complex128, 1000000)
	z = make([]complex128, 1000000)
)

func init() {
	for n := range x {
		x[n] = complex(float64(n), float64(n))
		y[n] = complex(float64(n), float64(n))
	}
}

func benchaxpyu(t *testing.B, n int, f func(a complex128, x, y []complex128)) {
	x, y := x[:n], y[:n]
	for i := 0; i < t.N; i++ {
		f(a, x, y)
	}
}

func naiveaxpyu(a complex128, x, y []complex128) {
	for i, v := range x {
		y[i] += a * v
	}
}

func BenchmarkC128AxpyUnitary1(t *testing.B)     { benchaxpyu(t, 1, AxpyUnitary) }
func BenchmarkC128AxpyUnitary2(t *testing.B)     { benchaxpyu(t, 2, AxpyUnitary) }
func BenchmarkC128AxpyUnitary3(t *testing.B)     { benchaxpyu(t, 3, AxpyUnitary) }
func BenchmarkC128AxpyUnitary4(t *testing.B)     { benchaxpyu(t, 4, AxpyUnitary) }
func BenchmarkC128AxpyUnitary5(t *testing.B)     { benchaxpyu(t, 5, AxpyUnitary) }
func BenchmarkC128AxpyUnitary10(t *testing.B)    { benchaxpyu(t, 10, AxpyUnitary) }
func BenchmarkC128AxpyUnitary100(t *testing.B)   { benchaxpyu(t, 100, AxpyUnitary) }
func BenchmarkC128AxpyUnitary1000(t *testing.B)  { benchaxpyu(t, 1000, AxpyUnitary) }
func BenchmarkC128AxpyUnitary5000(t *testing.B)  { benchaxpyu(t, 5000, AxpyUnitary) }
func BenchmarkC128AxpyUnitary10000(t *testing.B) { benchaxpyu(t, 10000, AxpyUnitary) }
func BenchmarkC128AxpyUnitary50000(t *testing.B) { benchaxpyu(t, 50000, AxpyUnitary) }

func BenchmarkLC128AxpyUnitary1(t *testing.B)     { benchaxpyu(t, 1, naiveaxpyu) }
func BenchmarkLC128AxpyUnitary2(t *testing.B)     { benchaxpyu(t, 2, naiveaxpyu) }
func BenchmarkLC128AxpyUnitary3(t *testing.B)     { benchaxpyu(t, 3, naiveaxpyu) }
func BenchmarkLC128AxpyUnitary4(t *testing.B)     { benchaxpyu(t, 4, naiveaxpyu) }
func BenchmarkLC128AxpyUnitary5(t *testing.B)     { benchaxpyu(t, 5, naiveaxpyu) }
func BenchmarkLC128AxpyUnitary10(t *testing.B)    { benchaxpyu(t, 10, naiveaxpyu) }
func BenchmarkLC128AxpyUnitary100(t *testing.B)   { benchaxpyu(t, 100, naiveaxpyu) }
func BenchmarkLC128AxpyUnitary1000(t *testing.B)  { benchaxpyu(t, 1000, naiveaxpyu) }
func BenchmarkLC128AxpyUnitary5000(t *testing.B)  { benchaxpyu(t, 5000, naiveaxpyu) }
func BenchmarkLC128AxpyUnitary10000(t *testing.B) { benchaxpyu(t, 10000, naiveaxpyu) }
func BenchmarkLC128AxpyUnitary50000(t *testing.B) { benchaxpyu(t, 50000, naiveaxpyu) }

func benchaxpyut(t *testing.B, n int, f func(d []complex128, a complex128, x, y []complex128)) {
	x, y, z := x[:n], y[:n], z[:n]
	for i := 0; i < t.N; i++ {
		f(z, a, x, y)
	}
}

func naiveaxpyut(d []complex128, a complex128, x, y []complex128) {
	for i, v := range x {
		d[i] = y[i] + a*v
	}
}

func BenchmarkC128AxpyUnitaryTo1(t *testing.B)     { benchaxpyut(t, 1, AxpyUnitaryTo) }
func BenchmarkC128AxpyUnitaryTo2(t *testing.B)     { benchaxpyut(t, 2, AxpyUnitaryTo) }
func BenchmarkC128AxpyUnitaryTo3(t *testing.B)     { benchaxpyut(t, 3, AxpyUnitaryTo) }
func BenchmarkC128AxpyUnitaryTo4(t *testing.B)     { benchaxpyut(t, 4, AxpyUnitaryTo) }
func BenchmarkC128AxpyUnitaryTo5(t *testing.B)     { benchaxpyut(t, 5, AxpyUnitaryTo) }
func BenchmarkC128AxpyUnitaryTo10(t *testing.B)    { benchaxpyut(t, 10, AxpyUnitaryTo) }
func BenchmarkC128AxpyUnitaryTo100(t *testing.B)   { benchaxpyut(t, 100, AxpyUnitaryTo) }
func BenchmarkC128AxpyUnitaryTo1000(t *testing.B)  { benchaxpyut(t, 1000, AxpyUnitaryTo) }
func BenchmarkC128AxpyUnitaryTo5000(t *testing.B)  { benchaxpyut(t, 5000, AxpyUnitaryTo) }
func BenchmarkC128AxpyUnitaryTo10000(t *testing.B) { benchaxpyut(t, 10000, AxpyUnitaryTo) }
func BenchmarkC128AxpyUnitaryTo50000(t *testing.B) { benchaxpyut(t, 50000, AxpyUnitaryTo) }

func BenchmarkLC128AxpyUnitaryTo1(t *testing.B)     { benchaxpyut(t, 1, naiveaxpyut) }
func BenchmarkLC128AxpyUnitaryTo2(t *testing.B)     { benchaxpyut(t, 2, naiveaxpyut) }
func BenchmarkLC128AxpyUnitaryTo3(t *testing.B)     { benchaxpyut(t, 3, naiveaxpyut) }
func BenchmarkLC128AxpyUnitaryTo4(t *testing.B)     { benchaxpyut(t, 4, naiveaxpyut) }
func BenchmarkLC128AxpyUnitaryTo5(t *testing.B)     { benchaxpyut(t, 5, naiveaxpyut) }
func BenchmarkLC128AxpyUnitaryTo10(t *testing.B)    { benchaxpyut(t, 10, naiveaxpyut) }
func BenchmarkLC128AxpyUnitaryTo100(t *testing.B)   { benchaxpyut(t, 100, naiveaxpyut) }
func BenchmarkLC128AxpyUnitaryTo1000(t *testing.B)  { benchaxpyut(t, 1000, naiveaxpyut) }
func BenchmarkLC128AxpyUnitaryTo5000(t *testing.B)  { benchaxpyut(t, 5000, naiveaxpyut) }
func BenchmarkLC128AxpyUnitaryTo10000(t *testing.B) { benchaxpyut(t, 10000, naiveaxpyut) }
func BenchmarkLC128AxpyUnitaryTo50000(t *testing.B) { benchaxpyut(t, 50000, naiveaxpyut) }

func benchaxpyinc(t *testing.B, ln, t_inc int, f func(alpha complex128, x, y []complex128, n, incX, incY, ix, iy uintptr)) {
	n, inc := uintptr(ln), uintptr(t_inc)
	var idx int
	if t_inc < 0 {
		idx = (-ln + 1) * t_inc
	}
	for i := 0; i < t.N; i++ {
		f(1+1i, x, y, n, inc, inc, uintptr(idx), uintptr(idx))
	}
}

func naiveaxpyinc(alpha complex128, x, y []complex128, n, incX, incY, ix, iy uintptr) {
	for i := 0; i < int(n); i++ {
		y[iy] += alpha * x[ix]
		ix += incX
		iy += incY
	}
}

func BenchmarkC128AxpyIncN1Inc1(b *testing.B) { benchaxpyinc(b, 1, 1, AxpyInc) }

func BenchmarkC128AxpyIncN2Inc1(b *testing.B)  { benchaxpyinc(b, 2, 1, AxpyInc) }
func BenchmarkC128AxpyIncN2Inc2(b *testing.B)  { benchaxpyinc(b, 2, 2, AxpyInc) }
func BenchmarkC128AxpyIncN2Inc4(b *testing.B)  { benchaxpyinc(b, 2, 4, AxpyInc) }
func BenchmarkC128AxpyIncN2Inc10(b *testing.B) { benchaxpyinc(b, 2, 10, AxpyInc) }

func BenchmarkC128AxpyIncN3Inc1(b *testing.B)  { benchaxpyinc(b, 3, 1, AxpyInc) }
func BenchmarkC128AxpyIncN3Inc2(b *testing.B)  { benchaxpyinc(b, 3, 2, AxpyInc) }
func BenchmarkC128AxpyIncN3Inc4(b *testing.B)  { benchaxpyinc(b, 3, 4, AxpyInc) }
func BenchmarkC128AxpyIncN3Inc10(b *testing.B) { benchaxpyinc(b, 3, 10, AxpyInc) }

func BenchmarkC128AxpyIncN4Inc1(b *testing.B)  { benchaxpyinc(b, 4, 1, AxpyInc) }
func BenchmarkC128AxpyIncN4Inc2(b *testing.B)  { benchaxpyinc(b, 4, 2, AxpyInc) }
func BenchmarkC128AxpyIncN4Inc4(b *testing.B)  { benchaxpyinc(b, 4, 4, AxpyInc) }
func BenchmarkC128AxpyIncN4Inc10(b *testing.B) { benchaxpyinc(b, 4, 10, AxpyInc) }

func BenchmarkC128AxpyIncN10Inc1(b *testing.B)  { benchaxpyinc(b, 10, 1, AxpyInc) }
func BenchmarkC128AxpyIncN10Inc2(b *testing.B)  { benchaxpyinc(b, 10, 2, AxpyInc) }
func BenchmarkC128AxpyIncN10Inc4(b *testing.B)  { benchaxpyinc(b, 10, 4, AxpyInc) }
func BenchmarkC128AxpyIncN10Inc10(b *testing.B) { benchaxpyinc(b, 10, 10, AxpyInc) }

func BenchmarkC128AxpyIncN1000Inc1(b *testing.B)  { benchaxpyinc(b, 1000, 1, AxpyInc) }
func BenchmarkC128AxpyIncN1000Inc2(b *testing.B)  { benchaxpyinc(b, 1000, 2, AxpyInc) }
func BenchmarkC128AxpyIncN1000Inc4(b *testing.B)  { benchaxpyinc(b, 1000, 4, AxpyInc) }
func BenchmarkC128AxpyIncN1000Inc10(b *testing.B) { benchaxpyinc(b, 1000, 10, AxpyInc) }

func BenchmarkC128AxpyIncN100000Inc1(b *testing.B)  { benchaxpyinc(b, 100000, 1, AxpyInc) }
func BenchmarkC128AxpyIncN100000Inc2(b *testing.B)  { benchaxpyinc(b, 100000, 2, AxpyInc) }
func BenchmarkC128AxpyIncN100000Inc4(b *testing.B)  { benchaxpyinc(b, 100000, 4, AxpyInc) }
func BenchmarkC128AxpyIncN100000Inc10(b *testing.B) { benchaxpyinc(b, 100000, 10, AxpyInc) }

func BenchmarkC128AxpyIncN100000IncM1(b *testing.B)  { benchaxpyinc(b, 100000, -1, AxpyInc) }
func BenchmarkC128AxpyIncN100000IncM2(b *testing.B)  { benchaxpyinc(b, 100000, -2, AxpyInc) }
func BenchmarkC128AxpyIncN100000IncM4(b *testing.B)  { benchaxpyinc(b, 100000, -4, AxpyInc) }
func BenchmarkC128AxpyIncN100000IncM10(b *testing.B) { benchaxpyinc(b, 100000, -10, AxpyInc) }

func BenchmarkLC128AxpyIncN1Inc1(b *testing.B) { benchaxpyinc(b, 1, 1, naiveaxpyinc) }

func BenchmarkLC128AxpyIncN2Inc1(b *testing.B)  { benchaxpyinc(b, 2, 1, naiveaxpyinc) }
func BenchmarkLC128AxpyIncN2Inc2(b *testing.B)  { benchaxpyinc(b, 2, 2, naiveaxpyinc) }
func BenchmarkLC128AxpyIncN2Inc4(b *testing.B)  { benchaxpyinc(b, 2, 4, naiveaxpyinc) }
func BenchmarkLC128AxpyIncN2Inc10(b *testing.B) { benchaxpyinc(b, 2, 10, naiveaxpyinc) }

func BenchmarkLC128AxpyIncN3Inc1(b *testing.B)  { benchaxpyinc(b, 3, 1, naiveaxpyinc) }
func BenchmarkLC128AxpyIncN3Inc2(b *testing.B)  { benchaxpyinc(b, 3, 2, naiveaxpyinc) }
func BenchmarkLC128AxpyIncN3Inc4(b *testing.B)  { benchaxpyinc(b, 3, 4, naiveaxpyinc) }
func BenchmarkLC128AxpyIncN3Inc10(b *testing.B) { benchaxpyinc(b, 3, 10, naiveaxpyinc) }

func BenchmarkLC128AxpyIncN4Inc1(b *testing.B)  { benchaxpyinc(b, 4, 1, naiveaxpyinc) }
func BenchmarkLC128AxpyIncN4Inc2(b *testing.B)  { benchaxpyinc(b, 4, 2, naiveaxpyinc) }
func BenchmarkLC128AxpyIncN4Inc4(b *testing.B)  { benchaxpyinc(b, 4, 4, naiveaxpyinc) }
func BenchmarkLC128AxpyIncN4Inc10(b *testing.B) { benchaxpyinc(b, 4, 10, naiveaxpyinc) }

func BenchmarkLC128AxpyIncN10Inc1(b *testing.B)  { benchaxpyinc(b, 10, 1, naiveaxpyinc) }
func BenchmarkLC128AxpyIncN10Inc2(b *testing.B)  { benchaxpyinc(b, 10, 2, naiveaxpyinc) }
func BenchmarkLC128AxpyIncN10Inc4(b *testing.B)  { benchaxpyinc(b, 10, 4, naiveaxpyinc) }
func BenchmarkLC128AxpyIncN10Inc10(b *testing.B) { benchaxpyinc(b, 10, 10, naiveaxpyinc) }

func BenchmarkLC128AxpyIncN1000Inc1(b *testing.B)  { benchaxpyinc(b, 1000, 1, naiveaxpyinc) }
func BenchmarkLC128AxpyIncN1000Inc2(b *testing.B)  { benchaxpyinc(b, 1000, 2, naiveaxpyinc) }
func BenchmarkLC128AxpyIncN1000Inc4(b *testing.B)  { benchaxpyinc(b, 1000, 4, naiveaxpyinc) }
func BenchmarkLC128AxpyIncN1000Inc10(b *testing.B) { benchaxpyinc(b, 1000, 10, naiveaxpyinc) }

func BenchmarkLC128AxpyIncN100000Inc1(b *testing.B)  { benchaxpyinc(b, 100000, 1, naiveaxpyinc) }
func BenchmarkLC128AxpyIncN100000Inc2(b *testing.B)  { benchaxpyinc(b, 100000, 2, naiveaxpyinc) }
func BenchmarkLC128AxpyIncN100000Inc4(b *testing.B)  { benchaxpyinc(b, 100000, 4, naiveaxpyinc) }
func BenchmarkLC128AxpyIncN100000Inc10(b *testing.B) { benchaxpyinc(b, 100000, 10, naiveaxpyinc) }

func BenchmarkLC128AxpyIncN100000IncM1(b *testing.B)  { benchaxpyinc(b, 100000, -1, naiveaxpyinc) }
func BenchmarkLC128AxpyIncN100000IncM2(b *testing.B)  { benchaxpyinc(b, 100000, -2, naiveaxpyinc) }
func BenchmarkLC128AxpyIncN100000IncM4(b *testing.B)  { benchaxpyinc(b, 100000, -4, naiveaxpyinc) }
func BenchmarkLC128AxpyIncN100000IncM10(b *testing.B) { benchaxpyinc(b, 100000, -10, naiveaxpyinc) }

func benchaxpyincto(t *testing.B, ln, t_inc int, f func(dst []complex128, incDst, idst uintptr, alpha complex128, x, y []complex128, n, incX, incY, ix, iy uintptr)) {
	n, inc := uintptr(ln), uintptr(t_inc)
	var idx int
	if t_inc < 0 {
		idx = (-ln + 1) * t_inc
	}
	for i := 0; i < t.N; i++ {
		f(z, inc, uintptr(idx), 1+1i, x, y, n, inc, inc, uintptr(idx), uintptr(idx))
	}
}

func naiveaxpyincto(dst []complex128, incDst, idst uintptr, alpha complex128, x, y []complex128, n, incX, incY, ix, iy uintptr) {
	for i := 0; i < int(n); i++ {
		dst[idst] = alpha*x[ix] + y[iy]
		ix += incX
		iy += incY
		idst += incDst
	}
}

func BenchmarkC128AxpyIncToN1Inc1(b *testing.B) { benchaxpyincto(b, 1, 1, AxpyIncTo) }

func BenchmarkC128AxpyIncToN2Inc1(b *testing.B)  { benchaxpyincto(b, 2, 1, AxpyIncTo) }
func BenchmarkC128AxpyIncToN2Inc2(b *testing.B)  { benchaxpyincto(b, 2, 2, AxpyIncTo) }
func BenchmarkC128AxpyIncToN2Inc4(b *testing.B)  { benchaxpyincto(b, 2, 4, AxpyIncTo) }
func BenchmarkC128AxpyIncToN2Inc10(b *testing.B) { benchaxpyincto(b, 2, 10, AxpyIncTo) }

func BenchmarkC128AxpyIncToN3Inc1(b *testing.B)  { benchaxpyincto(b, 3, 1, AxpyIncTo) }
func BenchmarkC128AxpyIncToN3Inc2(b *testing.B)  { benchaxpyincto(b, 3, 2, AxpyIncTo) }
func BenchmarkC128AxpyIncToN3Inc4(b *testing.B)  { benchaxpyincto(b, 3, 4, AxpyIncTo) }
func BenchmarkC128AxpyIncToN3Inc10(b *testing.B) { benchaxpyincto(b, 3, 10, AxpyIncTo) }

func BenchmarkC128AxpyIncToN4Inc1(b *testing.B)  { benchaxpyincto(b, 4, 1, AxpyIncTo) }
func BenchmarkC128AxpyIncToN4Inc2(b *testing.B)  { benchaxpyincto(b, 4, 2, AxpyIncTo) }
func BenchmarkC128AxpyIncToN4Inc4(b *testing.B)  { benchaxpyincto(b, 4, 4, AxpyIncTo) }
func BenchmarkC128AxpyIncToN4Inc10(b *testing.B) { benchaxpyincto(b, 4, 10, AxpyIncTo) }

func BenchmarkC128AxpyIncToN10Inc1(b *testing.B)  { benchaxpyincto(b, 10, 1, AxpyIncTo) }
func BenchmarkC128AxpyIncToN10Inc2(b *testing.B)  { benchaxpyincto(b, 10, 2, AxpyIncTo) }
func BenchmarkC128AxpyIncToN10Inc4(b *testing.B)  { benchaxpyincto(b, 10, 4, AxpyIncTo) }
func BenchmarkC128AxpyIncToN10Inc10(b *testing.B) { benchaxpyincto(b, 10, 10, AxpyIncTo) }

func BenchmarkC128AxpyIncToN1000Inc1(b *testing.B)  { benchaxpyincto(b, 1000, 1, AxpyIncTo) }
func BenchmarkC128AxpyIncToN1000Inc2(b *testing.B)  { benchaxpyincto(b, 1000, 2, AxpyIncTo) }
func BenchmarkC128AxpyIncToN1000Inc4(b *testing.B)  { benchaxpyincto(b, 1000, 4, AxpyIncTo) }
func BenchmarkC128AxpyIncToN1000Inc10(b *testing.B) { benchaxpyincto(b, 1000, 10, AxpyIncTo) }

func BenchmarkC128AxpyIncToN100000Inc1(b *testing.B)  { benchaxpyincto(b, 100000, 1, AxpyIncTo) }
func BenchmarkC128AxpyIncToN100000Inc2(b *testing.B)  { benchaxpyincto(b, 100000, 2, AxpyIncTo) }
func BenchmarkC128AxpyIncToN100000Inc4(b *testing.B)  { benchaxpyincto(b, 100000, 4, AxpyIncTo) }
func BenchmarkC128AxpyIncToN100000Inc10(b *testing.B) { benchaxpyincto(b, 100000, 10, AxpyIncTo) }

func BenchmarkC128AxpyIncToN100000IncM1(b *testing.B)  { benchaxpyincto(b, 100000, -1, AxpyIncTo) }
func BenchmarkC128AxpyIncToN100000IncM2(b *testing.B)  { benchaxpyincto(b, 100000, -2, AxpyIncTo) }
func BenchmarkC128AxpyIncToN100000IncM4(b *testing.B)  { benchaxpyincto(b, 100000, -4, AxpyIncTo) }
func BenchmarkC128AxpyIncToN100000IncM10(b *testing.B) { benchaxpyincto(b, 100000, -10, AxpyIncTo) }

func BenchmarkLC128AxpyIncToN1Inc1(b *testing.B) { benchaxpyincto(b, 1, 1, naiveaxpyincto) }

func BenchmarkLC128AxpyIncToN2Inc1(b *testing.B)  { benchaxpyincto(b, 2, 1, naiveaxpyincto) }
func BenchmarkLC128AxpyIncToN2Inc2(b *testing.B)  { benchaxpyincto(b, 2, 2, naiveaxpyincto) }
func BenchmarkLC128AxpyIncToN2Inc4(b *testing.B)  { benchaxpyincto(b, 2, 4, naiveaxpyincto) }
func BenchmarkLC128AxpyIncToN2Inc10(b *testing.B) { benchaxpyincto(b, 2, 10, naiveaxpyincto) }

func BenchmarkLC128AxpyIncToN3Inc1(b *testing.B)  { benchaxpyincto(b, 3, 1, naiveaxpyincto) }
func BenchmarkLC128AxpyIncToN3Inc2(b *testing.B)  { benchaxpyincto(b, 3, 2, naiveaxpyincto) }
func BenchmarkLC128AxpyIncToN3Inc4(b *testing.B)  { benchaxpyincto(b, 3, 4, naiveaxpyincto) }
func BenchmarkLC128AxpyIncToN3Inc10(b *testing.B) { benchaxpyincto(b, 3, 10, naiveaxpyincto) }

func BenchmarkLC128AxpyIncToN4Inc1(b *testing.B)  { benchaxpyincto(b, 4, 1, naiveaxpyincto) }
func BenchmarkLC128AxpyIncToN4Inc2(b *testing.B)  { benchaxpyincto(b, 4, 2, naiveaxpyincto) }
func BenchmarkLC128AxpyIncToN4Inc4(b *testing.B)  { benchaxpyincto(b, 4, 4, naiveaxpyincto) }
func BenchmarkLC128AxpyIncToN4Inc10(b *testing.B) { benchaxpyincto(b, 4, 10, naiveaxpyincto) }

func BenchmarkLC128AxpyIncToN10Inc1(b *testing.B)  { benchaxpyincto(b, 10, 1, naiveaxpyincto) }
func BenchmarkLC128AxpyIncToN10Inc2(b *testing.B)  { benchaxpyincto(b, 10, 2, naiveaxpyincto) }
func BenchmarkLC128AxpyIncToN10Inc4(b *testing.B)  { benchaxpyincto(b, 10, 4, naiveaxpyincto) }
func BenchmarkLC128AxpyIncToN10Inc10(b *testing.B) { benchaxpyincto(b, 10, 10, naiveaxpyincto) }

func BenchmarkLC128AxpyIncToN1000Inc1(b *testing.B)  { benchaxpyincto(b, 1000, 1, naiveaxpyincto) }
func BenchmarkLC128AxpyIncToN1000Inc2(b *testing.B)  { benchaxpyincto(b, 1000, 2, naiveaxpyincto) }
func BenchmarkLC128AxpyIncToN1000Inc4(b *testing.B)  { benchaxpyincto(b, 1000, 4, naiveaxpyincto) }
func BenchmarkLC128AxpyIncToN1000Inc10(b *testing.B) { benchaxpyincto(b, 1000, 10, naiveaxpyincto) }

func BenchmarkLC128AxpyIncToN100000Inc1(b *testing.B)  { benchaxpyincto(b, 100000, 1, naiveaxpyincto) }
func BenchmarkLC128AxpyIncToN100000Inc2(b *testing.B)  { benchaxpyincto(b, 100000, 2, naiveaxpyincto) }
func BenchmarkLC128AxpyIncToN100000Inc4(b *testing.B)  { benchaxpyincto(b, 100000, 4, naiveaxpyincto) }
func BenchmarkLC128AxpyIncToN100000Inc10(b *testing.B) { benchaxpyincto(b, 100000, 10, naiveaxpyincto) }

func BenchmarkLC128AxpyIncToN100000IncM1(b *testing.B) { benchaxpyincto(b, 100000, -1, naiveaxpyincto) }
func BenchmarkLC128AxpyIncToN100000IncM2(b *testing.B) { benchaxpyincto(b, 100000, -2, naiveaxpyincto) }
func BenchmarkLC128AxpyIncToN100000IncM4(b *testing.B) { benchaxpyincto(b, 100000, -4, naiveaxpyincto) }
func BenchmarkLC128AxpyIncToN100000IncM10(b *testing.B) {
	benchaxpyincto(b, 100000, -10, naiveaxpyincto)
}

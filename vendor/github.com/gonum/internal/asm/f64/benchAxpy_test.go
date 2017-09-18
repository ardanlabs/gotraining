// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.7

package f64

import (
	"fmt"
	"testing"
)

const (
	testLen = 1e5
)

var (
	a = 2.0
	x = make([]float64, testLen)
	y = make([]float64, testLen)
	z = make([]float64, testLen)
)

func init() {
	for n := range x {
		x[n] = float64(n)
		y[n] = float64(n)
	}
}

func BenchmarkAxpyUnitary(t *testing.B) {
	naiveaxpyu := func(a float64, x, y []float64) {
		for i, v := range x {
			y[i] += a * v
		}
	}
	tests := []struct {
		name string
		f    func(a float64, x, y []float64)
	}{
		{"AxpyUnitary", AxpyUnitary},
		{"NaiveAxpyUnitary", naiveaxpyu},
	}
	for _, test := range tests {
		for _, ln := range []uintptr{1, 3, 10, 30, 1e2, 3e2, 1e3, 3e3, 1e4, 3e4, 1e5} {
			t.Run(fmt.Sprintf("%s-%d", test.name, ln), func(b *testing.B) {
				b.SetBytes(int64(64 * ln))
				x, y := x[:ln], y[:ln]
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					test.f(a, x, y)
				}
			})
		}
	}
}

func BenchmarkAxpyUnitaryTo(t *testing.B) {
	naiveaxpyut := func(d []float64, a float64, x, y []float64) {
		for i, v := range x {
			d[i] = y[i] + a*v
		}
	}
	tests := []struct {
		name string
		f    func(z []float64, a float64, x, y []float64)
	}{
		{"AxpyUnitaryTo", AxpyUnitaryTo},
		{"NaiveAxpyUnitaryTo", naiveaxpyut},
	}
	for _, test := range tests {
		for _, ln := range []uintptr{1, 3, 10, 30, 1e2, 3e2, 1e3, 3e3, 1e4, 3e4, 1e5} {
			t.Run(fmt.Sprintf("%s-%d", test.name, ln), func(b *testing.B) {
				b.SetBytes(int64(64 * ln))
				x, y, z := x[:ln], y[:ln], z[:ln]
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					test.f(z, a, x, y)
				}
			})
		}
	}
}

var incsAxpy = []struct {
	len uintptr
	inc []int
}{
	{1, []int{1}},
	{2, []int{1, 2, 4, 10}},
	{3, []int{1, 2, 4, 10}},
	{4, []int{1, 2, 4, 10}},
	{5, []int{1, 2, 4, 10}},
	{10, []int{1, 2, 4, 10}},
	{500, []int{1, 2, 4, 10}},
	{1e3, []int{1, 2, 4, 10}},
	{1e4, []int{1, 2, 4, 10, -1, -2, -4, -10}},
}

func BenchmarkAxpyInc(t *testing.B) {
	naiveaxpyinc := func(alpha float64, x, y []float64, n, incX, incY, ix, iy uintptr) {
		for i := 0; i < int(n); i++ {
			y[iy] += alpha * x[ix]
			ix += incX
			iy += incY
		}
	}
	tests := []struct {
		name string
		f    func(alpha float64, x, y []float64, n, incX, incY, ix, iy uintptr)
	}{
		{"AxpyInc", AxpyInc},
		{"NaiveAxpyInc", naiveaxpyinc},
	}
	for _, test := range tests {
		for _, tt := range incsAxpy {
			for _, inc := range tt.inc {
				t.Run(fmt.Sprintf("%s-%d-inc(%d)", test.name, tt.len, inc), func(b *testing.B) {
					b.SetBytes(int64(64 * tt.len))
					var idx, tstInc uintptr = 0, uintptr(inc)
					if inc < 0 {
						idx = uintptr((-int(tt.len) + 1) * inc)
					}
					for i := 0; i < b.N; i++ {
						test.f(a, x, y, uintptr(tt.len), tstInc, tstInc, idx, idx)
					}
				})
			}
		}
	}
}

func BenchmarkAxpyIncTo(t *testing.B) {
	naiveaxpyincto := func(dst []float64, incDst, idst uintptr, alpha float64, x, y []float64, n, incX, incY, ix, iy uintptr) {
		for i := 0; i < int(n); i++ {
			dst[idst] = alpha*x[ix] + y[iy]
			ix += incX
			iy += incY
			idst += incDst
		}
	}
	tests := []struct {
		name string
		f    func(dst []float64, incDst, idst uintptr, alpha float64, x, y []float64, n, incX, incY, ix, iy uintptr)
	}{
		{"AxpyIncTo", AxpyIncTo},
		{"NaiveAxpyIncTo", naiveaxpyincto},
	}
	for _, test := range tests {
		for _, tt := range incsAxpy {
			for _, inc := range tt.inc {
				t.Run(fmt.Sprintf("%s-%d-inc(%d)", test.name, tt.len, inc), func(b *testing.B) {
					b.SetBytes(int64(64 * tt.len))
					var idx, tstInc uintptr = 0, uintptr(inc)
					if inc < 0 {
						idx = uintptr((-int(tt.len) + 1) * inc)
					}
					for i := 0; i < b.N; i++ {
						test.f(z, tstInc, idx, a, x, y, uintptr(tt.len),
							tstInc, tstInc, idx, idx)
					}
				})
			}
		}
	}
}

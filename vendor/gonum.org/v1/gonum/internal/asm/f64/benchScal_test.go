// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.7

package f64

import (
	"fmt"
	"testing"
)

var uniScal = []int64{1, 3, 10, 30, 1e2, 3e2, 1e3, 3e3, 1e4, 3e4}

func BenchmarkScalUnitary(t *testing.B) {
	tstName := "ScalUnitary"
	for _, ln := range uniScal {
		t.Run(fmt.Sprintf("%s-%d", tstName, ln), func(b *testing.B) {
			b.SetBytes(64 * ln)
			x := x[:ln]
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ScalUnitary(a, x)
			}
		})
	}
}

func BenchmarkScalUnitaryTo(t *testing.B) {
	tstName := "ScalUnitaryTo"
	for _, ln := range uniScal {
		t.Run(fmt.Sprintf("%s-%d", tstName, ln), func(b *testing.B) {
			b.SetBytes(int64(64 * ln))
			x, y := x[:ln], y[:ln]
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ScalUnitaryTo(y, a, x)
			}
		})
	}
}

var incScal = []struct {
	len uintptr
	inc []int
}{
	{1, []int{1}},
	{3, []int{1, 2, 4, 10}},
	{10, []int{1, 2, 4, 10}},
	{30, []int{1, 2, 4, 10}},
	{1e2, []int{1, 2, 4, 10}},
	{3e2, []int{1, 2, 4, 10}},
	{1e3, []int{1, 2, 4, 10}},
	{3e3, []int{1, 2, 4, 10}},
	{1e4, []int{1, 2, 4, 10}},
}

func BenchmarkScalInc(t *testing.B) {
	tstName := "ScalInc"
	for _, tt := range incScal {
		for _, inc := range tt.inc {
			t.Run(fmt.Sprintf("%s-%d-inc(%d)", tstName, tt.len, inc), func(b *testing.B) {
				b.SetBytes(int64(64 * tt.len))
				tstInc := uintptr(inc)
				for i := 0; i < b.N; i++ {
					ScalInc(a, x, uintptr(tt.len), tstInc)
				}
			})
		}
	}
}

func BenchmarkScalIncTo(t *testing.B) {
	tstName := "ScalIncTo"
	for _, tt := range incScal {
		for _, inc := range tt.inc {
			t.Run(fmt.Sprintf("%s-%d-inc(%d)", tstName, tt.len, inc), func(b *testing.B) {
				b.SetBytes(int64(64 * tt.len))
				tstInc := uintptr(inc)
				for i := 0; i < b.N; i++ {
					ScalIncTo(z, tstInc, a, x, uintptr(tt.len), tstInc)
				}
			})
		}
	}
}

// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.7

package f32

import (
	"fmt"
	"testing"
)

var (
	benchSink   float32
	benchSink64 float64
)

func BenchmarkDotUnitary(t *testing.B) {
	const name = "DotUnitary"
	for _, v := range []int64{1, 2, 3, 4, 5, 10, 100, 1e3, 5e3, 1e4, 5e4} {
		t.Run(fmt.Sprintf("%s-%d", name, v), func(b *testing.B) {
			x, y := x[:v], y[:v]
			b.SetBytes(32 * v)
			for i := 0; i < b.N; i++ {
				benchSink = DotUnitary(x, y)
			}
		})
	}
}

func BenchmarkDdotUnitary(t *testing.B) {
	const name = "DdotUnitary"
	for _, v := range []int64{1, 2, 3, 4, 5, 10, 100, 1e3, 5e3, 1e4, 5e4} {
		t.Run(fmt.Sprintf("%s-%d", name, v), func(b *testing.B) {
			x, y := x[:v], y[:v]
			b.SetBytes(32 * v)
			for i := 0; i < b.N; i++ {
				benchSink64 = DdotUnitary(x, y)
			}
		})
	}
}

var incsDot = []struct {
	len int
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
	{1e4, []int{1, 2, 4, 10, -1, -2, -4, -10}},
}

func BenchmarkDotInc(t *testing.B) {
	const name = "DotInc"
	for _, tt := range incsDot {
		for _, inc := range tt.inc {
			t.Run(fmt.Sprintf("%s-%d-inc(%d)", name, tt.len, inc), func(b *testing.B) {
				b.SetBytes(int64(32 * tt.len))
				idx := 0
				if inc < 0 {
					idx = (-tt.len + 1) * inc
				}
				for i := 0; i < b.N; i++ {
					benchSink = DotInc(x, y, uintptr(tt.len), uintptr(inc), uintptr(inc), uintptr(idx), uintptr(idx))
				}
			})
		}
	}
}

func BenchmarkDdotInc(t *testing.B) {
	const name = "DdotInc"
	for _, tt := range incsDot {
		for _, inc := range tt.inc {
			t.Run(fmt.Sprintf("%s-%d-inc(%d)", name, tt.len, inc), func(b *testing.B) {
				b.SetBytes(int64(32 * tt.len))
				idx := 0
				if inc < 0 {
					idx = (-tt.len + 1) * inc
				}
				for i := 0; i < b.N; i++ {
					benchSink64 = DdotInc(x, y, uintptr(tt.len), uintptr(inc), uintptr(inc), uintptr(idx), uintptr(idx))
				}
			})
		}
	}
}

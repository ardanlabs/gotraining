// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.7

package c128

import (
	"fmt"
	"testing"
)

func BenchmarkDotUnitary(t *testing.B) {
	for _, test := range []struct {
		name string
		f    func(x, y []complex128) complex128
	}{
		{"DotcUnitary", DotcUnitary},
		{"DotuUnitary", DotuUnitary},
	} {
		for _, v := range []int64{1, 2, 3, 4, 5, 10, 100, 1e3, 5e3, 1e4, 5e4} {
			t.Run(fmt.Sprintf("%s-%d", test.name, v), func(b *testing.B) {
				x, y := x[:v], y[:v]
				b.SetBytes(256 * v)
				for i := 0; i < b.N; i++ {
					benchSink = test.f(x, y)
				}
			})
		}
	}
}

func BenchmarkDotInc(t *testing.B) {
	for _, test := range []struct {
		name string
		f    func(x, y []complex128, n, incX, incY, ix, iy uintptr) complex128
	}{
		{"DotcInc", DotcInc},
		{"DotuInc", DotuInc},
	} {
		for _, ln := range []int{1, 2, 3, 4, 5, 10, 100, 1e3, 5e3, 1e4, 5e4} {
			for _, inc := range []int{1, 2, 4, 10, -1, -2, -4, -10} {
				t.Run(fmt.Sprintf("%s-%d-inc%d", test.name, ln, inc), func(b *testing.B) {
					b.SetBytes(int64(256 * ln))
					var idx int
					if inc < 0 {
						idx = (-ln + 1) * inc
					}
					for i := 0; i < b.N; i++ {
						benchSink = test.f(x, y, uintptr(ln),
							uintptr(inc), uintptr(inc),
							uintptr(idx), uintptr(idx))
					}
				})
			}
		}
	}
}

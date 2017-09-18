// Copyright ©2014 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat64

import (
	"math"
	"math/rand"
	"reflect"
	"testing"
)

func TestPool(t *testing.T) {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			m := NewDense(i, j, nil)
			for k := 0; k < 5; k++ {
				work := make([]*Dense, rand.Intn(10)+1)
				for l := range work {
					w := getWorkspace(i, j, true)
					if !reflect.DeepEqual(w.mat, m.mat) {
						t.Error("unexpected non-zeroed matrix returned by getWorkspace")
					}
					if w.capRows != m.capRows {
						t.Error("unexpected capacity matrix returned by getWorkspace")
					}
					if w.capCols != m.capCols {
						t.Error("unexpected capacity matrix returned by getWorkspace")
					}
					if cap(w.mat.Data) >= 2*len(w.mat.Data) {
						t.Errorf("r: %d c: %d -> len: %d cap: %d", i, j, len(w.mat.Data), cap(w.mat.Data))
					}
					w.Set(0, 0, math.NaN())
					work[l] = w
				}
				for _, w := range work {
					putWorkspace(w)
				}
			}
		}
	}
}

var benchmat *Dense

func poolBenchmark(n, r, c int, clear bool) {
	for i := 0; i < n; i++ {
		benchmat = getWorkspace(r, c, clear)
		putWorkspace(benchmat)
	}
}

func newBenchmark(n, r, c int) {
	for i := 0; i < n; i++ {
		benchmat = NewDense(r, c, nil)
	}
}

func BenchmarkPool10by10Uncleared(b *testing.B)   { poolBenchmark(b.N, 10, 10, false) }
func BenchmarkPool10by10Cleared(b *testing.B)     { poolBenchmark(b.N, 10, 10, true) }
func BenchmarkNew10by10(b *testing.B)             { newBenchmark(b.N, 10, 10) }
func BenchmarkPool100by100Uncleared(b *testing.B) { poolBenchmark(b.N, 100, 100, false) }
func BenchmarkPool100by100Cleared(b *testing.B)   { poolBenchmark(b.N, 100, 100, true) }
func BenchmarkNew100by100(b *testing.B)           { newBenchmark(b.N, 100, 100) }

func BenchmarkMulWorkspaceDense100Half(b *testing.B)        { denseMulWorkspaceBench(b, 100, 0.5) }
func BenchmarkMulWorkspaceDense100Tenth(b *testing.B)       { denseMulWorkspaceBench(b, 100, 0.1) }
func BenchmarkMulWorkspaceDense1000Half(b *testing.B)       { denseMulWorkspaceBench(b, 1000, 0.5) }
func BenchmarkMulWorkspaceDense1000Tenth(b *testing.B)      { denseMulWorkspaceBench(b, 1000, 0.1) }
func BenchmarkMulWorkspaceDense1000Hundredth(b *testing.B)  { denseMulWorkspaceBench(b, 1000, 0.01) }
func BenchmarkMulWorkspaceDense1000Thousandth(b *testing.B) { denseMulWorkspaceBench(b, 1000, 0.001) }
func denseMulWorkspaceBench(b *testing.B, size int, rho float64) {
	b.StopTimer()
	a, _ := randDense(size, rho, rand.NormFloat64)
	d, _ := randDense(size, rho, rand.NormFloat64)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		a.Mul(a, d)
	}
}

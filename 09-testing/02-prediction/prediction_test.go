// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/M7KmittyJH

// go test -run=XXX -bench=. -benchtime=20s

// Tests to show branch prediction benchmarks.
package prediction

import "testing"

func BenchmarkIfNotMostlyTrue(b *testing.B) {
	benchmark(b.N, ifNotMostlyTrue)
}

func BenchmarkIfMostlyTrue(b *testing.B) {
	benchmark(b.N, ifMostlyTrue)
}

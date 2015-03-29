// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/M7KmittyJH

// go test -run=XXX -bench=. -benchtime=20s

// Tests to show branch prediction benchmarks.
package prediction

import "testing"

func BenchmarkIfOnly(b *testing.B) {
	benchmark(b.N, ifOnly)
}

func BenchmarkIfElse(b *testing.B) {
	benchmark(b.N, ifElse)
}

func BenchmarkIfOnlyReversed(b *testing.B) {
	benchmark(b.N, ifOnlyReversed)
}

func BenchmarkIfElseReversed(b *testing.B) {
	benchmark(b.N, ifElseReversed)
}

// go test -run=XXX -bench=. -benchtime=20s
//
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

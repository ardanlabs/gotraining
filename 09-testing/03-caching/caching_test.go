// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/opI__KHj9a

// go test -run=XXX -bench=.

// Tests to show how why CPU caches matter.
package caching

import "testing"

// BenchmarkRowTraverse capture the time it takes to perform
// a row traversal.
func BenchmarkRowTraverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rowTraverse()
	}
}

// BenchmarkColTraverse capture the time it takes to perform
// a column traversal.
func BenchmarkColTraverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		colTraverse()
	}
}

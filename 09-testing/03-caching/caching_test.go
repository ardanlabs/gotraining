// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/opI__KHj9a

// go test -run=XXX -bench=.

// Tests to show how why CPU caches matter.
package caching

import "testing"

func BenchmarkRowTraverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rowTraverse()
	}
}

func BenchmarkColTraverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		colTraverse()
	}
}

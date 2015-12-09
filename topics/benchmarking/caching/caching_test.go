// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/F9KjBB84o4

// go test -run none -bench . -benchtime 3s -benchmem

// Tests to show how why CPU caches matter.
package caching

import "testing"

var fa int

// BenchmarkRowTraverse capture the time it takes to perform
// a row traversal.
func BenchmarkRowTraverse(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++ {
		a = rowTraverse()
	}

	fa = a
}

// BenchmarkColTraverse capture the time it takes to perform
// a column traversal.
func BenchmarkColTraverse(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++ {
		a = colTraverse()
	}

	fa = a
}

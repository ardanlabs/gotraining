// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/VVcx4Jg5E6

// go test -run none -bench . -benchtime 3s -benchmem

// Basic benchmark test.
package basic

import "testing"

// fib finds the nth fibonacci number.
func fib(n int) int {
	a := 0
	b := 1

	for i := 0; i < n; i++ {
		temp := a
		a = b
		b = temp + a
	}
	return a
}

var fa int

// BenchmarkFib provides performance numbers for the fibonacci function.
func BenchmarkFib(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++ {
		a = fib(1e5)
	}

	fa = a
}

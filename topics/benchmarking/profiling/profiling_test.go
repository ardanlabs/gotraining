// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Benchmark tests for program.
package main

import "testing"

var fv interface{}

func BenchmarkGetValue(b *testing.B) {
	var v interface{}

	for i := 0; i < b.N; i++ {
		variable := "#string:variable_name"
		v = getValue(variable, vars)
	}

	fv = v
}

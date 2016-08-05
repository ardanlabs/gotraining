// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go test -run none -bench . -benchtime 3s -benchmem

// Tests to see how each algorithm compare.
package main

import (
	"bytes"
	"testing"
)

// Capture the time it takes to execute algorithm one.
func BenchmarkAlgorithmOne(b *testing.B) {
	var output bytes.Buffer

	for i := 0; i < b.N; i++ {
		for _, d := range data {
			output.Reset()
			algorithmOne(d.input, &output)
		}
	}
}

// Capture the time it takes to execute algorithm two.
func BenchmarkAlgorithmTwo(b *testing.B) {
	var output bytes.Buffer

	for i := 0; i < b.N; i++ {
		for _, d := range data {
			output.Reset()
			algorithmTwo(d.input, &output)
		}
	}
}

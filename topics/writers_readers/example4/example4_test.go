// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go test -run none -bench . -benchtime 3s -benchmem

// Tests to see how each algorithm compare.
package main

import (
	"bytes"
	"testing"
)

// assembleInputStream appends all the input slices together to allow for
// consistent testing across all benchmarks.
func assembleInputStream() []byte {
	var out []byte
	for _, d := range data {
		out = append(out, d.input...)
	}
	return out
}

// Capture the time it takes to execute algorithm one.
func BenchmarkAlgorithmOne(b *testing.B) {
	var output bytes.Buffer
	in := bytes.NewReader(assembleInputStream())

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		output.Reset()
		in.Seek(0, 0)
		algOne(in, &output)
	}
}

// Capture the time it takes to execute algorithm two.
func BenchmarkAlgorithmTwo(b *testing.B) {
	var output bytes.Buffer
	in := bytes.NewReader(assembleInputStream())

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		output.Reset()
		in.Seek(0, 0)
		algTwo(in, &output)
	}
}

// Capture the time it takes to execute algorithm three.
func BenchmarkAlgorithmThree(b *testing.B) {
	var output bytes.Buffer
	in := bytes.NewReader(assembleInputStream())

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		output.Reset()
		in.Seek(0, 0)
		algThree(in, &output)
	}
}

// Capture the time it takes to execute algorithm four.
func BenchmarkAlgorithmFour(b *testing.B) {
	var output bytes.Buffer
	in := bytes.NewReader(assembleInputStream())

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		output.Reset()
		in.Seek(0, 0)
		algFour(in, &output)
	}
}

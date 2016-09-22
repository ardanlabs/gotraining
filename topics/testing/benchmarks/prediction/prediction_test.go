// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go test -run none -bench . -benchtime 3s -benchmem

// Package prediction provides code to show how branch
// prediction can affect performance.
package prediction

import (
	"math/rand"
	"testing"
)

// crunch is used to perform branch instructions.
func crunch(data []uint8) uint8 {
	var sum uint8
	for _, v := range data {
		if v < 128 {
			sum--
		} else {
			sum++
		}
	}
	return sum
}

var fa uint8

// BenchmarkPredictable runs the test when the branch is predictable.
func BenchmarkPredictable(b *testing.B) {
	data := make([]uint8, 1024)
	b.ResetTimer()

	var a uint8

	for i := 0; i < b.N; i++ {
		a = crunch(data)
	}

	fa = a
}

// BenchmarkUnpredictable runs the test when the branch is random.
func BenchmarkUnpredictable(b *testing.B) {
	data := make([]uint8, 1024)
	rand.Seed(0)

	// Fill data with (pseudo) random noise
	for i := range data {
		data[i] = uint8(rand.Uint32())
	}

	b.ResetTimer()

	var a uint8

	for i := 0; i < b.N; i++ {
		a = crunch(data)
	}

	fa = a
}

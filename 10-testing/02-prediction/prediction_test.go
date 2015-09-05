// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/hwZqjJNdbm

// go build -gcflags -m

// Package prediction provides code to show how branch
// prediction can affect performance.
package prediction

import (
	"math/rand"
	"testing"
)

func BenchmarkPredictable(b *testing.B) {
	data := make([]uint8, 1024)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		crunch(data)
	}
}

func BenchmarkUnpredictable(b *testing.B) {
	data := make([]uint8, 1024)
	rand.Seed(0)

	// Fill data with (pseudo) random noise
	for i := range data {
		data[i] = uint8(rand.Uint32())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		crunch(data)
	}
}

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

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// GOGC=off go test -cpu 1 -run none -bench . -benchtime 3s
// GOGC=off go test -run none -bench . -benchtime 3s

// Tests to show the different performance based on concurrency with
// or without parallelism.
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

var docs []string

func init() {
	rand.Seed(time.Now().UnixNano())
	docs = generateList(1e3)
	fmt.Printf("Processing %d documents using %d goroutines\n", len(docs), runtime.NumCPU())
}

func BenchmarkSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		find("test", docs)
	}
}

func BenchmarkConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findConcurrent(runtime.NumCPU(), "test", docs)
	}
}

func BenchmarkSequentialAgain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		find("test", docs)
	}
}

func BenchmarkConcurrentAgain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findConcurrent(runtime.NumCPU(), "test", docs)
	}
}

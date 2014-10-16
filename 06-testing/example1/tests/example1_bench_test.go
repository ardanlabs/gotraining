// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

/*
http://golang.org/cmd/go/#hdr-Description_of_testing_flags
go test -run=XXX -bench=BenchmarkStation -benchmem

Add Later
-benchtime=20s

The benchmark will call the function first with a value of b.N being 1. Then it will
continue to call the function for a minimum of 1 second to complete the test.
*/

// Benchmark tests for example 1.
package main

import (
	"testing"

	"github.com/ArdanStudios/gotraining/06-testing/example1/buoy"
)

// BenchmarkStation provides stats for how this code is performing.
func BenchmarkStation(b *testing.B) {
	// Perform call to find a station.
	for i := 0; i < b.N; i++ {
		buoy.FindStation("42002")
	}
}

/*
	100         78480563 ns/op
	ok      github.com/ArdanStudios/gotraining/08-testing/example1/tests/benchmarks 7.962s

	The function was called 100 times and the test ran for 7.962 seconds
	It takes 78480563 nanoseconds/operation (0.078 seconds/operation) to run FindStation
*/

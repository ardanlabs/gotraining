/*
http://golang.org/cmd/go/#hdr-Description_of_testing_flags
go test -run=XXX -bench=BenchmarkStation

Each benchmark is run for a minimum of 1 second by default. If the
second has not elapsed when the Benchmark function returns, the value
of b.N is increased in the sequence 1, 2, 5, 10, 20, 50, … and the
function run again.
*/

// Benchmark tests for example 1.
package main

import (
	"testing"

	"github.com/ArdanStudios/gotraining/08-testing/example1/buoy"
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

	Test ran for 7.962 seconds
	It takes 78480563 nanoseconds (0.078 seconds) to run FindStation
*/

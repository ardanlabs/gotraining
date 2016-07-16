// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to see what a trace will look like for basic
// channel latencies.
package main

import "testing"

// TestLatency provides a test to profile and trace channel latencies.
func TestLatency(t *testing.T) {
	ch := make(chan int)

	// Lanuch a goroutine that passes integers into the channel.
	go func() {

		// Close the channel when done.
		defer close(ch)

		// Pass 10k integers into the channel.
		for i := 0; i < 10000; i++ {
			ch <- i
		}
	}()

	// Receive each integer that is sent into the channel.
	for range ch {

		// do nothing (we're looking at overhead)
	}
}

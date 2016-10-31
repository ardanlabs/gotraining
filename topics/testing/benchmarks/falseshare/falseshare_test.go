// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests to show the effect of false sharing on concurrent
// memory writes.
package falseshare

import (
	"sync"
	"testing"
)

// cnt represents a data value for our counter.
type cnt struct {
	counter int64
}

// Create an array of 8 counters.
var countersPad [8]cnt

// BenchmarkGlobal tests the performance of 8 goroutines
// incrementing the global counter in parallel.
func BenchmarkGlobal(b *testing.B) {

	// Create and set the WaitGroup.
	var wg sync.WaitGroup

	for i := 0; i < b.N; i++ {

		// Add the 8 goroutines we will create.
		wg.Add(8)

		// Create the 8 goroutines.
		for g := 0; g < 8; g++ {

			// Have each goroutine loop and increment
			// their specific global counter.
			go func(i int) {
				for {

					// Increment the specific global counter.
					countersPad[i].counter++

					// Check if we have incremented it enough.
					if countersPad[i].counter%1e6 == 0 {

						// Report we are done and terminate.
						wg.Done()
						return
					}
				}
			}(g)
		}

		// Wait for all the goroutines to finish.
		wg.Wait()
	}
}

// BenchmarkLocal tests the performance of 8 goroutines
// incrementing their local counter.
func BenchmarkLocal(b *testing.B) {

	// Create and set the WaitGroup.
	var wg sync.WaitGroup

	for i := 0; i < b.N; i++ {

		// Add the 8 goroutines we will create.
		wg.Add(8)

		// Create the 8 goroutines.
		for g := 0; g < 8; g++ {

			// Have each goroutine loop and increment
			// their specific local counter.
			go func(i int) {

				// Init the local counter.
				var counter int64

				for {
					// Increment the local counter.
					counter++

					// Check if we have incremented it enough.
					if counter%1e6 == 0 {

						// Write the final counter to the
						// specific global counter.
						countersPad[i].counter = counter

						// Report we are done and terminate.
						wg.Done()
						return
					}
				}
			}(g)
		}

		// Wait for all the goroutines to finish.
		wg.Wait()
	}
}

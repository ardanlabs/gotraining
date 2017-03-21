// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to use the atomic package to
// provide safe access to numeric types.
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// counter is a variable incremented by all goroutines.
var counter int64

func main() {

	// Number of goroutines to use.
	const grs = 2

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(grs)

	// Create two goroutines.
	for i := 0; i < grs; i++ {
		go func() {
			incCounter()
			wg.Done()
		}()
	}

	// Wait for the goroutines to finish.
	wg.Wait()

	// Display the final value.
	fmt.Println("Final Counter:", counter)
}

// incCounter increments the package level counter variable.
func incCounter() {
	for count := 0; count < 2; count++ {

		// Safely Add One To Counter.
		atomic.AddInt64(&counter, 1)

		// Yield the thread and be placed back in queue.
		runtime.Gosched()
	}
}

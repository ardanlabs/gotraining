// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/rWL6rs-X3M

// go build -race

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

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

// main is the entry point for the application.
func main() {
	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// Create two goroutines.

	go func() {
		incCounter()
		wg.Done()
	}()

	go func() {
		incCounter()
		wg.Done()
	}()

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

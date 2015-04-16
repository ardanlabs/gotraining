// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// https://play.golang.org/p/LJETaLkVl0

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

// main is the entry point for all Go programs.
func main() {
	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// Create two goroutines.
	go incCounter(1)
	go incCounter(2)

	// Wait for the goroutines to finish.
	wg.Wait()

	// Display the final value.
	fmt.Println("Final Counter:", counter)
}

// incCounter increments the package level counter variable.
func incCounter(id int) {
	for count := 0; count < 2; count++ {
		// Safely Add One To Counter.
		atomic.AddInt64(&counter, 1)

		// Yield the thread and be placed back in queue.
		runtime.Gosched()
	}

	// Tell main we are done.
	wg.Done()
}

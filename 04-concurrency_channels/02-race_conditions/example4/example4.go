// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/4ZPCUnWyV5

// go build -race

// Sample program to show how to use a mutex to define critical
// sections of code that need synchronous access.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter is a variable incremented by all goroutines.
	counter int

	// wg is used to wait for the program to finish.
	wg sync.WaitGroup

	// mutex is used to define a critical section of code.
	mutex sync.Mutex
)

// main is the entry point for all Go programs.
func main() {
	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// Create two goroutines.
	go incCounter(1)
	go incCounter(2)

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}

// incCounter increments the package level Counter variable
// using the Mutex to synchronize and provide safe access.
func incCounter(id int) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// Only allow one goroutine through this
		// critical section at a time.
		mutex.Lock()
		{
			// Capture the value of counter.
			value := counter

			// Yield the thread and be placed back in queue.
			runtime.Gosched()

			// Increment our local value of counter.
			value++

			// Store the value back into counter.
			counter = value
		}
		// Release the lock and allow any
		// waiting goroutine through.
		mutex.Unlock()
	}
}

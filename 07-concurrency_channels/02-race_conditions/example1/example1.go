// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// https://play.golang.org/p/dMHhzuM-TM

// go build -race

// Sample program to show how to create race conditions in
// our programs. We don't want to do this.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// counter is a variable incremented by all goroutines.
var counter int

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
	fmt.Println("Final Counter:", counter)
}

// incCounter increments the package level counter variable.
func incCounter(id int) {
	for count := 0; count < 2; count++ {
		// Capture the value of Counter.
		value := counter

		// Yield the thread and be placed back in queue.
		// DO NOT USE IN PRODUCTION CODE!
		runtime.Gosched()

		// Increment our local value of Counter.
		value++

		// Store the value back into Counter.
		counter = value
	}

	// Tell main we are done.
	wg.Done()
}

/*
==================
WARNING: DATA RACE
Write by goroutine 5:
  main.incCounter()
      /Users/bill/Spaces/Go/Projects/src/github.com/ArdanStudios/gotraining/06-concurrency_channels/02-race_conditions/example1/example1.go:54 +0x76

Previous read by goroutine 6:
  main.incCounter()
      /Users/bill/Spaces/Go/Projects/src/github.com/ArdanStudios/gotraining/06-concurrency_channels/02-race_conditions/example1/example1.go:44 +0x46

Goroutine 5 (running) created at:
  main.main()
      /Users/bill/Spaces/Go/Projects/src/github.com/ArdanStudios/gotraining/06-concurrency_channels/02-race_conditions/example1/example1.go:32 +0x61

Goroutine 6 (running) created at:
  main.main()
      /Users/bill/Spaces/Go/Projects/src/github.com/ArdanStudios/gotraining/06-concurrency_channels/02-race_conditions/example1/example1.go:33 +0x78
==================
Final Counter: 2
Found 1 data race(s)
*/

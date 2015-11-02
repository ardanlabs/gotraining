// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/IwFKbnb1JO

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
	go incCounter()
	go incCounter()

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

// incCounter increments the package level counter variable.
func incCounter() {
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
Write by goroutine 6:
  main.incCounter()
      /Users/bill/.../example1/example1.go:52 +0x6f

Previous read by goroutine 7:
  main.incCounter()
      /Users/bill/.../example1/example1.go:42 +0x41

Goroutine 6 (running) created at:
  main.main()
      /Users/bill/.../example1/example1.go:30 +0x69

Goroutine 7 (running) created at:
  main.main()
      /Users/bill/.../example1/example1.go:31 +0x8a
==================
Final Counter: 2
Found 1 data race(s)
*/

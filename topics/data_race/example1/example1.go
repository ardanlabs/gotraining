// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/wj47Qz9MRa

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
}

/*
==================
WARNING: DATA RACE
Read by goroutine 7:
  main.incCounter()
      /Users/bill/.../example1/example1.go:50 +0x41
  main.main.func2()
      /Users/bill/.../example1/example1.go:37 +0x25

Previous write by goroutine 6:
  main.incCounter()
      /Users/bill/.../example1/example1.go:60 +0x6f
  main.main.func1()
      /Users/bill/.../example1/example1.go:32 +0x25

Goroutine 7 (running) created at:
  main.main()
      /Users/bill/.../example1/example1.go:39 +0x78

Goroutine 6 (finished) created at:
  main.main()
      /Users/bill/.../example1/example1.go:34 +0x60
==================
Final Counter: 4
Found 1 data race(s)
*/

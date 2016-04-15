// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

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
      /Users/bill/.../data_race/example1/example1.go:46 +0x41
  main.main.func1()
      /Users/bill/.../data_race/example1/example1.go:31 +0x25

Previous write by goroutine 6:
  main.incCounter()
      /Users/bill/.../data_race/example1/example1.go:56 +0x6f
  main.main.func1()
      /Users/bill/.../data_race/example1/example1.go:31 +0x25

Goroutine 7 (running) created at:
  main.main()
      /Users/bill/.../data_race/example1/example1.go:33 +0xa1

Goroutine 6 (finished) created at:
  main.main()
      /Users/bill/.../data_race/example1/example1.go:33 +0xa1
==================
Final Counter: 4
Found 1 data race(s)
*/

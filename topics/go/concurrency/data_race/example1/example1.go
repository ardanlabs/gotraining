// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build -race

// Sample program to show how to create race conditions in
// our programs. We don't want to do this.
package main

import (
	"fmt"
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
			for count := 0; count < 2; count++ {

				// Capture the value of Counter.
				value := counter

				// Increment our local value of Counter.
				value++

				// Store the value back into Counter.
				counter = value
			}

			wg.Done()
		}()
	}

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

/*
==================
WARNING: DATA RACE
Read at 0x0000011a5118 by goroutine 7:
  main.main.func1()
      /Users/bill/code/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example1/example1.go:33 +0x4e

Previous write at 0x0000011a5118 by goroutine 6:
  main.main.func1()
      /Users/bill/code/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example1/example1.go:39 +0x6d

Goroutine 7 (running) created at:
  main.main()
      /Users/bill/code/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example1/example1.go:43 +0xc3

Goroutine 6 (finished) created at:
  main.main()
      /Users/bill/code/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example1/example1.go:43 +0xc3
==================
Final Counter: 4
Found 1 data race(s)
*/

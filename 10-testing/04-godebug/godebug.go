// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/M6_9Ir79EB

// export GODEBUG=schedtrace=1000
// export GODEBUG=schedtrace=1000,scheddetail=1

// Sample program to review scheduler stats.
package main

import (
	"sync"
	"time"
)

// Create a waitgroup.
var wg sync.WaitGroup

// main is the entry point for the application.
func main() {
	// We are going to create 10 goroutines.
	wg.Add(10)

	// Create those 10 goroutines.
	for i := 0; i < 10; i++ {
		go goroutine()
	}

	// Wait for all the goroutines to complete.
	wg.Wait()

	// Wait to see the global runqueue deplete.
	time.Sleep(3 * time.Second)
}

// goroutine does some CPU bound work.
func goroutine() {
	time.Sleep(time.Second)

	var count int
	for i := 0; i < 1e10; i++ {
		count++
	}

	wg.Done()
}

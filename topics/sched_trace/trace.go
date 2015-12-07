// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/yUz3VkxLdo

// Sample program to review scheduler stats.
package main

import (
	"sync"
	"time"
)

// main is the entry point for the application.
func main() {
	// We are going to create 10 goroutines.
	var wg sync.WaitGroup
	wg.Add(10)

	// Create those 10 goroutines.
	for i := 0; i < 10; i++ {
		go func() {
			goroutine()
			wg.Done()
		}()
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
}

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that creates a fixed set of workers to generate random
// numbers. Discard any number divisible by 2. Continue receiving until 100
// numbers are received. Tell the workers to shut down before terminating.
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func main() {

	// Create the channel for sharing results.
	values := make(chan int)

	// Create a channel "shutdown" to tell goroutines when to terminate.
	shutdown := make(chan struct{})

	// Define the size of the worker pool. Use runtime.GOMAXPROCS(0) to size the pool based on number of processors.
	poolSize := runtime.GOMAXPROCS(0)

	// Create a sync.WaitGroup to monitor the Goroutine pool. Add the count.
	var wg sync.WaitGroup
	wg.Add(poolSize)

	// Create a fixed size pool of goroutines to generate random numbers.
	for i := 0; i < poolSize; i++ {
		go func(id int) {

			// Start an infinite loop.
			for {

				// Generate a random number up to 1000.
				n := rand.Intn(1000)

				// Use a select to either send the number or receive the shutdown signal.
				select {

				// In one case send the random number.
				case values <- n:
					fmt.Printf("Worker %d sent %d\n", id, n)

				// In another case receive from the shutdown channel.
				case <-shutdown:
					fmt.Printf("Worker %d shutting down\n", id)
					wg.Done()
					return
				}
			}
		}(i)
	}

	// Create a slice to hold the random numbers.
	var nums []int
	for i := range values {

		// continue the loop if the value was even.
		if i%2 == 0 {
			fmt.Println("Discarding", i)
			continue
		}

		// Store the odd number.
		fmt.Println("Keeping", i)
		nums = append(nums, i)

		// break the loop once we have 100 results.
		if len(nums) == 100 {
			break
		}
	}

	// Send the shutdown signal by closing the channel.
	fmt.Println("Receiver sending shutdown signal")
	close(shutdown)

	// Wait for the Goroutines to finish.
	wg.Wait()

	// Print the values in our slice.
	fmt.Printf("Result count: %d\n", len(nums))
	fmt.Println(nums)
}

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that uses goroutines to generate up to 100 random numbers.
// Do not send values that are divisible by 2. Have the main goroutine receive
// values and add them to a slice.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	goroutines = 100
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// Create the channel for sharing results.
	values := make(chan int)

	// Create a sync.WaitGroup to monitor the Goroutine pool. Add the count.
	var wg sync.WaitGroup
	wg.Add(goroutines)

	// Iterate and launch each goroutine.
	for gr := 0; gr < goroutines; gr++ {

		// Create an anonymous function for each goroutine.
		go func() {

			// Ensure the waitgroup is decremented when this function returns.
			defer wg.Done()

			// Generate a random number up to 1000.
			n := rand.Intn(1000)

			// Return early if the number is divisible by 2. n%2 == 0
			if n%2 == 0 {
				return
			}

			// Send the odd values through the channel.
			values <- n
		}()
	}

	// Create a goroutine that waits for the other goroutines to finish then
	// closes the channel.
	go func() {
		wg.Wait()
		close(values)
	}()

	// Receive from the channel until it is closed.
	// Store values in a slice of ints.
	var nums []int
	for n := range values {
		nums = append(nums, n)
	}

	// Print the values in our slice.
	fmt.Printf("Result count: %d\n", len(nums))
	fmt.Println(nums)
}

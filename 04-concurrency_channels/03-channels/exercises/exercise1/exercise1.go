// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/3ry3sCIfaC

// Write a program where two goroutines pass an integer back and forth
// ten times. Display when each goroutine receives the integer. Increment
// the integer with each pass. Once the interger equals ten, terminate
// the program cleanly.
package main

import (
	"fmt"
	"sync"
)

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

// main is the entry point for all Go programs.
func main() {
	// Create an unbuffered channel.
	share := make(chan int)

	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// Launch two goroutines.
	go goroutine("Bill", share)
	go goroutine("Lisa", share)

	// Start the sharing.
	share <- 1

	// Wait for the program to finish.
	wg.Wait()
}

// goroutine simulates sharing a value.
func goroutine(name string, share chan int) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()
	defer fmt.Printf("Goroutine %s Down\n", name)

	for {
		// Wait for the ball to be hit back to us.
		value, ok := <-share
		if !ok {
			// If the channel was closed, shutdown.
			return
		}

		// Display the value.
		fmt.Printf("Goroutine %s Inc %d\n", name, value)

		// Terminate when the value is 10.
		if value == 10 {
			close(share)
			return
		}

		// Share the value.
		share <- (value + 1)
	}
}

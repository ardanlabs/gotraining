// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that uses select to read and write from multiple channels.
package main

import (
	"fmt"
	"time"
)

func main() {

	// Create a channel for sending int values.
	data := make(chan int)

	// Create a channel "shutdown" to tell the goroutine when to terminate.
	shutdown := make(chan struct{})

	// Create a channel "complete" for the goroutine to tell main when it's done.
	complete := make(chan struct{})

	// Launch a goroutine to generate data through the channel.
	go func() {

		// Create an int variable i.
		var i int

		// Run an infinite loop that uses select to perform channel operations.
		for {
			select {

			// In one case send the value of i.
			case data <- i:
				fmt.Println("Sent", i)             // Print the number sent.
				i++                                // Increment i for the next iteration.
				time.Sleep(100 * time.Millisecond) // Sleep for 100ms to simulate some latency.

			// In another case receive from the shutdown channel.
			case <-shutdown:
				fmt.Println("Sender shutting down") // Print a shutdown message
				close(complete)                     // Close the "complete" channel so main knows we're done.
				return                              // Return from the anonymous function.
			}
		}

	}()

	// Use time.After to make a channel which will send in 1 second.
	tc := time.After(1 * time.Second)

	// Run an infinite loop that uses a label like "loop:"
loop:
	for {
		select {

		// In one case receive the value of i.
		case i := <-data:
			fmt.Println("Received", i) // Print the value received.

		// In another case receive from the timeout channel.
		case <-tc:
			fmt.Println("Initiating shutdown") // Print a message that main is initiating the shutdown sequence.
			close(shutdown)                    // Close the "shutdown" channel so the goroutine knows to terminate.
			break loop                         // Break the main loop.
		}
	}

	// Block waiting to receive from the "complete" channel.
	<-complete

	// Print a message that the program shut down cleanly.
	fmt.Println("Terminating program")
}

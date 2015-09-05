// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/G7O-DnJrEA

// Write a program where two goroutines pass an integer back and forth
// ten times. Display when each goroutine receives the integer. Increment
// the integer with each pass. Once the integer equals ten, terminate
// the program cleanly.
package main

// Add imports.

// Declare a wait group variable.

// Declare a function for the goroutine. Pass in a name for the
// routine and a channel of type int used to share the value back and forth.
func goroutine( /* parameters */ ) {
	for {
		// Receive on the channel, waiting for the integer.

		// Check if the channel is closed.
		{
			// Call done on the waitgroup.
			// Display the goroutine is finished and return.
		}

		// Display the integer value received.

		// Check is the value from the channel is 10.
		{
			// Close the channel.
			// Call done on the waitgroup.
			// Display the goroutine is finished and return.
		}

		// Increment the value by one and send in back through
		// the channel.
	}
}

// main is the entry point for all Go programs.
func main() {
	// Declare and initialize an unbuffered channel
	// of type int.

	// Increment the wait group for each goroutine
	// to be created.

	// Create the two goroutines.

	// Send the initial integer value into the channel.

	// Wait for all the goroutines to finish.
}

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program where two goroutines pass an integer back and forth
// ten times. Display when each goroutine receives the integer. Increment
// the integer with each pass. Once the integer equals ten, terminate
// the program cleanly.
package main

// Add imports.

// Declare a wait group variable.

func main() {

	// Create an unbuffered channel.

	// Set the waitgroup, one for each goroutine.

	// Launch the goroutine and handle Done.

	// Launch the goroutine and handle Done.

	// Send a value to start the counting.

	// Wait for the program to finish.
}

// goroutine simulates sharing a value.
func goroutine( /* parameters */ ) {
	for {

		// Wait for the value to be sent.
		// If the channel was closed, return.

		// Display the value.

		// Terminate when the value is 10.

		// Increment the value and send it
		// over the channel.
	}
}

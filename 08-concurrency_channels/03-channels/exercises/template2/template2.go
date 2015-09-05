// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/vc6c1-M2EB

// Write a problem that uses a buffered channel to maintain a buffer
// of four strings. In main, send the strings 'A', 'B', 'C' and 'D'
// into the channel. Then create 20 goroutines that receive a string
// from the channel, display the value and then send the string back
// into the channel. Once each goroutine is done performing that task,
// allow the goroutine to terminate.
package main

// Add Imports.

// Declare a constant and set the value for the number of goroutines.

// Declare a constant and set the value for the number of buffers.

// Declare a wait group.

// Declare a buffered channel of type string and initialize it based on
// the constant you declared above.

// Declare a function for the goroutine that will process work
// from the buffered channel.
func worker(worker int) {
	// Receive a string from the channel.

	// Display the string.

	// Send the string back into the channel.

	// Tell main this goroutine is done.
}

// main is the entry point for all Go programs.
func main() {
	// Increment the wait group for the number of
	// goroutines based on the value of the constant.

	// Create the number of goroutines based on the
	// value of the constant.

	// Add strings in the buffered channel.

	// Wait for all the work to get done.
}

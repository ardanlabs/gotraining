// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/Ki3X_YIny6

// Write a problem that uses a buffered channel to maintain a buffer
// of four strings. In main, send the strings 'A', 'B', 'C' and 'D'
// into the channel. Then create 20 goroutines that receive a string
// from the channel, display the value and then send the string back
// into the channel. Once each goroutine is done performing that task,
// allow the goroutine to terminate.
package main

// Add Imports.

// Declare constants for number of goroutines and capacity.

// Declare a wait group variable.

// Declare a buffered channel to manage strings
// with a capacity.
var resources = make(chan string, capacity)

// main is the entry point for the application.
func main() {
	// Launch the goroutines to handle the work.
	// Handle the Done when they are finished.

	// Add the strings to the buffered channel.

	// Wait for all the work to get done.
}

// worker is launched as a goroutine to process work from
// the buffered channel.
func worker( /* parameters */ ) {
	// Receive a string from the channel.

	// Display the value.

	// Place the string back.
}

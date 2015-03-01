// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/h0nMS_l1rO

// Write a program where two goroutines pass an integer back and forth
// ten times. Display when each goroutine receives the integer. Increment
// the integer with each pass. Once the integer equals ten, terminate
// the program cleanly.
package main

// Add imports.

// Declare a wait group variable.

// Declare a function for the goroutine. Pass in a name for the
// routine and a channel used to share the value back and forth.
func goroutine( /* parameters */ ) {
	for {
		// Receive on the channel, waiting for the integer. Check
		// if the channel is closed.

		// Display the integer value received.

		// Terminate the goroutine when the value is 10.
		{
			// Close the channel.
			// Wait on the wait group to hit 0.
			// Display the goroutine is finished and return.
		}

		// Increment the value and send in back through
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

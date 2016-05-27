// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that uses a fan out pattern to generate 100 random numbers
// concurrently. Have each goroutine generate a single random number and return
// that number to the main goroutine over a buffered channel. Set the size of
// the buffer channel so no send every blocks. Don't allocate more buffers than
// you need. Have the main goroutine display each random number is receives and
// then terminate the program.
package main

// Add imports.

// Declare constant for number of goroutines .

func init() {
	// Seed the random number generator.
}

func main() {

	// Create the buffer channel with a buffer for
	// each goroutine to be created.

	// Iterate and launch each goroutine.
	{

		// Create an anonymous function for each goroutine that
		// generates a random number and sends it on the channel.
	}

	// Create a variable to be used to track received messages.
	// Set the value to the number of goroutines created.

	// Iterate receiving each value until they are all received.
}

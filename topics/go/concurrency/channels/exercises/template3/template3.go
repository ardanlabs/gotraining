// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that uses goroutines to generate up to 100 random numbers.
// Do not send values that are divisible by 2. Have the main goroutine receive
// values and add them to a slice.
package main

// Declare constant for number of goroutines.
const goroutines = 100

func init() {
	// Seed the random number generator.
}

func main() {

	// Create the channel for sharing results.

	// Create a sync.WaitGroup to monitor the Goroutine pool. Add the count.

	// Iterate and launch each goroutine.
	{

		// Create an anonymous function for each goroutine.
		{

			// Ensure the waitgroup is decremented when this function returns.

			// Generate a random number up to 1000.

			// Return early if the number is even. (n%2 == 0)

			// Send the odd values through the channel.
		}
	}

	// Create a goroutine that waits for the other goroutines to finish then
	// closes the channel.

	// Receive from the channel until it is closed.
	// Store values in a slice of ints.

	// Print the values in our slice.
}

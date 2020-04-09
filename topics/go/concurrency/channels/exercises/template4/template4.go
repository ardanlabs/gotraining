// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that creates a fixed set of workers to generate random
// numbers. Discard any number divisible by 2. Continue receiving until 100
// numbers are received. Tell the workers to shut down before terminating.
package main

// Add imports.

func main() {

	// Create the channel for sharing results.

	// Create a channel "shutdown" to tell goroutines when to terminate.

	// Define the size of the worker pool. Use runtime.GOMAXPROCS(0) to size the pool based on number of processors.

	// Create a sync.WaitGroup to monitor the Goroutine pool. Add the count.

	// Create a fixed size pool of goroutines to generate random numbers.
	{
		{

			// Start an infinite loop.
			{

				// Generate a random number up to 1000.

				// Use a select to either send the number or receive the shutdown signal.
				{

					// In one case send the random number.

					// In another case receive from the shutdown channel.

				}
			}
		}
	}

	// Create a slice to hold the random numbers.

	// Receive from the values channel with range.
	{

		// continue the loop if the value was even.

		// Store the odd number.

		// break the loop once we have 100 results.
	}

	// Send the shutdown signal by closing the shutdown channel.

	// Wait for the Goroutines to finish.

	// Print the values in our slice.
}

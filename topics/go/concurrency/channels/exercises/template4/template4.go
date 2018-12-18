// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that uses select to read and write from multiple channels.
package main

// Add imports.

func main() {

	// Create a channel for sending int values.

	// Create a channel "shutdown" to tell the goroutine when to terminate.

	// Create a channel "complete" for the goroutine to tell main when it's done.

	// Launch a goroutine to generate data through the channel.
	{

		// Create an int variable i.

		// Run an infinite loop that uses select to perform channel operations.
		{
			{

				// In one case send the value of i.

				// Print the number sent.
				// Increment i for the next iteration.
				// Sleep for 100ms to simulate some latency.

				// In another case receive from the shutdown channel.

				// Print a shutdown message
				// Close the "complete" channel so main knows we're done.
				// Return from the anonymous function.
			}
		}

	}

	// Use time.After to make a channel which will send in 1 second.

	// Run an infinite loop that uses a label like "loop:"
	{
		{

			// In one case receive the value of i.

			// Print the value received.

			// In another case receive from the timeout channel.

			// Print a message that main is initiating the shutdown sequence.
			// Close the "shutdown" channel so the goroutine knows to terminate.
			// Break the main loop.
		}
	}

	// Block waiting to receive from the "complete" channel.

	// Print a message that the program shut down cleanly.
}

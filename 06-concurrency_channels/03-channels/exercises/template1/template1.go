// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/oaCmWz6_CR

// Write a program where two goroutines pass an integer back and forth
// ten times. Display when each goroutine receives the integer. Increment
// the integer with each pass. Once the integer equals ten, terminate
// the program cleanly.
package main

import (
	"fmt"
	"sync"
)

// wg is used to wait for the program to finish.
var waitgroup_name sync.waitgroup_type

// main is the entry point for all Go programs.
func main() {
	// Create an unbuffered channel.
	channel_name := make(chan type)

	// Add a count of two, one for each goroutine.
	waitgroup_name.add_method(N)

	// Launch two goroutines.
	keyword function_name("Name", channel_name)
	keywprd function_name("Name", channel_name)

	// Start the sharing.
	channel_name [operator] N

	// Wait for the program to finish.
	waitgroup_name.wait_method()
}

// goroutine simulates sharing a value.
func function_name(parameter_name type, channel_name chan type) {
	// Schedule this code when the function returns.
	defer func() {
		fmt.Printf("Goroutine %s Down\n", parameter_name)
		waitgroup_name.done_method()
	}()

	for {
		// Wait for the ball to be hit back to us.
		variable_name, flag_name := [operator]channel_name
		if !flag_name {
			// If the channel was closed, shutdown.
			return
		}

		// Display the value.
		fmt.Printf("Goroutine %s Inc %d\n", parameter_name, variable_name)

		// Terminate when the value is 10.
		if variable_name == N {
			close(channel_name)
			return
		}

		// Share the value.
		channel_name [operator] (variable_name + N)
	}
}

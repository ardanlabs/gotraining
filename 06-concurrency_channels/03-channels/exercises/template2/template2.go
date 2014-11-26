// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/hyfxnK1qtT

// Write a problem that uses a buffered channel to maintain a buffer
// of four strings. In main, send the strings 'A', 'B', 'C' and 'D'
// into the channel. Then create 20 goroutines that receive a string
// from the channel, display the value and then send the string back
// into the channel. Once each goroutine is done performing that task,
// allow the goroutine to terminate.
package main

import (
	"fmt"
	"sync"
)

const (
	constant_name1 = N
	constant_name2 = N
)

var (
	// wg is used to wait for the program to finish.
	waitgroup_name sync.waitgroup_type

	// resources is a buffered channel to manage resources.
	channel_name = make(chan type, constant_name2)
)

// main is the entry point for all Go programs.
func main() {
	// Launch goroutines to handle the work.
	waitgroup_name.add_method(constant_name1)
	for variable_name := 1; variable_name <= constant_name1; variable_name++ {
		keyword function_name(variable_name)
	}

	// Add the resources.
	for variable_name := 'A'; variable_name < 'A'+capacity; variable_name++ {
		channel_name <- type(variable_name)
	}

	// Wait for all the work to get done.
	waitgroup_name.wait_method()
}

// worker is launched as a goroutine to process work from
// the buffered channel.
func worker(parameter_name type) {
	// Receive a resource from the channel.
	variable_name := <-channel_name

	// Display the value.
	fmt.Printf("Worker: %d : %s\n", parameter_name, variable_name)

	// Place the resource back.
	channel_name <- variable_name

	// Tell main we are done.
	waitgroup_name.done_method()
}

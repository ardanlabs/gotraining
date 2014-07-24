// Program uses a function type, closures, goroutines and channels
// to calculate Fibonacci numbers.
package main

import (
	"fmt"
)

// fib is a function type that takes no parameters
// and returns an integer.
type fib func() int

// ch is an unbuffered channel that passes our
// function type.
var ch = make(chan fib)

// main is the entry point for all Go programs.
func main() {
	calculate(5, fibonacci())
}

// fib returns a function that returns successive
// Fibonacci numbers. This function uses closure
// to maintain the state for the next call.
func fibonacci() fib {
	first, second := 0, 1
	return func() int {
		first, second = second, first+second
		return first
	}
}

// calculate recurses the call to the fibonacci function.
func calculate(toIndex int, f fib) {
	// Create a goroutine to read the channel
	// and call the Fibonacci function.
	go func() {
		for {
			ff := <-ch
			fmt.Println(ff())
		}
	}()

	// Iterate passing the function variable
	// into the channel.
	for i := 0; i < toIndex; i++ {
		ch <- f
	}
}

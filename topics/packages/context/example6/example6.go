// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program that shows how to use the context package
// to avoid leaking goroutines.
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// Create a context that supports cancellation.
	ctx, cancel := context.WithCancel(context.Background())

	// Make sure all paths cancel the context to
	// avoid context leak.
	defer cancel()

	// Range of the gen channel.
	for n := range gen(ctx) {

		// Display the value we received.
		fmt.Println(n)

		// Once we get the number 5 quit.
		if n == 5 {
			cancel()
			break
		}
	}

	fmt.Println("Done generating numbers")

	time.Sleep(time.Millisecond)
	fmt.Println("Program done")
}

// gen is a generator that can be cancellable by cancelling the ctx.
func gen(ctx context.Context) <-chan int {

	// Create the channel to use for the generator.
	ch := make(chan int)

	// Create a goroutine to provide numbers.
	go func() {
		var n int

		for {
			select {

			// Avoid leaking of this goroutine when ctx is done.
			case <-ctx.Done():
				fmt.Println("Generator routine done")
				return

			// Send the current value of n and then increment.
			case ch <- n:
				n++
			}
		}
	}()

	// Return the channel for use.
	return ch
}

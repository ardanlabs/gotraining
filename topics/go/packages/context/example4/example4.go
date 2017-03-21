// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to use the WithTimeout function
// of the Context package.
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// Set a duration.
	duration := 150 * time.Millisecond

	// Create a context that is both manually cancellable and will signal
	// a cancel at the specificed duration.
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	// Create a channel to received a signal that work is done.
	ch := make(chan struct{})

	// Ask the goroutine to do some work for us.
	go func() {
		for {

			// Simulate work.
			time.Sleep(50 * time.Millisecond)

			// Report the work is done.
			ch <- struct{}{}
		}
	}()

	// Wait for the work to finish. If it takes too long move on.
	select {
	case <-ch:
		fmt.Println("work complete")

	case <-ctx.Done():
		fmt.Println("work cancelled")
	}
}

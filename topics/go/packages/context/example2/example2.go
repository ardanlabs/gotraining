// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to use the WithCancel function.
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// Create a context that is cancellable only manually.
	ctx, cancel := context.WithCancel(context.Background())

	// The cancel function must be called regardless of the outcome.
	defer cancel()

	// Ask the goroutine to do some work for us.
	go func() {
		for {

			// Simulate work.
			time.Sleep(50 * time.Millisecond)

			// Report the work is done.
			cancel()
		}
	}()

	// Wait for the work to finish. If it takes too long move on.
	select {
	case <-time.After(100 * time.Millisecond):
		fmt.Println("moving on")

	case <-ctx.Done():
		fmt.Println("work complete")
	}
}

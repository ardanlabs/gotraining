// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to use the WithCancel function
// of the Context package.
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	timeout()
	callCancel()
}

// timeout show how to handle a context that does not timeout.
func timeout() {

	// WithCancel returns a copy of parent with a new Done channel.
	// The returned context's Done channel is closed when the returned
	// cancel function is called or when the parent context's Done
	// channel is closed, whichever happens first.
	ctx, cancel := context.WithCancel(context.Background())

	select {
	case <-time.After(100 * time.Millisecond):
		fmt.Println("overslept")

	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}

	// Even though ctx should have expired already, it is good
	// practice to call its cancelation function in any case.
	// Failure to do so may keep the context and its parent alive
	// longer than necessary.
	cancel()
}

// callCancel show how cancel works with the context.
func callCancel() {

	// WithCancel returns a copy of parent with a new Done channel.
	// The returned context's Done channel is closed when the returned
	// cancel function is called or when the parent context's Done
	// channel is closed, whichever happens first.
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	select {
	case <-time.After(100 * time.Millisecond):
		fmt.Println("overslept")

	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context canceled"
	}

	// Even though we called cancel first, it is good
	// practice to call its cancelation function in any case.
	// Failure to do so may keep the context and its parent alive
	// longer than necessary.
	cancel()
}

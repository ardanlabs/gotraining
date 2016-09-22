// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to use the WithDeadline function
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

// timeout show how timeouts work with the context.
func timeout() {

	// Set a deadline of fifty milliseconds.
	deadline := time.Now().Add(50 * time.Millisecond)

	// WithDeadline returns a copy of the parent context with the deadline adjusted
	// to be no later than time. If the parent's deadline is already earlier than the,
	// time WithDeadline is semantically equivalent to parent. The returned
	// context's Done channel is closed when the deadline expires, when the returned
	// cancel function is called, or when the parent context's Done channel is closed,
	// whichever happens first.
	ctx, cancel := context.WithDeadline(context.Background(), deadline)

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

	// Set a deadline of fifty milliseconds.
	deadline := time.Now().Add(time.Hour)

	// WithDeadline returns a copy of the parent context with the deadline adjusted
	// to be no later than time. If the parent's deadline is already earlier than the,
	// time WithDeadline is semantically equivalent to parent. The returned
	// context's Done channel is closed when the deadline expires, when the returned
	// cancel function is called, or when the parent context's Done channel is closed,
	// whichever happens first.
	ctx, cancel := context.WithDeadline(context.Background(), deadline)

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

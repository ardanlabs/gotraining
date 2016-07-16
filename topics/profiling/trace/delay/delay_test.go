// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to see what a trace will look like when minimized to
// a simple example of channel communication.
package main

import (
	"sync"
	"testing"
	"time"
)

// TestDelay provides a test to profile and trace channel communication.
func TestDelay(t *testing.T) {
	wait := make(chan struct{})
	var mu sync.Mutex

	// Goroutine attempts to acquire a Lock inorder
	// to close the channel.
	go func() {

		// Acquire a lock on the mutex.
		mu.Lock()
		t.Log("Capture mutex lock")

		// Release the lock on the mutex when the
		// goroutine terminates.
		defer func() {
			t.Log("Release mutex lock")
			mu.Unlock()
		}()

		// Close the wait channel.
		t.Log("Closing wait channel")
		close(wait)

		// Wait a second before releasing the mutex lock.
		time.Sleep(1 * time.Second)
	}()

	// Wait for the goroutine to close the channel.
	// This will allow a context switch to occur.
	<-wait
	t.Log("Channel closed")

	// Wait for the goroutine to release the mutex lock.
	mu.Lock()
	defer mu.Unlock()
	t.Log("Test complete")
}

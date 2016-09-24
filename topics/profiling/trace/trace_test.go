// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to see what a memory leak looks like.
package main

import (
	"fmt"
	"testing"
	"time"
)

// TestSTWTrace provides an opportunity to learn how to use the trace
// tool to understand how the GC is affecting the program and the
// health of the heap.
func TestSTWTrace(t *testing.T) {
	ch := make(chan bool)

	// This goroutine allocates memory in the heap by using
	// a map to add keys. In one example we will not remove
	// keys and in the next we will.
	go func() {
		fmt.Println("Goroutine Started, Pounding GC")

		// Make the map.
		m := make(map[int]int)

		// Loop endlessly adding keys to the map.
		for i := 0; ; i++ {

			// Add a key and value.
			m[i] = i

			// See if we have been asked to stop.
			select {
			case <-ch:
				ch <- true
				break

			default:
			}

			// This code removes a million keys once a million
			// keys have been added. We will run the program with
			// and without this code. Without this code we are
			// essentially leaking memory.

			// if i == 1000000 {
			// 	fmt.Println("Deleting")
			// 	for j := 0; j < 1000000; j++ {
			// 		delete(m, j)
			// 	}
			// 	i = -1
			// }
		}
	}()

	// Wait 15 seconds and then tell the goroutine to quit.
	time.Sleep(15 * time.Second)
	ch <- true

	// Wait for it to respond it is done.
	<-ch

	fmt.Println("Goroutine Terminated")
}

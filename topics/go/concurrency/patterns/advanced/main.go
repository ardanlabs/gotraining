// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This sample program demonstrates how the logger package works.
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

// device allows us to mock a device we write logs to.
type device struct {
	mu      sync.RWMutex
	problem bool
}

// Write implements the io.Writer interface.
func (d *device) Write(p []byte) (n int, err error) {

	// Simulate disk problems.
	for d.isProblem() {
		time.Sleep(time.Second)
	}

	fmt.Print(string(p))
	return len(p), nil
}

// isProblem checks in a safe way if there is a problem.
func (d *device) isProblem() bool {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.problem
}

// flipProblem reverses the problem flag to the opposite value.
func (d *device) flipProblem() {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.problem = !d.problem
}

func main() {

	// Number of goroutines that will be writing logs.
	const grs = 10

	// Create a logger value with a buffer of capacity
	// for each goroutine that will be logging.
	var d device
	l := log.New(&d, "prefix", 0)

	// Generate goroutines, each writing to disk.
	for i := 0; i < grs; i++ {
		go func(id int) {
			for {
				l.Println(fmt.Sprintf("%d: log data", id))
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}

	// We want to control the simulated disk blocking. Capture
	// interrupt signals to toggle device issues. Use <ctrl> z
	// to kill the program.

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	for {
		<-sigChan
		d.flipProblem()
	}
}

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This sample program demonstrates how the logger package works.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/ardanlabs/gotraining/topics/concurrency/patterns/logger"
)

func main() {

	// Number of goroutines to simulate.
	const grs = 10

	// Create a logger value with 3 times the capacity.
	l := logger.New(grs * 3)

	// Generate gorutines each writing to disk.
	for i := 0; i < grs; i++ {
		go func(id int) {
			for {
				l.Write(fmt.Sprintf("%d: log data", id))
				time.Sleep(50 * time.Millisecond)
			}
		}(i)
	}

	// We want to control the simulated disk blocking.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	<-sigChan
	fmt.Println("Simulate the Disk is out of space.")
	l.DiskFull()

	<-sigChan
	fmt.Println("Simulate the Disk good again.")
	l.DiskFull()

	<-sigChan
	fmt.Println("Shutting down.")
	l.Shutdown()

	fmt.Println("DOWN")
}

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This sample program demonstrates how the logger package works.
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ardanlabs/gotraining/topics/concurrency_patterns/logger"
)

func main() {

	// Create a logger value.
	l := logger.New(10)

	// Generate 100 gorutines each writing to disk.
	for i := 0; i < 100; i++ {
		go func(id int) {
			for {
				l.Write(fmt.Sprintf("%d: log data", id))
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}

	time.Sleep(1 * time.Second)

	// Simulate large latencies now so data will be thrown away.
	log.Println("Simulate the Disk is out of space.")
	l.DiskFull()

	time.Sleep(3 * time.Second)

	// Shutdown the log while goroutines are still writing to it.
	log.Println("Shutdown the log.")
	l.Shutdown()

	log.Println("DOWN")
}

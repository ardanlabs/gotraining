// Sample program to show how to use a timer channel and hook
// into the OS using a channel to receive OS events.
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync/atomic"
	"time"
)

// Give the program 10 seconds to complete the work.
const timeoutSeconds = 10 * time.Second

// Flag to indicate to running goroutines to shut
// down the program immediately.
var shutdown int32 = 0

func main() {
	log.Println("Starting Process")

	// Create a channel to talk with the OS.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	// Set the timeout channel.
	timeout := time.After(timeoutSeconds)

	// Launch the process.
	log.Println("Launching Processor")
	complete := make(chan error)
	go processor(complete)

ControlLoop:
	for {
		select {
		case <-sigChan:
			// Interrupt event signaled by the operation system.
			log.Println("OS INTERRUPT - Shutting Down Early")

			// Set the flag to indicate the program should be shutdown early.
			// Try to shutdown cleanly.
			atomic.StoreInt32(&shutdown, 1)
			continue

		case <-timeout:
			// We have taken too much time. Kill the app hard.
			log.Println("Timeout - Killing Program")
			os.Exit(1)

		case err := <-complete:
			// Everything complete within the time given.
			log.Printf("Task Complete: Error[%s]", err)
			break ControlLoop
		}
	}

	// Program finished.
	log.Println("Process Ended")
	return
}

// isShutdown returns the value of the shutdown flag.
func isShutdown() bool {
	value := atomic.LoadInt32(&shutdown)
	if value == 1 {
		return true
	}

	return false
}

// processor provides the main program logic for the program.
func processor(complete chan<- error) {
	log.Println("Processor - Starting")

	// Message returned through the complete channel.
	var err error

	// Schedule this anonymous function to be executed when
	// the function returns.
	defer func() {
		log.Println("Processor - Completed")

		// Signal the goroutine is shutdown.
		complete <- err
	}()

	// Simulate some iterative work.
	for {
		log.Println("Processor - Doing Work")
		time.Sleep(5 * time.Second)

		// Check if we are being asked to shutdown before we
		// complete our work.
		if isShutdown() {
			err = fmt.Errorf("Processor - Shutting Down")
			return
		}
	}
}

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This sample program demonstrates how to use a channel to
// monitor the amount of time the program is running and terminate
// the program if it runs too long.
package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

// Give the program 3 seconds to complete the work.
const timeoutSeconds = 3 * time.Second

var (
	// sigChan receives os signals.
	sigChan = make(chan os.Signal, 1)

	// timeout limits the amount of time the program has.
	timeout = time.After(timeoutSeconds)

	// complete is used to report processing is done.
	complete = make(chan error)

	// shutdown provides system wide notification.
	shutdown = make(chan struct{})
)

func main() {
	log.Println("Starting Process")

	// We want to receive all interrupt based signals.
	signal.Notify(sigChan, os.Interrupt)

	// Launch the process.
	log.Println("Launching Processors")
	go processor(complete)

ControlLoop:
	for {
		select {
		case <-sigChan:

			// Interrupt event signaled by the operation system.
			log.Println("OS INTERRUPT")

			// Close the channel to signal to the processor
			// it needs to shutdown.
			close(shutdown)

			// Set the channel to nil so we no longer process
			// any more of these events.
			sigChan = nil

		case <-timeout:

			// We have taken too much time. Kill the app hard.
			log.Println("Timeout - Killing Program")
			os.Exit(1)

		case err := <-complete:

			// Everything completed within the time given.
			log.Printf("Task Completed: Error[%s]", err)
			break ControlLoop
		}
	}

	// Program finished.
	log.Println("Process Ended")
}

// processor provides the main program logic for the program.
func processor(complete chan<- error) {
	log.Println("Processor - Starting")

	// Variable to store any error that occurs.
	// Passed into the defer function via closures.
	var err error

	// Defer the send on the channel so it happens
	// regardless of how this function terminates.
	defer func() {

		// Capture any potential panic.
		if r := recover(); r != nil {
			log.Println("Processor - Panic", r)
		}

		// Signal the goroutine we have shutdown.
		complete <- err
	}()

	// Perform the work.
	err = doWork()

	log.Println("Processor - Completed")
}

// doWork simulates task work.
func doWork() error {
	log.Println("Processor - Task 1")
	time.Sleep(2 * time.Second)

	if checkShutdown() {
		return errors.New("Early Shutdown")
	}

	log.Println("Processor - Task 2")
	time.Sleep(1 * time.Second)

	if checkShutdown() {
		return errors.New("Early Shutdown")
	}

	log.Println("Processor - Task 3")
	time.Sleep(1 * time.Second)

	return nil
}

// checkShutdown checks the shutdown flag to determine
// if we have been asked to interrupt processing.
func checkShutdown() bool {
	select {
	case <-shutdown:

		// We have been asked to shutdown cleanly.
		log.Println("checkShutdown - Shutdown Early")
		return true

	default:

		// If the shutdown channel was not closed,
		// presume with normal processing.
		return false
	}
}

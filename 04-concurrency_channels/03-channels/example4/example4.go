// http://play.golang.org/p/xIpMqh4Fis

// Sample program to show how to use a timer channel and hook
// into the OS using a channel to receive OS events.
package main

import (
	"log"
	"os"
	"os/signal"
	"time"
)

// Give the program 5 seconds to complete the work.
const timeoutSeconds = 5 * time.Second

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

// main is the entry point for all Go programs.
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
			log.Println("OS INTERRUPT - Shutting Down Early")
			close(shutdown)
			sigChan = nil

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
	for work := 0; work < 5; work++ {
		log.Println("Processor - Doing Work")
		time.Sleep(1 * time.Second)

		select {
		case <-shutdown:
			log.Println("Processor - Shutdown Early")
			// We have been asked to shutdown cleanly.
			return

		default:
			// If the shutdown channel was not closed,
			// presume with normal processing.
		}
	}
}

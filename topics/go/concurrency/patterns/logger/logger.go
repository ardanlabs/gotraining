// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package logger shows a pattern of using a buffer to handle log write
// continuity but deal with write latencies by throwing away log data.
package logger

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Logger provides support to throw log lines away if log
// writes start to timeout due to latency.
type Logger struct {
	write chan string    // Channel to send/recv data to be logged.
	wg    sync.WaitGroup // Helps control the shutdown.

	// Flag to simulate disk issues and latencies.
	full int32
}

// =============================================================================

// New creates a logger value and initializes it for use. The user can
// pass the size of the buffer to use for continuity.
func New(size int) *Logger {

	// Create a value of type logger and init the channel
	// and timer value.
	l := Logger{
		write: make(chan string, size), // Buffered channel if size > 0.
	}

	// Add one to the waitgroup to track the write goroutine.
	l.wg.Add(1)

	// Create the write goroutine that performs the actual
	// writes to disk.
	go func() {

		// Range over the channel and write each data received to disk.
		// Once the channel is close and flushed the loop will terminate.
		for d := range l.write {

			// Simulate disk blocking when signaled.
			for {
				if atomic.LoadInt32(&l.full) == 1 {
					time.Sleep(250 * time.Millisecond)
					continue
				}
				break
			}

			// Simulate write to disk.
			fmt.Println(d)
		}

		// Mark that we are done and terminated.
		l.wg.Done()
	}()

	return &l
}

// Shutdown closes the logger and wait for the writer goroutine
// to terminate.
func (l *Logger) Shutdown() {

	// Close the channel which will cause the write goroutine
	// to finish what is has in its buffer and terminate.
	close(l.write)

	// Wait for the write goroutine to terminate.
	l.wg.Wait()
}

// Write is used to write the data to the log.
func (l *Logger) Write(data string) {

	// Perform the channel operations.
	select {
	case l.write <- data:
		// The writing goroutine got it.

	default:
		// Drop the write.
		fmt.Println("***** DROPPING WRITE")
	}
}

// =============================================================================

// DiskFull can be used by the sample app to create large latencies.
func (l *Logger) DiskFull() {
	if atomic.LoadInt32(&l.full) == 0 {
		atomic.StoreInt32(&l.full, 1)
	} else {
		atomic.StoreInt32(&l.full, 0)
	}
}

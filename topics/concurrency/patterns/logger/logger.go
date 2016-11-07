// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package logger shows a pattern of using a buffer to handle log write
// continuity but deal with write latencies by throwing away log data.
package logger

import (
	"log"
	"sync"
	"sync/atomic"
	"time"
)

// Logger provides support to throw log lines away if log
// writes start to timeout due to latency.
type Logger struct {
	write         chan string    // Channel to send/recv data to be logged.
	timer         *time.Timer    // Timer to deal with latency and timeouts.
	mu            sync.Mutex     // Provides synchronization for log writes.
	wg            sync.WaitGroup // Helps control the shutdown.
	pendingWrites int32          // Counter to identify how many pending writes exist.
	loggingOff    bool           // Flag to indicate the logging is off.

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
		write: make(chan string, size),  // Buffered channel if size > 0.
		timer: time.NewTimer(time.Hour), // Some abitrary large value.
	}

	// Add one to the waitgroup to track
	// the write goroutine.
	l.wg.Add(1)

	// Create the write goroutine that performs the actual
	// writes to disk.
	go func() {

		// Range over the channel and write each data received to disk.
		// Once the channel is close and flushed the loop will terminate.
		for d := range l.write {

			// Help to simulate disk latency issues.
			// WOULD NOT NEED THIS IN PRODUCTION CODE.
			l.pretendDiskFull()

			// Write to disk and decrement the pendingWrites counter.
			log.Println(d)
			atomic.AddInt32(&l.pendingWrites, -1)
		}

		// Mark that we are done and terminated.
		l.wg.Done()
	}()

	return &l
}

// Shutdown closes the logger and wait for the writer goroutine
// to terminate.
func (l *Logger) Shutdown() {
	l.mu.Lock()
	{
		// Set a pending write and turn off logging.
		// We don't want anything else to be written.
		l.pendingWrites = 1
		l.loggingOff = true

		// Close the channel which will cause the write goroutine
		// to finish what is has in its buffer and terminate.
		close(l.write)

		// Wait for the write goroutine to terminate.
		l.wg.Wait()
	}
	l.mu.Unlock()
}

// Write is used to write the data to the log.
func (l *Logger) Write(data string) {
	l.mu.Lock()
	{
		// If logging is off because of latency issues, we will
		// want to throw the data away. Check if the latency issues
		// are gone first.
		if l.loggingOff {

			// If there are pending writes, the buffer is NOT flushed.
			if pw := atomic.LoadInt32(&l.pendingWrites); pw > 0 {
				log.Println("**** DROPPING LOG DATA : ", pw)
				l.mu.Unlock()
				return
			}

			// The buffer has been flushed and the latency is gone.
			l.loggingOff = false
			log.Println("**** LOG WARNING: LOGGING WAS OFF - PLEASE REPORT ****")
		}

		// For now we will not wait longer than 25 millisecond to
		// get our write processed.
		l.timer.Reset(25 * time.Millisecond)

		// Perform the channel operations.
		select {

		// Try to send the data to the write goroutine.
		case l.write <- data:
			atomic.AddInt32(&l.pendingWrites, 1)
			l.timer.Stop()

		// Hit our latency timeout, throw the data away and turn off logging.
		case <-l.timer.C:
			l.loggingOff = true
		}
	}
	l.mu.Unlock()
}

// =============================================================================

// DiskFull is used by the sample app to create large latencies.
func (l *Logger) DiskFull() {
	atomic.StoreInt32(&l.full, 1)
}

// pretendDiskFull checks the full flag and blocks the write to disk
// for one second. Then reset the full flag.
func (l *Logger) pretendDiskFull() {
	if atomic.LoadInt32(&l.full) == 1 {
		time.Sleep(1 * time.Second)
		atomic.StoreInt32(&l.full, 0)
	}
}

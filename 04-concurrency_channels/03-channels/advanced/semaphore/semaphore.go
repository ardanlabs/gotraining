// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/ezGmsjAbiC

// This sample program demonstrates how to implement a semaphore using
// channels that can allow multiple reads but a single write.
//
// It uses the generator pattern to create channels and goroutines.
//
// Multiple reader/writers can be created and run concurrently. Then after
// a timeout period, the program shutdowns cleanly.
//
// http://www.golangpatterns.info/concurrency/semaphores
package main

import (
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type (
	// semaphore type represents a channel that implements the semaphore pattern.
	semaphore chan struct{}
)

type (
	// readerWriter provides a structure for safely reading and writing a shared resource.
	// It supports multiple readers and a single writer goroutine using a semaphore construct.
	readerWriter struct {
		// The name of this object.
		name string

		// write forces reading to stop to allow the write to occur safely.
		write sync.WaitGroup

		// readerControl is a semaphore that allows a fixed number of reader goroutines
		// to read at the same time. This is our semaphore.
		readerControl semaphore

		// shutdown is used to signal to running goroutines to shutdown.
		shutdown chan struct{}

		// reportShutdown is used by the goroutines to report they are shutdown.
		reportShutdown sync.WaitGroup

		// maxReads defined the maximum number of reads that can occur at a time.
		maxReads int

		// maxReaders defines the number of goroutines launched to perform read operations.
		maxReaders int

		// currentReads keeps a safe count of the current number of reads occurring
		// at any give time.
		currentReads int32
	}
)

// init is called when the package is initialized. This code runs first.
func init() {
	// Seed the random number generator
	rand.Seed(time.Now().Unix())
}

// main is the entry point for the application
func main() {
	log.Println("Starting Process")

	// Create a new readerWriter with a max of 3 reads at a time
	// and a total of 6 reader goroutines.
	first := start("First", 3, 6)

	// Create a new readerWriter with a max of 1 reads at a time
	// and a total of 1 reader goroutines.
	second := start("Second", 2, 2)

	// Let the program run for 2 seconds.
	time.Sleep(2 * time.Second)

	// Shutdown both of the readerWriter processes.
	shutdown(first, second)

	log.Println("Process Ended")
	return
}

// start uses the generator pattern to create the readerWriter value. It launches
// goroutines to process the work, returning the created ReaderWriter value.
func start(name string, maxReads int, maxReaders int) *readerWriter {
	// Create a value of readerWriter and initialize.
	rw := readerWriter{
		name:          name,
		shutdown:      make(chan struct{}),
		maxReads:      maxReads,
		maxReaders:    maxReaders,
		readerControl: make(semaphore, maxReads),
	}

	// Launch a number of reader goroutines and let them start reading.
	rw.reportShutdown.Add(maxReaders)
	for goroutine := 0; goroutine < maxReaders; goroutine++ {
		go rw.reader(goroutine)
	}

	// Launch the single writer goroutine and let it start writing.
	rw.reportShutdown.Add(1)
	go rw.writer()

	return &rw
}

// shutdown stops all of the existing readerWriter processes concurrently.
func shutdown(readerWriters ...*readerWriter) {
	// Create a WaitGroup to track the shutdowns.
	var waitShutdown sync.WaitGroup
	waitShutdown.Add(len(readerWriters))

	// Launch each call to the stop method as a goroutine.
	for _, readerWriter := range readerWriters {
		go readerWriter.stop(&waitShutdown)
	}

	// Wait for all the goroutines to report they are done.
	waitShutdown.Wait()
}

// stop signals to all goroutines to shutdown and reports back
// when that is complete
func (rw *readerWriter) stop(waitShutdown *sync.WaitGroup) {
	// Schedule the call to Done for once the method returns.
	defer waitShutdown.Done()

	log.Printf("%s\t: #####> Stop", rw.name)

	// Close the channel which will causes all the goroutines waiting on
	// this channel to receive the notification to shutdown.
	close(rw.shutdown)

	// Wait for all the goroutine to call Done on the waitgroup we
	// are waiting on.
	rw.reportShutdown.Wait()

	log.Printf("%s\t: #####> Stopped", rw.name)
}

// reader is a goroutine that listens on the shutdown channel and
// performs reads until the channel is signaled.
func (rw *readerWriter) reader(reader int) {
	// Schedule the call to Done for once the method returns.
	defer rw.reportShutdown.Done()

	for {
		select {
		case <-rw.shutdown:
			log.Printf("%s\t: #> Reader Shutdown", rw.name)
			return

		default:
			rw.performRead(reader)
		}
	}
}

// performRead performs the actual reading work.
func (rw *readerWriter) performRead(reader int) {
	// Get a read lock for this critical section.
	rw.ReadLock(reader)

	// Safely increment the current reads counter
	count := atomic.AddInt32(&rw.currentReads, 1)

	// Simulate some reading work
	log.Printf("%s\t: [%d] Start\t- [%d] Reads\n", rw.name, reader, count)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	// Safely decrement the current reads counter
	count = atomic.AddInt32(&rw.currentReads, -1)
	log.Printf("%s\t: [%d] Finish\t- [%d] Reads\n", rw.name, reader, count)

	// Release the read lock for this critical section.
	rw.ReadUnlock(reader)
}

// writer is a goroutine that listens on the shutdown channel and
// performs writes until the channel is signaled.
func (rw *readerWriter) writer() {
	// Schedule the call to Done for once the method returns.
	defer rw.reportShutdown.Done()

	for {
		select {
		case <-rw.shutdown:
			log.Printf("%s\t: #> Writer Shutdown", rw.name)
			return
		default:
			rw.performWrite()
		}
	}
}

// performWrite performs the actual write work.
func (rw *readerWriter) performWrite() {
	// Wait a random number of milliseconds before we write again.
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	log.Printf("%s\t: *****> Writing Pending\n", rw.name)

	// Get a write lock for this critical section.
	rw.WriteLock()

	// Simulate some writing work.
	log.Printf("%s\t: *****> Writing Start", rw.name)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("%s\t: *****> Writing Finish", rw.name)

	// Release the write lock for this critical section.
	rw.WriteUnlock()
}

// ReadLock guarantees only the maximum number of goroutines can read at a time.
func (rw *readerWriter) ReadLock(reader int) {
	// If a write is occurring, wait for it to complete.
	rw.write.Wait()

	// Acquire a buffer from the semaphore channel.
	rw.readerControl.Acquire(1)
}

// ReadUnlock gives other goroutines waiting to read their opporunity.
func (rw *readerWriter) ReadUnlock(reader int) {
	// Release the buffer back into the semaphore channel.
	rw.readerControl.Release(1)
}

// WriteLock blocks all reading so the write can happen safely.
func (rw *readerWriter) WriteLock() {
	// Add 1 to the waitGroup so reads will block
	rw.write.Add(1)

	// Acquire all the buffers from the semaphore channel.
	rw.readerControl.Acquire(rw.maxReads)
}

// WriteUnlock releases the write lock and allows reads to occur.
func (rw *readerWriter) WriteUnlock() {
	// Release all the buffers back into the semaphore channel.
	rw.readerControl.Release(rw.maxReads)

	// Release the write lock.
	rw.write.Done()
}

// Acquire attempts to secure the specified number of buffers from the
// semaphore channel.
func (s semaphore) Acquire(buffers int) {
	var e struct{}

	// Write data to secure each buffer.
	for buffer := 0; buffer < buffers; buffer++ {
		s <- e
	}
}

// Release returns the specified number of buffers back into the semaphore channel.
func (s semaphore) Release(buffers int) {
	// Read the data from the channel to release buffers.
	for buffer := 0; buffer < buffers; buffer++ {
		<-s
	}
}

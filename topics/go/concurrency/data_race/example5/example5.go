// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to use a read/write mutex to define critical
// sections of code that needs synchronous access.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// data is a slice that will be shared.
	data []string

	// rwMutex is used to define a critical section of code.
	rwMutex sync.RWMutex

	// Number of reads occurring at ay given time.
	readCount int64
)

// init is called prior to main.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(1)

	// Create the writer goroutine.
	go func() {
		writer()
		wg.Done()
	}()

	// To keep the sample simple we are allowing the runtime to
	// kill the reader goroutines. This is something we should
	// control before allowing main to exit.

	// Create seven reader goroutines.
	for i := 1; i <= 7; i++ {
		go reader(i)
	}

	// Wait for the write goroutine to finish.
	wg.Wait()
	fmt.Println("Program Complete")
}

// writer adds 10 new strings to the slice in random intervals.
func writer() {
	for i := 1; i <= 10; i++ {

		// Only allow one goroutine to read/write to the
		// slice at a time.
		rwMutex.Lock()
		{
			// Capture the current read count.
			// Keep this safe though we can due without this call.
			rc := atomic.LoadInt64(&readCount)

			// Perform some work since we have a full lock.
			fmt.Printf("****> : Performing Write : RCount[%d]\n", rc)
			data = append(data, fmt.Sprintf("String: %d", i))
		}
		rwMutex.Unlock()
		// Release the lock and allow any waiting goroutines
		// to continue using the slice.

		// Sleep a random amount of time.
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

// reader wakes up and iterates over the data slice.
func reader(id int) {
	for {

		// Any goroutine can read when no write
		// operation is taking place.
		rwMutex.RLock()
		{
			// Increment the read count value by 1.
			rc := atomic.AddInt64(&readCount, 1)

			// Perform some read work and display values.
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
			fmt.Printf("%d : Performing Read : Length[%d] RCount[%d]\n", id, len(data), rc)

			// Decrement the read count value by 1.
			atomic.AddInt64(&readCount, -1)
		}
		rwMutex.RUnlock()
		// Release the read lock.
	}
}

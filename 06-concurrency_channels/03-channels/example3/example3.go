// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/aHI23AlFD7

// This sample program demonstrates how to use a buffered
// channel to receive results from other goroutines in a guaranteed way.
package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// numInserts of inserts to perform.
const numInserts = 10

// init called before main.
func init() {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())
}

// main is the entry point for all Go programs.
func main() {
	performInserts()
}

// performInserts coordinates the possible inserts that need to take place.
func performInserts() {
	log.Println("Inserts Started")

	// Waitgroup to know when all inserts are complete.
	var wg sync.WaitGroup

	// Buffered channel to receive information about any possible insert.
	ch := make(chan error, numInserts)

	// Perform any possible number of inserts.
	for i := 0; i < numInserts; i++ {
		// Do we need to insert document A?
		if isNecessary() {
			wg.Add(1)
			go func(id int) {
				ch <- insertDoc(id)
				wg.Done()
			}(i)
		}
	}

	// Wait to be told all the inserts are done.
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Process the insert results as they complete. Wait for
	// the channel to be closed.
	for err := range ch {
		if err != nil {
			log.Println("Received error:", err)
			continue
		}

		log.Println("Received nil error")
	}

	log.Println("Inserts Complete")
}

// insertDoc simulates a database operation.
func insertDoc(id int) error {
	log.Println("Insert document: ", id)

	// Randomize if the insert fails or not.
	if rand.Intn(2)%2 == 0 {
		return fmt.Errorf("Document ID: %d", id)
	}

	return nil
}

// isNecessary determine if we need to perform the insert.
func isNecessary() bool {
	// Randomize if this insert is necessary or not.
	if rand.Intn(2)%2 == 0 {
		return true
	}

	return false
}

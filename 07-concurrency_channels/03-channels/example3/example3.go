// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/jT4-vZBpMm

// This sample program demonstrates how to use a buffered
// channel to receive results from other goroutines in a guaranteed way.
package main

import (
	"fmt"
	"log"
	"math/rand"
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

	// Buffered channel to receive information about any possible insert.
	ch := make(chan error, numInserts)

	// Number of responses we need to handle.
	var waitResponses int

	// Perform any possible number of inserts.
	for i := 0; i < numInserts; i++ {
		// Do we need to insert document A?
		if isNecessary() {
			waitResponses++
			go func(id int) {
				ch <- insertDoc(id)
			}(i)
		}
	}

	// Process the insert results as they complete.
	for waitResponses > 0 {
		// Wait for a response from a goroutine.
		err := <-ch

		// Display the result.
		if err != nil {
			log.Println("Received error:", err)
		} else {
			log.Println("Received nil error")
		}

		// Decrement the wait count and determine if we are done.
		waitResponses--
	}

	log.Println("Inserts Complete")
}

// insertDoc simulates a database operation.
func insertDoc(id int) error {
	log.Println("Insert document: ", id)

	// Randomize if the insert fails or not.
	if rand.Intn(2) == 0 {
		return fmt.Errorf("Document ID: %d", id)
	}

	return nil
}

// isNecessary determine if we need to perform the insert.
func isNecessary() bool {
	// Randomize if this insert is necessary or not.
	if rand.Intn(2) == 0 {
		return true
	}

	return false
}

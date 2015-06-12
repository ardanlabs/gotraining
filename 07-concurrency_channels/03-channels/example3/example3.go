// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/T3QupG_7_X

// This sample program demonstrates how to use a buffered
// channel to receive results from other goroutines in a guaranteed way.
package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

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

	// Set the number of routines and inserts.
	const routines = 10
	const inserts = routines * 2

	// Buffered channel to receive information about any possible insert.
	ch := make(chan error, inserts)

	// Number of responses we need to handle.
	waitInserts := inserts

	// Perform all the inserts.
	for i := 0; i < routines; i++ {
		go func(id int) {
			ch <- insertDoc1(id)
			ch <- insertDoc2(id)
		}(i)
	}

	// Process the insert results as they complete.
	for waitInserts > 0 {
		// Wait for a response from a goroutine.
		err := <-ch

		// Display the result.
		if err != nil {
			log.Println("Received error:", err)
		} else {
			log.Println("Received nil error")
		}

		// Decrement the wait count and determine if we are done.
		waitInserts--
	}

	log.Println("Inserts Complete")
}

// insertDoc1 simulates a database operation.
func insertDoc1(id int) error {
	log.Println("Insert document 1: ", id)

	// Randomize if the insert fails or not.
	if rand.Intn(10) == 0 {
		return fmt.Errorf("Document ID: %d", id)
	}

	return nil
}

// insertDoc2 simulates a database operation.
func insertDoc2(id int) error {
	log.Println("Insert document 2: ", id)

	// Randomize if the insert fails or not.
	if rand.Intn(10) == 0 {
		return fmt.Errorf("Document ID: %d", id)
	}

	return nil
}

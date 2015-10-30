// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This sample program demonstrates how to use the work package
// to use a pool of goroutines to get work done.
package main

import (
	"log"
	"sync"
	"time"

	"github.com/ardanlabs/gotraining/09-concurrency_patterns/task"
)

// names provides a set of names to display.
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

// namePrinter provides special support for printing names.
type namePrinter struct {
	name string
}

// Work implements the Worker interface.
func (m namePrinter) Work() {
	log.Println(m.name)
	time.Sleep(3 * time.Second)
}

// main is the entry point for all Go programs.
func main() {
	// Create a task pool with 4 goroutines.
	t := task.New(10)

	var wg sync.WaitGroup
	wg.Add(10 * len(names))

	for i := 0; i < 10; i++ {
		// Iterate over the slice of names.
		for _, name := range names {
			// Create a namePrinter and provide the
			// specific name.
			np := namePrinter{
				name: name,
			}

			go func() {
				// Submit the task to be worked on. When Do
				// returns, we know it is being handled.
				t.Do(np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	// Shutdown the task pool and wait for all existing work
	// to be completed.
	t.Shutdown()
}

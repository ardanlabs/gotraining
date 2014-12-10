// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// This example is provided with help by Jason Waldrip.

// This sample program demostrates how to use the work package
// to use a pool of goroutines to get work done.
package main

import (
	"fmt"
	"time"

	"github.com/ArdanStudios/gotraining/07-concurrency_patterns/work"
)

// names provides a set of names to display.
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
	"kelly",
	"paul",
	"dina",
	"chris",
	"lisa",
	"tom",
	"travis",
}

// namePrinter provides special support for printing names.
type namePrinter struct {
	name string
}

// Work implements the Worker interface.
func (m *namePrinter) Work() {
	fmt.Println(m.name)
	time.Sleep(time.Second)
}

// main is the entry point for all Go programs.
func main() {
	// Create a work value with 2 goroutines.
	w := work.New(2)

	// Iterate over the slice of names.
	for _, name := range names {
		// Create a namePrinter and provide the
		// specfic name.
		np := namePrinter{
			name: name,
		}

		// Submit the task to be worked on. When RunTask
		// returns we know it is being handled.
		w.RunTask(&np)
	}

	// Shutdown the work and wait for all existing work
	// to be completed.
	w.Shutdown()
}

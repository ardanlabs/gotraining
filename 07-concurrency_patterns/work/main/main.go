// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// This sample program demostrates how to use the work package
// to use a pool of goroutines to get work done.
package main

import (
	"fmt"
	"log"
	"sync"
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
}

// namePrinter provides special support for printing names.
type namePrinter struct {
	name string
}

// Work implements the Worker interface.
func (m *namePrinter) Work(id int) {
	fmt.Println(m.name)
	time.Sleep(time.Second)
}

// logger is called by the work pool when
// it has things to log.
func logger(message string) {
	log.Println(message)
}

// main is the entry point for all Go programs.
func main() {
	// Create a work value with 2 goroutines.
	w, err := work.New(2, time.Second, logger)
	if err != nil {
		return
	}

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		// Iterate over the slice of names.
		for _, name := range names {
			// Create a namePrinter and provide the
			// specfic name.
			np := namePrinter{
				name: name,
			}

			go func() {
				// Submit the task to be worked on. When RunTask
				// returns we know it is being handled.
				w.Run(&np)
				wg.Done()
			}()
		}
	}

	for {
		// Enter a number and hit enter to change the size
		// of the work pool.
		var c int
		fmt.Scanf("%d", &c)
		if c == 0 {
			break
		}

		w.Add(c)
	}

	wg.Wait()

	// Shutdown the work and wait for all existing work
	// to be completed.
	w.Shutdown()
}

// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/K9gNyTRA0s

// Write a problem that uses a buffered channel to maintain a buffer
// of four strings. In main, send the strings 'A', 'B', 'C' and 'D'
// into the channel. Then create 20 goroutines that receive a string
// from the channel, display the value and then send the string back
// into the channel. Once each goroutine is done performing that task,
// allow the goroutine to terminate.
package main

import (
	"fmt"
	"sync"
)

const (
	goroutines = 20
	capacity   = 4
)

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

// resources is a buffered channel to manage strings.
var resources = make(chan string, capacity)

// main is the entry point for all Go programs.
func main() {
	// Launch goroutines to handle the work.
	wg.Add(goroutines)
	for gr := 1; gr <= goroutines; gr++ {
		go worker(gr)
	}

	// Add the strings.
	for rune := 'A'; rune < 'A'+capacity; rune++ {
		resources <- string(rune)
	}

	// Wait for all the work to get done.
	wg.Wait()
}

// worker is launched as a goroutine to process work from
// the buffered channel.
func worker(worker int) {
	// Receive a string from the channel.
	value := <-resources

	// Display the value.
	fmt.Printf("Worker: %d : %s\n", worker, value)

	// Place the string back.
	resources <- value

	// Tell main we are done.
	wg.Done()
}

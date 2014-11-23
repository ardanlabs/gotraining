// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/B9npiUVveE

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

var (
	// wg is used to wait for the program to finish.
	wg sync.WaitGroup

	// resources is a buffered channel to manage resources.
	resources = make(chan string, capacity)
)

// main is the entry point for all Go programs.
func main() {
	// Launch goroutines to handle the work.
	wg.Add(goroutines)
	for gr := 1; gr <= goroutines; gr++ {
		go worker(gr)
	}

	// Add the resources.
	for rune := 'A'; rune < 'A'+capacity; rune++ {
		resources <- string(rune)
	}

	// Wait for all the work to get done.
	wg.Wait()
}

// worker is launched as a goroutine to process work from
// the buffered channel.
func worker(worker int) {
	// Report that we just returned.
	defer wg.Done()

	// Receive a resource from the channel.
	value, ok := <-resources
	if !ok {
		// This means the channel is empty and closed.
		fmt.Printf("Worker: %d : Shutting Down\n", worker)
		return
	}

	// Display the value.
	fmt.Printf("Worker: %d : %s\n", worker, value)

	// Place the resource back.
	resources <- value
}

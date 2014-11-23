// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/viYA-f4zBI

// Sample program to show how the goroutine scheduler
// will time slice goroutines on a single thread.
package main

import (
	"fmt"
	"sync"
)

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

// main is the entry point for all Go programs.
func main() {
	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// Create two goroutines.
	fmt.Println("Create Goroutines")
	go printPrime("A")
	go printPrime("B")

	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

// printPrime displays prime numbers for the first 5000 numbers.
func printPrime(prefix string) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

next:
	for outer := 1; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}

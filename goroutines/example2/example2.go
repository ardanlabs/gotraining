// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/SvJj6T4Jhi

// Sample program to show how the goroutine scheduler
// will time slice goroutines on a single thread.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

// init is called prior to main.
func init() {
	// Allocate one logical processor for the scheduler to use.
	runtime.GOMAXPROCS(1)
}

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
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}

	fmt.Println("Completed", prefix)

	// Tell main we are done.
	wg.Done()
}

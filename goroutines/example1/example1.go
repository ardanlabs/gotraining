// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/ki1woWvmzW

// Sample program to show how to create goroutines and
// how the scheduler behaves.
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

	fmt.Println("Start Goroutines")

	// Create a goroutine from the lowercase function.
	go lowercase()

	// Create a goroutine from the uppercase function.
	go uppercase()

	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

// lowercase displays the set of lowercase letters three times.
func lowercase() {
	// Display the alphabet three times
	for count := 0; count < 3; count++ {
		for rune := 'a'; rune < 'a'+26; rune++ {
			fmt.Printf("%c ", rune)
		}
	}

	// Tell main we are done.
	wg.Done()
}

// uppercase displays the set of uppercase letters three times.
func uppercase() {
	// Display the alphabet three times
	for count := 0; count < 3; count++ {
		for rune := 'A'; rune < 'A'+26; rune++ {
			fmt.Printf("%c ", rune)
		}
	}

	// Tell main we are done.
	wg.Done()
}

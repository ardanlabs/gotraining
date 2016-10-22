// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to create goroutines and
// how the goroutine scheduler behaves with two contexts.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {

	// Allocate two logical processors for the scheduler to use.
	runtime.GOMAXPROCS(2)
}

func main() {

	// wg is used to wait for the program to finish.
	// Add a count of two, one for each goroutine.
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// Declare an anonymous function and create a goroutine.
	go func() {

		// Display the alphabet three times.
		for count := 0; count < 3; count++ {
			for rune := 'a'; rune < 'a'+26; rune++ {
				fmt.Printf("%c ", rune)
			}
		}

		// Tell main we are done.
		wg.Done()
	}()

	// Declare an anonymous function and create a goroutine.
	go func() {

		// Display the alphabet three times.
		for count := 0; count < 3; count++ {
			for rune := 'A'; rune < 'A'+26; rune++ {
				fmt.Printf("%c ", rune)
			}
		}

		// Tell main we are done.
		wg.Done()
	}()

	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/sFfYEQQJFW

// Sample program to show how to create goroutines and
// how the scheduler behaves.
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
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	// Display the alphabet three times
	for count := 0; count < 3; count++ {
		for rune := 'a'; rune < 'a'+26; rune++ {
			fmt.Printf("%c ", rune)
		}
	}
}

// uppercase displays the set of uppercase letters three times.
func uppercase() {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	// Display the alphabet three times
	for count := 0; count < 3; count++ {
		for rune := 'A'; rune < 'A'+26; rune++ {
			fmt.Printf("%c ", rune)
		}
	}
}

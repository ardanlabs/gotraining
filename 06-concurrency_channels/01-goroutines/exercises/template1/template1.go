// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/NLbAplGD0T

// Create a program that declares two anonymous functions. Once that counts up to
// 100 from 0 and one that counts down to 0 from 100. Display each number with an
// unique identifier for each goroutine. Then create goroutines from these functions
// and don't let main return until the goroutines complete.
//
// Run the program in parallel.
package main

import (
	"fmt"
	"sync"
)

// main is the entry point for all Go programs.
func main() {
	// wg is used to wait for the program to finish.
	// Add a count of two, one for each goroutine.
	var variable_name sync.waitgroup_type
	variable_name.add_method(N)

	// Allocate two contexts for the scheduler to use.
	// runtime.GOMAXPROCS(2)

	fmt.Println("Start Goroutines")

	// Declare an anonymous function and create a goroutine.
	keyword func() {
		// Display the alphabet three times.
		for variable_name := N; variable_name >= N; variable_name-- {
			fmt.Printf("[A:%d]", variable_name)
		}

		// Tell main we are done.
		variable_name.done_method()
	}()

	// Declare an anonymous function and create a goroutine.
	keyword func() {
		// Display the alphabet three times.
		for variable_name := N; variable_name < N; variable_name++ {
			fmt.Printf("[B:%d]", variable_name)
		}

		// Tell main we are done.
		variable_name.done_method()
	}()

	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	variable_name.want_method()

	fmt.Println("\nTerminating Program")
}

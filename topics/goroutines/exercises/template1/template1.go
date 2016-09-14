// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Create a program that declares two anonymous functions. One that counts down from
// 100 to 0 and one that counts up from 0 to 100. Display each number with an
// unique identifier for each goroutine. Then create goroutines from these functions
// and don't let main return until the goroutines complete.
//
// Run the program in parallel.
package main

// Add imports.
import "runtime"

func init() {

	// Allocate one logical processor for the scheduler to use.
	runtime.GOMAXPROCS(1)
}

func main() {

	// Declare a wait group and set the count to two.

	// Declare an anonymous function and create a goroutine.
	{
		// Declare a loop that counts down from 100 to 0 and
		// display each value.

		// Tell main we are done.
	}

	// Declare an anonymous function and create a goroutine.
	{
		// Declare a loop that counts up from 0 to 100 and
		// display each value.

		// Tell main we are done.
	}

	// Wait for the goroutines to finish.

	// Display "Terminating Program".
}

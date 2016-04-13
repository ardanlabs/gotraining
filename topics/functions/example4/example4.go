// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/mmhY-UL7KQ

// Sample program to show how anonymous functions and closures work.
package main

import "fmt"

// main is the entry point for the application.
func main() {
	var n int

	// Defer the call to fmt.Println till after main returns.
	defer fmt.Println("2:", n)

	// Defer the call to the anonymous function till after main returns.
	defer func() {
		fmt.Println("1:", n)
	}()

	// Set the value of n to 3 before the return.
	n = 3
}

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how anonymous functions and closures work.
package main

import "fmt"

func main() {
	var n int

	// Declare an anonymous function and call it.
	func() {
		fmt.Println("Direct:", n)
	}()

	// Declare an anonymous function and assign it to a variable.
	f := func() {
		fmt.Println("Variable:", n)
	}

	// Call the anonymous function through the variable.
	f()

	// Defer the call to the anonymous function till after main returns.
	defer func() {
		fmt.Println("Defer 1:", n)
	}()

	// Set the value of n to 3 before the return.
	n = 3

	// Call the anonymous function through the variable.
	f()

	// Defer the call to the anonymous function till after main returns.
	defer func() {
		fmt.Println("Defer 2:", n)
	}()
}

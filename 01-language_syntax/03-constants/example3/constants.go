// Sample program to show how constants do have a parallel type system.
package main

import "fmt"

// Much larger value than int64.
const bigger int64 = 9223372036854775808543522345

// main is the entry point for the application.
func main() {
	fmt.Println("Will NOT Compile")
}

// Compiler Error:
// ./constants.go:6: constant 9223372036854775808543522345 overflows int64

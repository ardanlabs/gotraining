// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show the syntax and mechanics of type
// switches and the empty interface.
package main

import "fmt"

func main() {

	// fmt.Println can be called with values of any type
	fmt.Println("Hello, world")
	fmt.Println(12345)
	fmt.Println(3.14159)
	fmt.Println(true)

	// How can we do the same?
	myPrintln("Hello, world")
	myPrintln(12345)
	myPrintln(3.14159)
	myPrintln(true)

	// - An interface is satisfied by any type with certain methods.
	// - The empty interface has no methods.
	// - All types have AT LEAST no methods.
	// Therefore all types satisfy the empty interface.

	// - The empty interface says nothing useful for the compiler.
	// Checks must be performed at runtime to inspect the data
	// stored in the interface variable.
	// - Abstracting around the shape of data is costly both in
	// terms of performance and complexity.
	// - Abstract around behavior when possible.
}

func myPrintln(a interface{}) {

	switch a.(type) {
	case string:
		fmt.Println("Passed a string")
	case int:
		fmt.Println("Passed an int")
	case float64:
		fmt.Println("Passed a float64")
	default:
		fmt.Println("Passed something else")
	}

	switch v := a.(type) {
	case string:
		fmt.Println("[string ]: ", v)
	case int:
		fmt.Println("[int    ]:", v)
	case float64:
		fmt.Println("[float64]:", v)
	default:
		fmt.Println("[???????]:,", v)
	}
}

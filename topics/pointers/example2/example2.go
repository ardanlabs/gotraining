// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show the basic concept of using a pointer
// to share data.
package main

import "fmt"

func main() {

	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	println("Before:", count, &count)

	// Pass the "address of" the variable count.
	increment(&count)

	println("After: ", count, &count)
}

// increment declares count as a pointer variable whose value is
// always an address and points to values of type int.
func increment(inc *int) {

	// Increment the value that the "pointer points to". (de-referencing)
	*inc++
	println("Inc:   ", *inc, &inc, inc)

	// Do this to prevent inlining.
	var x int
	fmt.Sprintf("Prevent Inlining: %d", x)
}

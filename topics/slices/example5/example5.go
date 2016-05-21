// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how one needs to be careful when appending
// to a slice when you have a reference to an element.
package main

import "fmt"

func main() {

	// Declare a slice of integers with 7 values.
	x := []int{100, 200, 300, 400, 500, 600, 700}

	// Set a pointer to the second element of the slice.
	twohundred := &x[1]

	// Append a new value to the slice.
	x = append(x, 800)

	// Change the value of the second element of the slice.
	x[1] = 250

	// Display the value that the pointer points to and the
	// second element of the slice.
	fmt.Println("Pointer:", *twohundred, "Element", x[1])
}

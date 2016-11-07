// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how the capacity of the slice
// is not available for use.
package main

import "fmt"

func main() {

	// Create a slice with a length of 5 elements.
	slice := make([]string, 5)
	slice[0] = "Apple"
	slice[1] = "Orange"
	slice[2] = "Banana"
	slice[3] = "Grape"
	slice[4] = "Plum"

	// You can't access an index of a slice beyond its length.
	slice[5] = "Runtime error"

	// Error: panic: runtime error: index out of range

	fmt.Println(slice)
}

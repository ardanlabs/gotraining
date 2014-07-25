// http://play.golang.org/p/zU_00KiuyS

// Sample program to show how the capacity of the slice
// is not available for use.
package main

import (
	"fmt"
)

// main is the entry point for the application.
func main() {
	// Create a slice with a length of 5 elements and a capacity of 8.
	slice := make([]string, 5, 8)
	slice[0] = "Apple"
	slice[1] = "Orange"
	slice[2] = "Banana"
	slice[3] = "Grape"
	slice[4] = "Plum"

	// You can't access an element of a slice beyond its length.
	slice[5] = "Runtime error"

	// Error: panic: runtime error: index out of range

	fmt.Println(slice)
}

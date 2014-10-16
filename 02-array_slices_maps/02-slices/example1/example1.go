// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/4r90uFQwJn

// Sample program to show how the capacity of the slice
// is not available for use.
package main

import (
	"fmt"
)

// main is the entry point for the application.
func main() {
	// Create a slice with a length of 5 elements.
	slice := make([]string, 5)
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

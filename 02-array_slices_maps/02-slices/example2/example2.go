// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/DB8hwJ0hw9

// Sample program to show the components of a slice. It has a
// length, capacity and the underlying array.
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

	inspectSlice(slice)
}

// inspectSlice exposes the slice header for review.
func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for index, value := range slice {
		fmt.Printf("[%d] %p %s\n",
			index,
			&slice[index],
			value)
	}
}

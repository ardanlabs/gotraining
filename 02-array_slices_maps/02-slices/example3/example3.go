// http://play.golang.org/p/PyZthd9EFl

// Sample program to show how to takes slices of slices to create different
// views of and make changes to the underlying array.
package main

import (
	"fmt"
)

// main is the entry point for the application.
func main() {
	// Create a slice with a length of 5 elements and a capacity of 8.
	slice1 := make([]string, 5, 8)
	slice1[0] = "Apple"
	slice1[1] = "Orange"
	slice1[2] = "Banana"
	slice1[3] = "Grape"
	slice1[4] = "Plum"

	inspectSlice(slice1)

	// For slice[i:j] with an underlying array of capacity k
	// Length: j - i
	// Capacity: k - i

	// Take a slice of slice1. We want just elements 2 and 3.
	// slice2[0] = "Banana"
	// slice2[1] = "Grape"
	// Length:   4 - 2
	// Capacity: 8 - 2
	slice2 := slice1[2:4]
	inspectSlice(slice2)

	// Take a slice of slice2. We want all the elements from 1.
	// slice3[0] = "Grape"
	// slice3[1] = "Plum"
	// Length:   6 - 1
	// Capacity: 6 - 1
	slice3 := slice2[1:cap(slice2)]
	inspectSlice(slice3)

	fmt.Println("*************************")

	// Change the value of the first element of slice3.
	slice3[0] = "CHANGED"

	// Display the change across all existing slices.
	inspectSlice(slice1)
	inspectSlice(slice2)
	inspectSlice(slice3)
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

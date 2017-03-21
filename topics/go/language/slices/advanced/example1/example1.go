// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to use a third index slice.
package main

import "fmt"

func main() {

	// Create a slice of strings with different types of fruit.
	slice := []string{"Apple", "Orange", "Banana", "Grape", "Plum"}
	inspectSlice(slice)

	// Take a slice of slice. We want just index 2
	takeOne := slice[2:3]
	inspectSlice(takeOne)

	// Take a slice of just index 2 with a length and capacity of 1
	takeOneCapOne := slice[2:3:3] // Use the third index position to
	inspectSlice(takeOneCapOne)   // set the capacity to 1.

	// Append a new element which will create a new
	// underlying array to increase capacity.
	takeOneCapOne = append(takeOneCapOne, "Kiwi")
	inspectSlice(takeOneCapOne)
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

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show the components of a slice. It has a
// length, capacity and the underlying array.
package main

import "fmt"

func main() {

	// Create a slice with a length of 5 elements and a capacity of 8.
	fruits := make([]string, 5, 8)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	// ----- ruben
	for i := range fruits {
		fmt.Println(fruits[i], &fruits[i])
	}

	fmt.Printf("\n\n")
	// ----- ruben

	inspectSlice(fruits)
}

// inspectSlice exposes the slice header for review.
func inspectSlice(slice []string) {
	// ----- ruben
	for i := range slice {
		fmt.Println(slice[i], &slice[i])
	}

	fmt.Printf("\n\n")
	// ----- ruben

	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i, s := range slice {
		fmt.Printf("[%d] %p %s %p\n",
			i,
			&slice[i],
			s,
			&s) // I added this '&s'
	}
}

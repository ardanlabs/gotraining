// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how the range will make a copy of the supplied
// slice value when the second value is requested during iteration.
package main

import "fmt"

func main() {

	// Create a slices of strings.
	a := []string{"a1", "a2", "a3", "a4", "a5"}
	b := []string{"b1", "b2", "b3", "b4", "b5"}

	// The range will make a copy of the slice header.
	for i, v := range a {

		// Replace slice a for slice b but it won't matter. We are
		// iterating over our own copy of the slice header for a.
		a = b

		// We continue to iterate over the orginal backing array
		// for the a slice.
		fmt.Println(i, v)
	}

	// The range will make a copy of the slice header.
	for i, v := range a {

		// Slice off the last two element but it won't matter. We are
		// iterating over our own copy of the slice header for a.
		a = a[:3]

		// We continue to iterate over the orginal backing array
		// for the a slice.
		fmt.Println(i, v)
	}
}

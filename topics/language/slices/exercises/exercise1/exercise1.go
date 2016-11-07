// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

import "fmt"

func main() {

	// Declare a nil slice of integers.
	var numbers []int

	// Append numbers to the slice.
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i*10)
	}

	// Display each value.
	for _, number := range numbers {
		fmt.Println(number)
	}

	// Declare a slice of strings.
	names := []string{"Bill", "Lisa", "Jim", "Cathy", "Beth"}

	// Display each index position and name.
	for i, name := range names {
		fmt.Printf("Index: %d  Name: %s\n", i, name)
	}

	// Take a slice of index 1 and 2.
	slice := names[1:3]

	// Display the value of the new slice.
	for i, name := range slice {
		fmt.Printf("Index: %d  Name: %s\n", i, name)
	}
}

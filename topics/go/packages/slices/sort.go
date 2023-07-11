// This program showcases
// the `slices` package's sort function.
// The aim of this test is to sort an array of
// integers to ascending order.
// This program requires Go 1.21rc1
package main

import (
	"fmt"
	"slices"
)

func main() {

	// This example demonstrates
	// how the sort function can be used
	// with a slice of numbers.
	ints := []int{
		2, 4, 6, 7, 8, 9, 1, 0,
	}

	fmt.Println("Original", ints)

	slices.Sort(ints)

	fmt.Println("Sorted", ints)

}

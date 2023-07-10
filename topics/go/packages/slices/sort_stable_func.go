// This program showcases
// the `slices` package's SortStableFunc function.
// The aim of this test is to define a sort function
// and use it to sort a slice of integers in ascending
// order.
// This program requires Go 1.21rc1
package main

import (
	"fmt"
	"slices"
)

func main() {

	// This example demonstrates
	// how sort stable func can be used
	// with a slice of numbers.
	ints := []int{
		2, 4, 6, 7, 8, 9, 1, 0,
	}[:]

	var less = func(a, b int) int { return a - b }
	slices.SortStableFunc(ints, less)

	fmt.Println("Is slice int sorted", slices.IsSorted(ints))
	fmt.Println("Resulting slice after sort", ints)

}

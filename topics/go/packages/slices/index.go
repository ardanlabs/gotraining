// This program showcases
// the `slices` package's index function.
// The aim of this test is to determine
// the index of an element within a slice
// This program requires Go 1.21rc1
package main

import (
	"fmt"
	"slices"
)

func main() {

	a := []int{
		1, 2, 3, 4, 5,
	}

	c := []int{
		6, 6, 6, 6, 6,
	}

	fmt.Println(
		"Index of 4 in slice a",
		slices.Index(a, 4),
	)

	// This next example showcases
	// how the index function will
	// return the first matching element
	// by looking for the number 6
	// within slice C
	fmt.Println(
		"Index of 6 in slice c",
		slices.Index(c, 6),
	)
}

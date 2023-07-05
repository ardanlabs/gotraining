// This program showcases
// the `slices` package's compare function.
// The aim of this test is unclear
// Need additional clarification
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

	b := []int{
		1, 2, 6, 4, 5,
	}

	c := []int{
		1, 2, 3, 4, 5,
	}

	fmt.Println(
		"Compare Slice a and b",
		slices.Compare(a, b),
	)

	fmt.Println(
		"Compare Slice a and c",
		slices.Compare(a, c),
	)
}

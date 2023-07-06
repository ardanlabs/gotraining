// This program showcases
// the `slices` package's compare function.
// The aim of this test is to leverage
// the compare function to determine
// which array's length is greater
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
		1, 2, 6, 4, 5, 6,
	}

	c := []int{
		1, 2, 3, 4,
	}

	d := []int{
		1, 2, 3, 4,
	}

	// d is short for
	// dictionary and translates
	// the output from the compare
	// function into something that is
	// human readable.
	dict := map[int]string{
		-1: "First slice is shorter",
		0:  "Both slices are equal",
		1:  "Second slice is shorter",
	}
	fmt.Println(
		"Compare Slice a and b",
		dict[slices.Compare(a, b)],
	)

	fmt.Println(
		"Compare Slice a and c",
		dict[slices.Compare(a, c)],
	)

	fmt.Println(
		"Compare Slice c and d",
		dict[slices.Compare(c, d)],
	)
}

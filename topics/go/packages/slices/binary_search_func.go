// This program showcases
// the `slices` package's binary search func function
// The aim of this program is to perform
// a binary search,
// a search on an array that is already
// sorted in ascending order.
// This program requires Go 1.21rc1
package main

import (
	"fmt"
	"slices"
)

func main() {

	ints := []int{
		1, 2, 4, 5, 6,
	}

	cmp := func(a, b int) int {
		return a - b
	}

	fmt.Println("Slice ints", ints)

	_, found := slices.BinarySearchFunc(
		ints,
		3,
		cmp,
	)

	if !found {
		fmt.Println("3 not found in slice")
	}

	index, found := slices.BinarySearchFunc(
		ints,
		5,
		cmp,
	)

	if found {
		fmt.Println("found element 5 at index:", index)
	}

}

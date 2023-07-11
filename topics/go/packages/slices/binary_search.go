// This program showcases
// the `slices` package's binarySearch function
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

	a := []int{
		1, 2, 4, 5, 6,
	}

	fmt.Println("Slice a", a)

	_, found := slices.BinarySearch(
		a,
		3,
	)

	if !found {
		fmt.Println("3 not found in slice")
	}

	index, found := slices.BinarySearch(
		a,
		5,
	)

	if found {
		fmt.Println("found element 5 at index:", index)
	}

}

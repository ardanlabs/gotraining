// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show the `slices` package's binarySearch function
// and search an array, that is already sorted in ascending order, to find the specified element.
package main

import (
	"fmt"

	"golang.org/x/exp/slices"
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

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program shows how to use the BinarySearch APIs from the
// slices package.
package main

import (
	"fmt"
	"slices"
)

// BinarySearch searches for target in a sorted slice and returns the
// position where target is found, or the position where target would
// appear in the sort order.

func main() {
	list := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Slice", list)

	// -------------------------------------------------------------------------
	// BinarySearch

	index, found := slices.BinarySearch(list, 9)
	fmt.Printf("Looking for 9, idx[%d], found[%v]\n", index, found)

	index, found = slices.BinarySearch(list, 5)
	fmt.Printf("Looking for 5, idx[%d], found[%v]\n", index, found)

	// -------------------------------------------------------------------------
	// BinarySearchFunc

	index, found = slices.BinarySearchFunc(list, 7, compare)
	fmt.Printf("Looking for 7, idx[%d], found[%v]\n", index, found)

	index, found = slices.BinarySearchFunc(list, 2, compare)
	fmt.Printf("Looking for 2, idx[%d], found[%v]\n", index, found)
}

// Compare needs to return 0 if the two values are the same, a positive
// number of a > b, and a negative number of a < b.
func compare(a int, b int) int {
	return a - b
}

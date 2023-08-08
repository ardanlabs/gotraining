// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program shows how to use the Compare API from the slices package.
package main

import (
	"fmt"
	"slices"
)

// result translates the result of the compare API.
var result = map[int]string{
	-1: "First slice is shorter",
	0:  "Both slices are equal",
	1:  "Second slice is shorter",
}

// Compare compares the elements between two slices. The elements are compared
// sequentially, starting at index 0, until one element is not equal to the
// other. The result of comparing the first non-matching elements is returned.

func main() {
	list1 := []int{1, 2, 3, 4, 5}
	list2 := []int{1, 2, 6, 4, 5, 6}
	list3 := []int{1, 2, 3, 4}
	list4 := []int{1, 2, 3, 4}

	fmt.Println("Slice1", list1)
	fmt.Println("Slice2", list2)
	fmt.Println("Slice3", list3)
	fmt.Println("Slice4", list4)

	// -------------------------------------------------------------------------
	// Compare list1 and list2

	fmt.Printf("list1 vs list2: Compare(%s), Func(%s)\n",
		result[slices.Compare(list1, list2)],
		result[slices.CompareFunc(list1, list2, compare)],
	)

	// -------------------------------------------------------------------------
	// Compare list1 and list3

	fmt.Printf("list1 vs list3: Compare(%s), Func(%s)\n",
		result[slices.Compare(list1, list3)],
		result[slices.CompareFunc(list1, list3, compare)],
	)

	// -------------------------------------------------------------------------
	// Compare list3 and list4

	fmt.Printf("list3 vs list4: Compare(%s), Func(%s)\n",
		result[slices.Compare(list3, list4)],
		result[slices.CompareFunc(list3, list4, compare)],
	)
}

// compare evaluates values in increasing index order, and the comparisons stop
// after the first time the function returns non-zero. Return 0 is the two
// values match, return -1 if a < b, and 1 if a > b.
func compare(a int, b int) int {
	if a < b {
		return -1
	}

	if a > b {
		return 1
	}

	return 0
}

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program shows how to use the Equal API from the slices package.
package main

import (
	"fmt"
	"slices"
)

// Equal reports whether two slices are equal by comparing the length of
// each slice and testing that all elements are equal.

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
	// Equal

	fmt.Println("list1 == list2", slices.Equal(list1, list2))
	fmt.Println("list3 == list4", slices.Equal(list3, list4))

	// -------------------------------------------------------------------------
	// EqualFunc

	fmt.Println("list1 == list2", slices.EqualFunc(list1, list2, compare))
	fmt.Println("list3 == list4", slices.EqualFunc(list3, list4, compare))
}

func compare(a, b int) bool {
	return a == b
}

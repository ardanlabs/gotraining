// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program shows how to use the Sort APIs from the slices package.
package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

// Sort sorts a slice of any ordered type in ascending order.
// IsSorted reports whether x is sorted in ascending order.

func main() {

	// -------------------------------------------------------------------------
	// Sort

	list := []int{1, 4, 5, 2, 8, 3, 6, 9, 7}
	fmt.Printf("List: Addr(%x), %v\n", &list[0], list)

	slices.Sort(list)
	fmt.Printf("List: Addr(%x), %v\n", &list[0], list)

	is := slices.IsSorted(list)
	fmt.Println("Is list sorted:", is)

	// -------------------------------------------------------------------------
	// SortFunc

	list = []int{1, 4, 5, 2, 8, 3, 6, 9, 7}
	fmt.Printf("List: Addr(%x), %v\n", &list[0], list)

	slices.SortFunc(list, compare)
	fmt.Printf("List: Addr(%x), %v\n", &list[0], list)

	is = slices.IsSortedFunc(list, compare)
	fmt.Println("Is list sorted:", is)

	// -------------------------------------------------------------------------
	// SortStableFunc

	list = []int{1, 4, 5, 2, 8, 3, 6, 9, 7}
	fmt.Printf("List: Addr(%x), %v\n", &list[0], list)

	slices.SortStableFunc(list, compare)
	fmt.Printf("List: Addr(%x), %v\n", &list[0], list)

	is = slices.IsSortedFunc(list, compare)
	fmt.Println("Is list sorted:", is)
}

func compare(a, b int) bool {
	return a < b
}

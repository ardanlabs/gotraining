// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program shows how to use the Index API from the slices package.
package main

import (
	"fmt"
	"slices"
)

// Index returns the index of the first occurrence or -1 if not present.

func main() {
	list := []int{1, 1, 2, 2, 1, 1, 3, 3, 4, 5}
	fmt.Println("Slice", list)

	// -------------------------------------------------------------------------
	// Index

	fmt.Printf("Looking for 5, idx[%d]\n", slices.Index(list, 5))
	fmt.Printf("Looking for 0, idx[%d]\n", slices.Index(list, 0))
	fmt.Printf("Looking for 2, idx[%d]\n", slices.Index(list, 2))

	// -------------------------------------------------------------------------
	// IndexFunc

	fmt.Printf("Looking for 5, idx[%d]\n", slices.IndexFunc(list, compare(5)))
	fmt.Printf("Looking for 0, idx[%d]\n", slices.IndexFunc(list, compare(0)))
	fmt.Printf("Looking for 2, idx[%d]\n", slices.IndexFunc(list, compare(2)))
}

func compare(a int) func(int) bool {
	return func(b int) bool {
		return a == b
	}
}

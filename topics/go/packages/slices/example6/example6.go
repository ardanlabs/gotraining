// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program shows how to use the Contains API from the slices package.
package main

import (
	"fmt"
	"slices"
)

// Contains reports whether a specified value is present in slice.

func main() {
	list := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Slice", list)

	// -------------------------------------------------------------------------
	// Contains

	hasZero := slices.Contains(list, 0)
	fmt.Println("Does the list contain 0:", hasZero)

	hasFour := slices.Contains(list, 4)
	fmt.Println("Does the list contain 4:", hasFour)

	// -------------------------------------------------------------------------
	// ContainsFunc

	hasZero = slices.ContainsFunc(list, compare(0))
	fmt.Println("Does the list contain 0:", hasZero)

	hasFour = slices.ContainsFunc(list, compare(4))
	fmt.Println("Does the list contain 4:", hasFour)
}

func compare(a int) func(int) bool {
	return func(b int) bool {
		return a == b
	}
}

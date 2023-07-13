// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program shows how to use the Compact APIs from the slices package.
package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

// Compact replaces consecutive runs of equal elements with a single copy.
// Compact modifies the contents of the slice and does not create a new slice.

func main() {

	// -------------------------------------------------------------------------
	// Compact

	list := []int{1, 1, 2, 2, 1, 1, 3, 3, 4, 5}
	fmt.Printf("List: Addr(%x), %v\n", &list[0], list)

	compact := slices.Compact(list)
	fmt.Printf("List: Addr(%x), %v\n", &compact[0], compact)

	// -------------------------------------------------------------------------
	// CompactFunc

	list = []int{1, 1, 2, 2, 1, 1, 3, 3, 4, 5}
	fmt.Printf("List: Addr(%x), %v\n", &list[0], list)

	compact = slices.CompactFunc(list, compare)
	fmt.Printf("List: Addr(%x), %v\n", &compact[0], compact)
}

// compare needs to return true if the two values are the same.
func compare(a int, b int) bool {
	return a == b
}

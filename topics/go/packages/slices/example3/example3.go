// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program shows how to use the Clone API from the slices package.
package main

import (
	"fmt"
	"slices"
)

// Clone creates a new slice value and underlying array with a shallow
// copy of the elements.

func main() {
	list := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("List: Addr(%x), %v\n", &list[0], list)

	// -------------------------------------------------------------------------
	// Clone

	list = slices.Clone(list)
	fmt.Printf("Copy: Addr(%x), %v\n", &list[0], list)
}

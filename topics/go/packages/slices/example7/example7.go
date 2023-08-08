// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program shows how to use the Delete API from the slices package.
package main

import (
	"fmt"
	"slices"
)

// Delete removes a contiguous set of elements from the slice.
// Delete modifies the contents of the slice and does not create a new slice.

func main() {
	list := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("List: Addr(%x), %v\n", &list[0], list)

	// -------------------------------------------------------------------------
	// Delete - Remove 2 from the list

	list = slices.Delete(list, 1, 2)
	fmt.Printf("List: Addr(%x), %v\n", &list[0], list)

	// -------------------------------------------------------------------------
	// Delete - Remove 4 and 5 from the list

	list = slices.Delete(list, 2, 4)
	fmt.Printf("List: Addr(%x), %v\n", &list[0], list)
}

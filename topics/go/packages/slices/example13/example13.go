// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program shows how to use the Replace API from the slices package.
package main

import (
	"fmt"
	"slices"
)

// Replace replaces the elements specified for a given range and returns
// the modified slice.

func main() {
	list := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("List: Addr(%x), %v\n", &list[0], list)

	// -------------------------------------------------------------------------
	// Replace - Change 3 to 7

	list = slices.Replace(list, 2, 3, 7)
	fmt.Printf("List: Addr(%x), %v\n", &list[0], list)

	// -------------------------------------------------------------------------
	// Replace - Change 4, 5, 6 to 8, 9

	list = slices.Replace(list, 3, 6, 8, 9)
	fmt.Printf("List: Addr(%x), %v\n", &list[0], list)
}

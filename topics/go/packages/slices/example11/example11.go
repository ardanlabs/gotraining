// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program shows how to use the Insert API from the slices package.
package main

import (
	"fmt"
	"slices"
)

// Insert adds values into the slices at the specified index and
// returns the modified slice.

func main() {
	list := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("List: Addr(%x), %v\n", &list[0], list)

	// -------------------------------------------------------------------------
	// Insert - add 7, 8, 9 to the end of the list

	list = slices.Insert(list, 6, 7, 8, 9)
	fmt.Printf("List: Addr(%x), %v\n", &list[0], list)

	// -------------------------------------------------------------------------
	// Insert - add 0 to the beginning of the list

	list = slices.Insert(list, 0, 0)
	fmt.Printf("List: Addr(%x), %v\n", &list[0], list)
}

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program shows how to use the Clip API from the slices package.
package main

import (
	"fmt"
	"slices"
)

// Clip removes the unused capacity from the slice.

func main() {
	list := make([]string, 0, 10)
	fmt.Printf("Addr(%x), Len(%d), Cap(%d)\n", &list, len(list), cap(list))

	// -------------------------------------------------------------------------
	// Append a string to the slice

	list = append(list, "A")
	fmt.Printf("Addr(%x), Len(%d), Cap(%d)\n", &list[0], len(list), cap(list))

	// -------------------------------------------------------------------------
	// Clip

	list = slices.Clip(list)
	fmt.Printf("Addr(%x), Len(%d), Cap(%d)\n", &list[0], len(list), cap(list))
}

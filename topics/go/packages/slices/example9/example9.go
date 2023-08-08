// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program shows how to use the Grow API from the slices package.
package main

import (
	"fmt"
	"slices"
)

// Grow increases the slice's capacity to guarantee space for another n elements.

func main() {
	list := make([]string, 0, 1)

	fmt.Printf("New Slice  : Len(%d), Cap(%d)\n",
		len(list), cap(list))

	// -------------------------------------------------------------------------
	// Append a string to the slice

	list = append(list, "A")

	fmt.Printf("Append 'A' : Addr(%x), Len(%d), Cap(%d)\n",
		&list[0], len(list), cap(list))

	// -------------------------------------------------------------------------
	// Grow the capacity by 1 element

	list = slices.Grow(list, 1)

	fmt.Printf("Grow By 1  : Addr(%x), Len(%d), Cap(%d)\n",
		&list[0], len(list), cap(list))

	// -------------------------------------------------------------------------
	// Append a string to the slice

	list = append(list, "B")

	fmt.Printf("Append 'B' : Addr(%x), Len(%d), Cap(%d)\n",
		&list[0], len(list), cap(list))

	// -------------------------------------------------------------------------
	// Grow the capacity by 10 elements

	list = slices.Grow(list, 10)

	fmt.Printf("Grow By 10 : Addr(%x), Len(%d), Cap(%d)\n",
		&list[0], len(list), cap(list))
}

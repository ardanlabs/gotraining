// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This program showcases
// the `slices` package's compact func function.
// The aim of this test is to remove consecutive duplicate
// elements within a slice
// This program requires Go 1.21rc1
package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {

	a := []int{
		1, 1, 1, 2, 2, 2, 1, 1, 3, 3, 4, 3, 5,
	}

	fmt.Println("Original", a, len(a))

	comp := func(a, b int) bool {
		// preserve subsequent 1's
		if a == 1 && b == 1 {
			return false
		}

		return a == b
	}

	a = slices.CompactFunc(a, comp)
	fmt.Println("Compacted", a, len(a))
}

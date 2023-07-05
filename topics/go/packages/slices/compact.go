// This program showcases
// the `slices` package's compact function.
// The aim of this test is to remove consecutive duplicate
// elements within a slice
// This program requires Go 1.21rc1
package main

import (
	"fmt"
	"slices"
)

func main() {

	a := []int{
		1, 1, 2, 2, 1, 1, 3, 3, 4, 5,
	}

	fmt.Println("Original", a)

	a = slices.Compact(a)
	fmt.Println("Compacted", a)
}

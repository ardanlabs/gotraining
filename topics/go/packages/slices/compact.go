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
		1, 2, 1, 1, 3, 4, 5,
	}

	fmt.Println("Original", a)
	cache := map[int]bool{}
	a = slices.CompactFunc(a, func(e1 int, e2 int) bool {

		_, ok := cache[e1]

		if ok {
			return true
		}

		cache[e1] = true

		return false
	})
	fmt.Println("Compact", a)
}

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This program showcases
// the `slices` package's IsSorted function
// to determine if an array is in ascending
// order.
// This program requires Go 1.21rc1
package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {

	a := []int{
		1, 2, 3, 4, 5, 6,
	}

	b := []int{
		1, 7, 2, 3, 4, 5, 6,
	}

	fmt.Println("Array A:", a)
	fmt.Println("Array B:", b)

	isSortedA := slices.IsSorted(a)
	isSortedB := slices.IsSorted(b)

	fmt.Println("Is Array A in ascending order:", isSortedA)
	fmt.Println("Is Array B in ascending order:", isSortedB)
}

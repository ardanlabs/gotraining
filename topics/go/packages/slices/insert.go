// This program showcases
// the `slices` package's insert function.
// The aim of this test is determine
// the behavior the package's insert function
// and how it modifies an array.
// This program requires Go 1.21rc1
package main

import (
	"fmt"
	"slices"
)

func main() {

	a := []int{
		1, 4, 5,
	}

	b := []int{
		1, 2, 3,
	}

	// Adding 2,3 at index 1 in array a
	a = slices.Insert(a, 1, 2, 3)

	fmt.Println("Array a after insert", a)

	// Adding 4,5 at index 3 in array b
	b = slices.Insert(b, 3, 4, 5)

	fmt.Println("Array b after insert", b)
}

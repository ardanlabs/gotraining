// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This program showcases
// the `slices` package's insert function.
// The aim of this test is determine
// the behavior the package's insert function
// and how it modifies a slice.
// This program requires Go 1.21rc1
package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {

	a := []int{
		1, 4, 5,
	}

	b := []int{
		1, 2, 3,
	}

	fmt.Println("Slice a", a)
	fmt.Println("Slice b", b)

	// Adding 2,3 at index 1 in array a
	a = slices.Insert(a, 1, 2, 3)

	fmt.Println("Array a after insert", a)

	// Adding 4,5 at index 3 in array b
	b = slices.Insert(b, 3, 4, 5)

	fmt.Println("Array b after insert", b)
}

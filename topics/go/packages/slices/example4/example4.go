// This program showcases
// the `slices` package's clone function.
// The aim of this test is to understand
// the effects of cloning a slice
// This program requires Go 1.21rc1
package main

import (
	"fmt"
	"slices"
)

func main() {

	a := []int{
		1, 2, 3, 4, 5,
	}

	c := slices.Clone(a)

	fmt.Println(
		"Array a",
		a,
	)

	fmt.Println(
		"Array c",
		c,
	)

	// An example with pointers to demonstrate
	// how the clone function outputs a shallow copy.
	val1 := 2
	val2 := 1

	pA := []*int{
		&val1,
		&val2,
	}

	pC := slices.Clone(pA)
	*pC[0]++
	*pC[1] = 20

	fmt.Println("Value of element at index 0 of Array pA", *pA[0])
	fmt.Println("Value of element at index 1 of Array pA", *pA[1])
}

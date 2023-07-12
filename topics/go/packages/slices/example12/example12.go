// This program showcases
// the `slices` package's equal function.
// The aim of this test is to determine
// if two slices are equal.
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

	b := []int{
		1, 2, 3, 4, 5,
	}

	c := []int{
		6, 6, 6, 6, 6,
	}

	fmt.Println("Slice a", a)
	fmt.Println("Slice b", b)
	fmt.Println("Slice c", c)

	fmt.Println(
		"Is slice a and b equal",
		slices.Equal(a, b),
	)

	fmt.Println(
		"Is slice b and c equal",
		slices.Equal(a, c),
	)
}

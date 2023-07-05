// This program showcases
// the `slices` package's compact function.
// The aim of this test is to define a sort function
// and use it to sort a slice of integers in ascending
// order.
// This program requires Go 1.21rc1
package main

import (
	"fmt"
	"slices"
)

func main() {

	a := []int{
		5, 4, 1, 1, 6, 3,
	}

	fmt.Println("Original", a)

	a = slices.SortFunc(a, NumSort)

	fmt.Println("Sorted", a)
}

type NumSort func(p1, p2 int) bool

func (n NumSort) Sort(numbers []int) {
	fmt.Println(numbers)
}

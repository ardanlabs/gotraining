// This program showcases
// the `slices` package's SortFunc function.
// The aim of this test is to define a sort function
// and use it to sort a slice of integers in ascending
// order and another to sort  a complex type.
// This program requires Go 1.21rc1
package main

import (
	"fmt"
	"slices"
	"time"
)

type Order struct {
	Date     time.Time
	Name     string
	Complete bool
}

func main() {

	// This example demonstrates
	// how sort func can be used
	// with a slice of numbers.
	ints := []int{
		2, 4, 6, 7, 8, 9, 1, 0,
	}

	var less = func(a, b int) int { return a - b }
	slices.SortFunc(ints, less)

	fmt.Println("Is slice int sorted", slices.IsSorted(ints))
	fmt.Println("Resulting slice after sort", ints)

	// This example showcases how the
	// sort function can be leveraged
	// to sort complex types.
	orders := []Order{
		Order{
			Date:     time.Now().AddDate(0, 0, 2),
			Name:     "Bob",
			Complete: false,
		},
		Order{
			Date:     time.Now(),
			Name:     "Alice",
			Complete: true,
		},
	}

	// code is unsafe, in the case
	// of the date difference exceeding the max `int` size
	slices.SortFunc(orders, func(a, b Order) int { return int(a.Date.Unix() - b.Date.Unix()) })

	fmt.Println("Sorted slices of orders:", orders)
}

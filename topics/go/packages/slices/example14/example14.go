// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This program showcases
// the `slices` package's grow function.
// The aim of this test is grow the size (cap) of
// a slice.
// This program requires Go 1.21rc1
package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {

	// Here an array with cap. 10
	// is defined.
	a := make([]string, 0, 5)

	// Two elements are appended to the
	// array.
	a = append(a, "Hello", "World")

	// THe output of the cap here will be
	// 5.
	fmt.Println("Original", a, cap(a))

	// Grow will increase the cap of the array.
	// The function will panic if the runtime
	// is unable to increase capacity. A great
	// pre-emptive measure prior to increasing a
	// slices capacity.
	// if the supplied size is 1 then
	// no change is observed
	a = slices.Grow(a, 1)

	// After growing the cap will be
	// 5.
	fmt.Println("Grow 1", a, cap(a))

	// Growing with value 8 will change the cap to 10 due
	// to the pre-existing elemnts in the slice.
	a = slices.Grow(a, 8)

	// After growing the cap will be reduced
	// to 10.
	fmt.Println("Grow 8", a, cap(a))
}

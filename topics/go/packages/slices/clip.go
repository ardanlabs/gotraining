// This program showcases
// the `slices` package's clip function.
// This program requires Go 1.21rc1
package main

import (
	"fmt"
	"slices"
)

func main() {

	// Here an array with cap. 10
	// is defined.
	a := make([]string, 0, 10)

	// Two elements are appended to the
	// array.
	a = append(a, "Hello", "World")

	// THe output of the cap here will be
	// 10.
	fmt.Println("Original", a, cap(a))
	a = slices.Clip(a)

	// After clipping the cap will be reduced
	// to 2.
	fmt.Println("Clipped", a, cap(a))
}

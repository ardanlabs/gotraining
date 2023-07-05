// This program showcases
// the `slices` package's clip function.
// This program requires Go 1.21rc1
package main

import (
	"fmt"
	"slices"
)

func main() {

	a := make([]string, 0, 10)
	a = append(a, "Hello", "World")

	fmt.Println("Original", a, cap(a))
	a = slices.Clip(a)
	fmt.Println("Clipped", a, cap(a))
}

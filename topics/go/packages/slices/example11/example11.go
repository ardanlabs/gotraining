// This program showcases
// the `slices` package's delete function
// to remove an element from an array.
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
	a = slices.Delete(a, 1, 2)
	fmt.Println("Modified", a, cap(a))
}

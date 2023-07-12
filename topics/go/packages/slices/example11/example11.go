// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This program showcases
// the `slices` package's delete function
// to remove an element from an array.
// This program requires Go 1.21rc1
package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {

	a := make([]string, 0, 10)
	a = append(a, "Hello", "World")

	fmt.Println("Original", a, cap(a))
	a = slices.Delete(a, 1, 2)
	fmt.Println("Modified", a, cap(a))
}

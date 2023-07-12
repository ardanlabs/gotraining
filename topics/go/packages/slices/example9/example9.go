// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This program showcases
// the `slices` package's contain function.
// The aim of this test is to determine
// if a slice contains an element
// This program requires Go 1.21rc1
package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {

	a := []int{
		1, 2, 3, 4, 5,
	}

	fmt.Println("Array", a)

	containSix := slices.Contains(a, 6)
	containTwo := slices.Contains(a, 2)

	fmt.Println("Does the array contain 6:", containSix)
	fmt.Println("Does the array contain 2:", containTwo)
}

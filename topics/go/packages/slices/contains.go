// This program showcases
// the `slices` package's compact function.
// The aim of this test is to determine
// if a slice contains an element
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

	fmt.Println("Array", a)

	containSix := slices.Contains(a, 6)
	containTwo := slices.Contains(a, 2)

	fmt.Println("Does the array contain 6:", containSix)
	fmt.Println("Does the array contain 2:", containTwo)
}

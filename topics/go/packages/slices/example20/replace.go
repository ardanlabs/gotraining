// This program showcases
// the `slices` package's replace function.
// The aim of this test is determine
// the behavior the package's replace function
// and how it modifies a slice.
// The examples here will be used
// This program requires Go 1.21rc1
package main

import (
	"fmt"
	"slices"
)

func main() {

	a := []string{
		"How", "you", "?",
	}

	b := []string{
		"Good", "morning", "Gopher",
	}

	fmt.Println("Slice a", a)
	fmt.Println("Slice b", b)

	// The following op will
	// add 'are' to the sentence fragments
	// to do this, I'll replace the slice at [0:1]
	// with `How` and `are`
	a = slices.Replace(a, 0, 1, "How", "are")

	fmt.Println("Array a after replace", a)

	// This operation wil replace the word
	// Gopher with Rustacean.
	b = slices.Replace(b, 2, 3, "Rustacean")

	fmt.Println("Array b after replace", b)
}

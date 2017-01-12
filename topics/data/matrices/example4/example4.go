// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example4

// Sample program to illustrate various ways to format matrix output.
package main

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)

func main() {

	// Create a small matrix.
	a := mat64.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})

	// Create a matrix formatting value with a prefix.
	fa := mat64.Formatted(a, mat64.Prefix("    "))

	// Print the matrix with and without zero value elements.
	fmt.Printf("with all values:\na = %v\n\n", fa)
	fmt.Printf("with only non-zero values:\na = % v\n\n", fa)

	// Modify the matrix.
	a.Set(0, 2, 0)

	// Print it again without zero value elements.
	fmt.Printf("after modification with only non-zero values:\na = % v\n\n", fa)
}

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example3

// Sample program to access values within a matrix.
package main

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)

func main() {

	// Create a small matrix.
	a := mat64.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})

	// Get a single value from the matrix.
	val := a.At(0, 1)
	fmt.Printf("The value of a at (0,1) is: %.2f\n\n", val)

	// Get the values in a specific column.
	col := mat64.Col(nil, 2, a)
	fmt.Printf("The values in the 3rd column are: %v\n\n", col)

	// Get the values in a specific row.
	row := mat64.Row(nil, 2, a)
	fmt.Printf("The values in the 3rd row are: %v\n\n", row)

	// Get a "view" of a portion of the matrix extending from a starting
	// point (the first two arguments) a certain number of rows and columns
	// (the last two arguments).
	b := a.View(0, 0, 2, 2)

	// Print it again without zero value elements.
	fb := mat64.Formatted(b, mat64.Prefix("    "))
	fmt.Printf("The \"view\" of a looks like:\nb = % v\n\n", fb)
}

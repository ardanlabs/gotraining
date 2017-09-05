// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example2

// Sample program to show modifications to matrices.
package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {

	// Create a small matrix.
	a := mat.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})

	// Modify a single element.
	a.Set(0, 2, 0)

	// Modify an entire row.
	a.SetRow(0, []float64{3.0, 2.0, 1.0})

	// Modify an entire column.
	a.SetCol(0, []float64{1.0, 3.0, 2.0})

	// Print it again without zero value elements.
	fa := mat.Formatted(a, mat.Prefix("    "))
	fmt.Printf("after modification:\na = % v\n\n", fa)
}

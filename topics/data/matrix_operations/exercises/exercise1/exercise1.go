// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./exercise1

// Sample program to divide a matrix by its norm.
package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {

	// Create an example matrix.
	m := mat.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})

	// Get the Euclidean norm of the matrix.
	matNorm := mat.Norm(m, 2)

	// Multiply the matrix by 1 / matNorm.
	matScaled := mat.NewDense(0, 0, nil)
	matScaled.Scale(1/matNorm, m)

	// Output the matrix to standard out.
	ft := mat.Formatted(matScaled, mat.Prefix("         "))
	fmt.Printf("a/norm = %v\n\n", ft)
}

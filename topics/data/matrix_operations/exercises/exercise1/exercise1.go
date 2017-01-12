// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./exercise1

// Sample program to divide a matrix by its norm.
package main

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)

func main() {

	// Create an example matrix.
	mat := mat64.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})

	// Get the Euclidean norm of the matrix.
	matNorm := mat64.Norm(mat, 2)

	// Multiply the matrix by 1 / matNorm.
	matScaled := mat64.NewDense(0, 0, nil)
	matScaled.Scale(1/matNorm, mat)

	// Output the matrix to standard out.
	ft := mat64.Formatted(matScaled, mat64.Prefix("         "))
	fmt.Printf("a/norm = %v\n\n", ft)
}

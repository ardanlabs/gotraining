// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example4

// Sample program to compute vector and matrix norms.
package main

import (
	"fmt"

	"github.com/gonum/floats"
	"github.com/gonum/matrix/mat64"
)

func main() {

	// Create a "vector", really just a float64 slice.
	vec := []float64{1.2, 2.3, 4.1, 5.8}

	// Get the Euclidean norm of the vector.
	vecNorm := floats.Norm(vec, 2.0)
	fmt.Printf("\nThe Vector norm of vec is: %0.2f\n\n", vecNorm)

	// Create an example matrix.
	mat := mat64.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})

	// Get the Euclidean norm of the matrix.
	matNorm := mat64.Norm(mat, 2)
	fmt.Printf("The Matrix norm of mat is: %0.2f\n\n", matNorm)

}

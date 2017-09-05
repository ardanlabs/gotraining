// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example3

// Sample program to solve an eigenvalue/vector problem.
package main

import (
	"fmt"
	"log"

	"gonum.org/v1/gonum/mat"
)

func main() {

	// Create two matrices of the same size, a and b.
	a := mat.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})

	// Solve the eigenvalue problem.
	var eig mat.Eigen
	if ok := eig.Factorize(a, false, true); !ok {
		log.Fatal("Could not factorize the EigenSym value.")
	}

	// Output the eigenvalues.
	fmt.Printf("\neignvalues = %v\n\n", eig.Values(nil))

	// Output the eigenvectors.
	vectors := eig.Vectors()
	fv := mat.Formatted(vectors, mat.Prefix("               "))
	fmt.Printf("eigenvectors = %v\n\n", fv)
}

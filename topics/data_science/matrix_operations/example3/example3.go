// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example3

// Sample program to solve an eigenvalue/vector problem.
package main

import (
	"fmt"
	"log"

	"github.com/gonum/matrix/mat64"
)

func main() {

	// Create two matrices of the same size, a and b.
	a := mat64.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})

	// Solve the eigenvalue problem.
	var eig mat64.Eigen
	if ok := eig.Factorize(a, true); !ok {
		log.Fatal("Could not factorize the EigenSym value.")
	}

	// Output the eigenvalues.
	fmt.Printf("\neignvalues = %v\n\n", eig.Values(nil))

	// Output the eigenvectors.
	vectors := eig.Vectors()
	fv := mat64.Formatted(vectors, mat64.Prefix("               "))
	fmt.Printf("eigenvectors = %v\n\n", fv)
}

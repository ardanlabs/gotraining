// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./exercise1

// Sample program to project iris data on to 3 principal components.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gonum/matrix/mat64"
	"github.com/gonum/stat"
	"github.com/kniren/gota/dataframe"
)

func main() {

	// Open the iris dataset file.
	irisFile, err := os.Open("../../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Parse the file into a Gota dataframe.
	irisDF := dataframe.ReadCSV(irisFile)

	// Form a matrix from the dataframe.
	mat := irisDF.Select([]string{"sepal_length", "sepal_width", "petal_length", "petal_width"}).Matrix()

	// Calculate the principal component direction vectors
	// and variances.
	vecs, _, ok := stat.PrincipalComponents(mat, nil)
	if !ok {
		log.Fatal("Could not calculate prinicple components")
	}

	// Project the data onto the first 3 principal components.
	var proj mat64.Dense
	proj.Mul(mat, vecs.View(0, 0, 4, 3))

	// Output the resulting projected features to standard out.
	fmt.Printf("proj = %.4f\n", mat64.Formatted(&proj, mat64.Prefix("       ")))
}

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example1

// Sample program to illustrate the calculation of principal components.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gonum/stat"
	"github.com/kniren/gota/dataframe"
)

func main() {

	// Open the iris dataset file.
	irisFile, err := os.Open("../data/iris.csv")
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
	_, vars, ok := stat.PrincipalComponents(mat, nil)
	if !ok {
		log.Fatal("Could not calculate principal components")
	}

	// Output the principal component direction variances to standard out.
	fmt.Printf("\nvariances = %.4f\n\n", vars)
}

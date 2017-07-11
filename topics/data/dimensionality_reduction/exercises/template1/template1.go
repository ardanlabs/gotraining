// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./template1

// Sample program to project iris data on to 3 principal components.
package main

import (
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {

	// Open the iris dataset file.
	f, err := os.Open("../../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Parse the file into a Gota dataframe.
	irisDF := dataframe.ReadCSV(f)

	// Form a matrix from the dataframe.
	mat := irisDF.Select([]string{"sepal_length", "sepal_width", "petal_length", "petal_width"}).Matrix()

	// Calculate the principal component direction vectors
	// and variances.

	// Project the data onto the first 3 principal components.

	// Output the resulting projected features to standard out.
}

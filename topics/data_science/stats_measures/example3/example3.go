// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example3

// Sample program to calculate standard deviation and variance.
package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gonum/stat"
	"github.com/kniren/gota/data-frame"
)

func main() {

	// Pull in the CSV data.
	irisData, err := ioutil.ReadFile("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Create a dataframe from the CSV string.
	// The types of the columns will be inferred.
	irisDF := df.ReadCSV(string(irisData))

	// Get the float values from the "sepal_length" column as
	// we will be looking at the measures for this variable.
	sepalLength, err := irisDF.Col("sepal_length").Float()
	if err != nil {
		log.Fatal(err)
	}

	// Calculate the variance of the variable.
	varianceVal := stat.Variance(sepalLength, nil)

	// Calculate the standard deviation of the variable.
	stdDevVal := stat.StdDev(sepalLength, nil)

	// Output the results to standard out.
	fmt.Printf("\nSepal Length Summary Statistics:\n")
	fmt.Printf("Variance value: %0.2f\n", varianceVal)
	fmt.Printf("Std Dev value: %0.2f\n\n", stdDevVal)
}

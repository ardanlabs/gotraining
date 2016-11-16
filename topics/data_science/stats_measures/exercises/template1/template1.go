// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./template1

// Sample program to calculate both central tendency and statistical dispersion
// measures for the iris dataset.
package main

import (
	"io/ioutil"
	"log"

	"github.com/kniren/gota/data-frame"
)

func main() {

	// Pull in the CSV data.
	irisData, err := ioutil.ReadFile("../../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Create a dataframe from the CSV string.
	// The types of the columns will be inferred.
	irisDF := df.ReadCSV(string(irisData))

	// Loop over the float columns.

	// Only look at the numeric columns.

	// Get the float values from the column.

	// Calculate the Mean of the variable.

	// Calculate the Mode of the variable.

	// Calculate the Median of the variable.

	// Calculate the variance of the variable.

	// Calculate the standard deviation of the variable.

	// Output the results to standard out.
}

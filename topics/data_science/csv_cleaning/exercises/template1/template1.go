// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./template1

// Sample program to read in a CSV, create three filtered datasets, and
// save those datasets to three separate files.
package main

import (
	"io/ioutil"
	"log"
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

	// Filter the dataset into three separate dataframes,
	// each corresponding to one of the Iris species.

	// Save each of the species dataframe to a file.
}

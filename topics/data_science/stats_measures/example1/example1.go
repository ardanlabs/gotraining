// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example1

// Sample program to calculate means, modes, and medians.
package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gonum/stat"
	"github.com/kniren/gota/data-frame"
	"github.com/montanaflynn/stats"
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

	// Calculate the Mean of the variable.
	meanVal := stat.Mean(sepalLength, nil)

	// Calculate the Mode of the variable.
	modeVal, modeCount := stat.Mode(sepalLength, nil)

	// Calculate the Median of the variable.
	medianVal, err := stats.Median(sepalLength)
	if err != nil {
		log.Fatal(err)
	}

	// Output the results to standard out.
	fmt.Printf("\nSepal Length Summary Statistics:\n")
	fmt.Printf("Mean value: %0.2f\n", meanVal)
	fmt.Printf("Mode value: %0.2f\n", modeVal)
	fmt.Printf("Mode count: %d\n", int(modeCount))
	fmt.Printf("Media value: %0.2f\n\n", medianVal)
}

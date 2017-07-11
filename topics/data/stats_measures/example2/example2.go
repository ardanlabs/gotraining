// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example2

// Sample program to calculate means, modes, and medians.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gonum/floats"
	"github.com/kniren/gota/dataframe"
)

func main() {

	// Open the iris dataset file.
	f, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	irisDF := dataframe.ReadCSV(f)

	// Get the float values from the "sepal_length" column as
	// we will be looking at the measures for this variable.
	sepalLength := irisDF.Col("sepal_length").Float()

	// Calculate the Max of the variable.
	minVal := floats.Min(sepalLength)

	// Calculate the Max of the variable.
	maxVal := floats.Max(sepalLength)

	// Calculate the Median of the variable.
	rangeVal := maxVal - minVal

	// Output the results to standard out.
	fmt.Printf("\nSepal Length Summary Statistics:\n")
	fmt.Printf("Max value: %0.2f\n", maxVal)
	fmt.Printf("Min value: %0.2f\n", minVal)
	fmt.Printf("Range value: %0.2f\n\n", rangeVal)
}

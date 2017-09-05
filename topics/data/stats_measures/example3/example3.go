// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example3

// Sample program to calculate standard deviation and variance.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/gonum/stat"
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

	// Calculate the variance of the variable.
	varianceVal := stat.Variance(sepalLength, nil)

	// Calculate the standard deviation of the variable.
	stdDevVal := stat.StdDev(sepalLength, nil)

	// Output the results to standard out.
	fmt.Printf("\nSepal Length Summary Statistics:\n")
	fmt.Printf("Variance value: %0.2f\n", varianceVal)
	fmt.Printf("Std Dev value: %0.2f\n\n", stdDevVal)
}

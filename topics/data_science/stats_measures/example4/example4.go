// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example4

// Sample program to calculate quantiles
package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gonum/floats"
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
	floatCol, ok := irisDF.Col("sepal_length").Elements.(df.FloatElements)
	if !ok {
		log.Fatal(fmt.Errorf("Could not parse float column"))
	}

	// Convert the Gota float values to a normal slice of floats.
	var sepalLength []float64
	for _, val := range floatCol {
		sepalLength = append(sepalLength, *val.Float())
	}

	// Sort the values.
	inds := make([]int, len(sepalLength))
	floats.Argsort(sepalLength, inds)

	// Get the Quantiles.
	quant25 := stat.Quantile(0.25, stat.Empirical, sepalLength, nil)

	// Output the results to standard out.
	fmt.Printf("\nSepal Length Summary Statistics:\n")
	fmt.Printf("25 Quantile: %0.2f\n", quant25)
	//fmt.Printf("Std Dev value: %0.2f\n\n", stdDevVal)
}

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example4

// Sample program to calculate quantiles
package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/gonum/floats"
	"github.com/gonum/stat"
	"github.com/kniren/gota/dataframe"
	"github.com/pachyderm/pachyderm/src/client"
)

func main() {

	// Connect to Pachyderm on our localhost.  By default
	// Pachyderm will be exposed on port 30650.
	c, err := client.NewFromAddress("0.0.0.0:30650")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Get the Iris dataset from Pachyderm's data
	// versioning at the latest commit.
	var b bytes.Buffer
	if err := c.GetFile("iris", "master", "iris.csv", 0, 0, "", false, nil, &b); err != nil {
		log.Fatal()
	}

	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	irisDF := dataframe.ReadCSV(bytes.NewReader(b.Bytes()))

	// Get the float values from the "sepal_length" column as
	// we will be looking at the measures for this variable.
	sepalLength := irisDF.Col("sepal_length").Float()

	// Sort the values.
	inds := make([]int, len(sepalLength))
	floats.Argsort(sepalLength, inds)

	// Get the Quantiles.
	quant25 := stat.Quantile(0.25, stat.Empirical, sepalLength, nil)
	quant50 := stat.Quantile(0.50, stat.Empirical, sepalLength, nil)
	quant75 := stat.Quantile(0.75, stat.Empirical, sepalLength, nil)

	// Output the results to standard out.
	fmt.Printf("\nSepal Length Summary Statistics:\n")
	fmt.Printf("25 Quantile: %0.2f\n", quant25)
	fmt.Printf("50 Quantile: %0.2f\n", quant50)
	fmt.Printf("75 Quantile: %0.2f\n\n", quant75)
}

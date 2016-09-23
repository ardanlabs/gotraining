// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example1

// Sample program to read in records from an example CSV file and form
// a matrix with gonum.
package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gonum/matrix/mat64"
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

	// Sequentially move the columns into a slice of floats.
	floatData := make([]float64, 4*irisDF.Nrow())
	var dataIndex int
	for colIndex, colName := range irisDF.Names() {

		// If the column is one of the float columns, move it
		// into the slice of floats.
		if colIndex < 4 {

			// Extract the columns as a slice of floats.
			floatCol, ok := irisDF.Col(colName).Elements.(df.FloatElements)
			if !ok {
				log.Fatal(fmt.Errorf("Could not parse float column."))
			}

			// Append the float values to floatData.
			for _, floatVal := range floatCol {
				floatData[dataIndex] = *floatVal.Float()
				dataIndex++
			}
		}
	}

	// Form the matrix.
	mat := mat64.NewDense(irisDF.Nrow(), 4, floatData)

	// As a sanity check, output the matrix to standard out.
	fMat := mat64.Formatted(mat, mat64.Prefix("      "))
	fmt.Printf("mat = %v\n\n", fMat)
}

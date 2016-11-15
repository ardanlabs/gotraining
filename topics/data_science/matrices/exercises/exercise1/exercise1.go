// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./exercise1

// Sample program to read in records from a CSV file and form
// a matrix with gonum.
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gonum/matrix/mat64"
)

func main() {

	// Open the iris dataset file.
	csvFile, err := os.Open("../../data/diabetes.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = 11

	// Read in all of the CSV records
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Sequentially move the rows into a slice of floats.
	floatData := make([]float64, 11*len(rawCSVData))
	var dataIndex int
	for i, record := range rawCSVData {

		// Skip the header.
		if i == 0 {
			continue
		}

		// Loop over the columns.
		for _, rawVal := range record {

			// Convert the value to a float.
			val, err := strconv.ParseFloat(rawVal, 64)
			if err != nil {
				log.Fatal(fmt.Errorf("Could not parse float value"))
			}

			// Add the float value to the slice of floats.
			floatData[dataIndex] = val
			dataIndex++
		}
	}

	// Form the matrix.
	mat := mat64.NewDense(len(rawCSVData), 11, floatData)

	// Get the first 10 rows.
	firstTen := mat.View(0, 0, 10, 11)

	// As a sanity check, output the rows to standard out.
	fMat := mat64.Formatted(firstTen, mat64.Prefix("      "))
	fmt.Printf("mat = %v\n\n", fMat)
}

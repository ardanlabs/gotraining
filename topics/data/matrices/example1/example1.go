// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example1

// Sample program to read in records from an example CSV file and form
// a matrix with gonum.
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"gonum.org/v1/gonum/mat"
)

func main() {

	// Open the iris dataset file.
	f, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 5

	// Read in all of the CSV records
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// floatData will hold all the float values that will eventually be
	// used to form out matrix.
	floatData := make([]float64, 4*len(rawCSVData))

	// dataIndex will track the current index of the matrix values.
	var dataIndex int

	// Sequentially move the rows into a slice of floats.
	for idx, record := range rawCSVData {

		// Skip the header row.
		if idx == 0 {
			continue
		}

		// Loop over the float columns.
		for i := 0; i < 4; i++ {

			// Convert the value to a float.
			val, err := strconv.ParseFloat(record[i], 64)
			if err != nil {
				log.Fatal("Could not parse float value")
			}

			// Add the float value to the slice of floats.
			floatData[dataIndex] = val
			dataIndex++
		}
	}

	// Form the matrix.
	m := mat.NewDense(len(rawCSVData), 4, floatData)

	// As a sanity check, output the matrix to standard out.
	fMat := mat.Formatted(m, mat.Prefix("      "))
	fmt.Printf("mat = %v\n\n", fMat)
}

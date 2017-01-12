// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example1

// Sample program to read in records from an example CSV file and form
// a matrix with gonum.
package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"strconv"

	"github.com/gonum/matrix/mat64"
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

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(bytes.NewReader(b.Bytes()))
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
	mat := mat64.NewDense(len(rawCSVData), 4, floatData)

	// As a sanity check, output the matrix to standard out.
	fMat := mat64.Formatted(mat, mat64.Prefix("      "))
	fmt.Printf("mat = %v\n\n", fMat)
}

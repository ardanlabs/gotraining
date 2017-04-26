// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./template1

// Sample program to read in records from a CSV file and form
// a matrix with gonum.
package main

import (
	"bytes"
	"encoding/csv"
	"log"

	"github.com/dwhitena/pachyderm/src/client"
)

func main() {

	// Connect to Pachyderm on our localhost.  By default
	// Pachyderm will be exposed on port 30650.
	c, err := client.NewFromAddress("0.0.0.0:30650")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Get the diabetes dataset from Pachyderm's data
	// versioning at the latest commit.
	var b bytes.Buffer
	if err := c.GetFile("diabetes", "master", "diabetes.csv", 0, 0, &b); err != nil {
		log.Fatal()
	}

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(bytes.NewReader(b.Bytes()))
	reader.FieldsPerRecord = 11

	// Read in all of the CSV records
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Sequentially move the rows into a slice of floats.

	// Loop over the columns.

	// Add the float values to a slice of floats.

	// Form the matrix.

	// Get the first 10 rows.

	// As a sanity check, output the rows to standard out.
}

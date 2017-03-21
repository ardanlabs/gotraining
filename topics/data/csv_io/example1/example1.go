// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example1

// Sample program to read in records from an example CSV file.
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
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

	// Assume we don't know the number of fields per line.  By setting
	// FieldsPerRecord negative, each row may have a variable
	// number of fields.
	reader.FieldsPerRecord = -1

	// Read in all of the CSV records.
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// As a sanity check, display the records to stdout.
	for _, each := range rawCSVData {
		fmt.Println(each)
	}
}

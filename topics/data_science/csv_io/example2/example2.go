// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example2

// Sample program to read in records from an example CSV file,
// and catch an unexpected extra field in the data.
package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func main() {

	// Open the iris dataset file.
	f, err := os.Open("../data/iris_extra_field.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)

	// We should have 5 fields per line.  By setting
	// FieldsPerRecord to 5, we can validate that each of the
	// rows in our CSV has the correct number of fields.
	reader.FieldsPerRecord = 5

	// rawCSVData will hold our succesfully parsed rows.
	var rawCSVData [][]string

	// Read in the records looking for unexpected numbers of fields.
	for {

		// Read in a row. Check if we are at the end of the file.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// If we had a parsing error, log the error and move on.
		if err != nil {
			log.Println(err)
			continue
		}

		// Append the record to our data set, if it has the expected
		// number of fields.
		rawCSVData = append(rawCSVData, record)
	}

	// Output the number of records sucessfully read to stdout.
	log.Printf("Sucessfully parsed %d lines\n", len(rawCSVData))
}

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example3

// Sample program to read in records from an example CSV file,
// and catch an unexpected types in a single column.
package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	// Open the iris dataset file.
	csvFile, err := os.Open("../data/iris_mixed_types.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(csvFile)

	// Read in the records looking for unexpected types in the second column.
	var secondColumn []float64
	line := 1
	for {

		// Read in a row. Check if we are at the end of the file.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// Let's say that we want to gather the second column in the file
		// and validate that it includes only float values (e.g., because
		// we utilize this as a slice of floats later in our application.
		value, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		// Append the record to our slice, if it has the expected type.
		secondColumn = append(secondColumn, value)
		line++
	}

	// Output the number of records sucessfully read to stdout.
	log.Printf("Sucessfully parsed %d lines of the second column\n", len(secondColumn))
}

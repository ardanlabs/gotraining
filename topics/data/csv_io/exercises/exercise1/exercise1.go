// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./exercise1

// Sample program to read in records from an example CSV file,
// and catch an unexpected types in any of the columns.
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// CSVRecord contains a sucessfully parsed row of the CSV file.
type CSVRecord struct {
	SepalLength float64
	SepalWidth  float64
	PetalLength float64
	PetalWidth  float64
	Species     string
	ParseError  error
}

func main() {

	// Open the iris dataset file.
	f, err := os.Open("../../data/iris_multiple_mixed_types.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)

	// Create a slice value that will hold all of the successfully parsed
	// records from the CSV.
	var csvData []CSVRecord

	// line will help us keep track of line number for logging.
	line := 1

	// Read in the records looking for unexpected types.
	for {

		// Read in a row. Check if we are at the end of the file.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// Create a CSVRecord value for the row.
		var csvRecord CSVRecord

		// Parse each of the values in the record based on an expected type.
		for idx, value := range record {

			// Parse the value in the record as a string for the string column.
			if idx == 4 {

				// Validate that the value is not an empty string.  If the
				// value is an empty string break the parsing loop.
				if value == "" {
					log.Printf("Parsing line %d failed, unexpected type in column %d\n", line, idx)
					csvRecord.ParseError = fmt.Errorf("Empty string value")
					break
				}

				// Add the string value to the CSVRecord.
				csvRecord.Species = value
				continue
			}

			// Otherwise, parse the value in the record as a float64.
			// floatValue will hold the parsed float value of the record
			// for the numeric columns.
			var floatValue float64

			// If the value can not be parsed as a float, log and break the
			// parsing loop.
			if floatValue, err = strconv.ParseFloat(value, 64); err != nil {
				log.Printf("Parsing line %d failed, unexpected type in column %d\n", line, idx)
				csvRecord.ParseError = fmt.Errorf("Could not parse float")
				break
			}

			// Add the float value to the respective field in the CSVRecord.
			switch idx {
			case 0:
				csvRecord.SepalLength = floatValue
			case 1:
				csvRecord.SepalWidth = floatValue
			case 2:
				csvRecord.PetalLength = floatValue
			case 3:
				csvRecord.PetalWidth = floatValue
			}
		}

		// Append successfully parsed records to the slice defined above.
		if csvRecord.ParseError == nil {
			csvData = append(csvData, csvRecord)
		}

		// Increment the line counter.
		line++
	}

	// Output the number of records sucessfully parsed to stdout.
	log.Printf("Sucessfully parsed %d rows\n", len(csvData))
}

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./exercise2

// Sample program to read in records from an example CSV file,
// catch an unexpected types in any of the columns, and output
// processed data to a different CSV file.
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
}

func main() {

	// Read in and parse the CSV file into a slice of CSVRecord.
	csvData, err := cleanFile("../../data/iris_multiple_mixed_types.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Convert the records back in [][]string such that we can use
	// encoding/csv to save the records.
	var records [][]string
	for _, recordIn := range csvData {

		// Convert the float columns and add the species.
		recordOut := []string{
			strconv.FormatFloat(recordIn.SepalLength, 'f', 2, 64),
			strconv.FormatFloat(recordIn.SepalWidth, 'f', 2, 64),
			strconv.FormatFloat(recordIn.PetalLength, 'f', 2, 64),
			strconv.FormatFloat(recordIn.PetalWidth, 'f', 2, 64),
			recordIn.Species,
		}

		// Append the record.
		records = append(records, recordOut)
	}

	// Save the records to a CSV file called processed.csv.
	file, err := os.Create("processed.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a CSV writer.
	w := csv.NewWriter(file)

	// Write all the records out to the file.
	w.WriteAll(records)
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

// cleanFile parses and cleans the file similar to what we did in exercise1.
func cleanFile(filename string) ([]CSVRecord, error) {

	// Open the iris dataset file.
	csvFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(csvFile)

	// Create a slice value that will hold all of the successfully parsed
	// records from the CSV.
	var csvData []CSVRecord

	// Read in the records looking for unexpected types.
	line := 1
	for {

		// Read in a row. Check if we are at the end of the file.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// Parse each of the values in the record based on an expected type.
		var csvRecord CSVRecord
		var parseErr error
		for idx, value := range record {

			// Parse the value according to the execected type.
			var floatValue float64
			switch idx {
			case 0, 1, 2, 3:
				if floatValue, err = strconv.ParseFloat(value, 64); err != nil {
					log.Printf("Parsing line %d failed, unexpected type in column %d\n", line, idx)
					parseErr = err
					continue
				}
			case 4:
				if value == "" {
					log.Printf("Parsing line %d failed, unexpected type in column %d\n", line, idx)
					parseErr = fmt.Errorf("Could not parse string value")
					continue
				}
				csvRecord.Species = value

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
		if parseErr != nil {
			continue
		}

		// Append successfully parsed records to the slice defined above.
		csvData = append(csvData, csvRecord)
		line++
	}

	return csvData, nil
}

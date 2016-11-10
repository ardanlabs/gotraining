// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example3

// Sample program to read in records from an example CSV file,
// and catch an unexpected types in a single column.
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/gonum/stat"
)

func main() {

	// Open the continuous observations and predictions.
	csvFile, err := os.Open("../data/continuous.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(csvFile)

	// Read in the records looking for unexpected types in the columns.
	var observed []float64
	var predicted []float64
	line := 1
	for {

		// Read in a row. Check if we are at the end of the file.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// Skip the header.
		if line == 1 {
			line++
			continue
		}

		// Read in the observed and predicted values.
		observedVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		predictedVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		// Append the record to our slice, if it has the expected type.
		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++
	}

	// Calculate the R^2 value.
	rSquared := stat.RSquaredFrom(observed, predicted, nil)

	// Output the R^2 value to standard out.
	fmt.Printf("\nR^2 = %0.2f\n\n", rSquared)
}

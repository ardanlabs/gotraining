// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./template2

// Sample program to calculate a mean squared error.
package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	// Open the continuous observations and predictions.
	f, err := os.Open("../../data/continuous.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)

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

	// Calculate the mean squared error.

	// Output the MSE value to standard out.
}

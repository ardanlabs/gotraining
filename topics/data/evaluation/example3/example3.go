// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example3

// Sample program to calculate a accuracy.
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	// Open the binary observations and predictions.
	f, err := os.Open("../data/labeled.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)

	// observed and predicted will hold the parsed observed and predicted values
	// form the labeled data file.
	var observed []int
	var predicted []int

	// line will track row numbers for logging.
	line := 1

	// Read in the records looking for unexpected types in the columns.
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
		observedVal, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		predictedVal, err := strconv.Atoi(record[1])
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		// Append the record to our slice, if it has the expected type.
		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++
	}

	// This variable will hold our count of true positive and
	// true negative values.
	var truePosNeg int

	// Accumulate the true positive/negative count.
	for idx, oVal := range observed {
		if oVal == predicted[idx] {
			truePosNeg++
		}
	}

	// Calculate the accuracy (subset accuracy).
	accuracy := float64(truePosNeg) / float64(len(observed))

	// Output the Accuracy value to standard out.
	fmt.Printf("\nAccuracy = %0.2f\n\n", accuracy)
}

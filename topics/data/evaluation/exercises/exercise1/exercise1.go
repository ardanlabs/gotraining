// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./exercise1

// Sample program to calculate specificity.
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

	// Open the labeled observations and predictions.
	f, err := os.Open("../../data/labeled.csv")
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

	// classes contains the three possible classes in the labeled data.
	classes := []int{0, 1, 2}

	// Loop over each class.
	for _, class := range classes {

		// These variables will hold our count of true negatives and
		// our total count of negatives.
		var trueNeg int
		var negatives int

		// Accumulate the true negative and total negative counts.
		for idx, oVal := range observed {

			// Test to see if we have a negative case.
			if oVal != class {

				// Increment of tally of negatives.
				negatives++

				// See if we correctly identified the negative case.
				if predicted[idx] != class {
					trueNeg++
				}
			}
		}

		// Calculate the specificity.
		specificity := float64(trueNeg) / float64(negatives)

		// Output the specificity value to standard out.
		fmt.Printf("\nSpecificity (class %d) = %0.2f", class, specificity)
	}

	// Spacing for pretty output.
	fmt.Printf("\n\n")
}

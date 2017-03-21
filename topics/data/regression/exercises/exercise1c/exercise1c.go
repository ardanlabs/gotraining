// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./exercise1c

// Sample program to validate a trained multiple regression model on a holdout data set.
package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/pachyderm/pachyderm/src/client"
)

func main() {

	// Connect to Pachyderm on our localhost.  By default
	// Pachyderm will be exposed on port 30650.
	c, err := client.NewFromAddress("0.0.0.0:30650")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Get the holdout dataset from Pachyderm's data
	// versioning at the latest commit.
	var b bytes.Buffer
	if err := c.GetFile("regression_split", "master", "holdout.csv", 0, 0, "", false, nil, &b); err != nil {
		log.Fatal()
	}

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(bytes.NewReader(b.Bytes()))

	// Read in all of the CSV records
	reader.FieldsPerRecord = 11
	holdoutData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Loop over the holdout data predicting y and evaluating the prediction
	// with the mean absolute error.
	var mAE float64
	for i, record := range holdoutData {

		// Skip the header.
		if i == 0 {
			continue
		}

		// Parse the observed diabetes progression measure, or "y".
		yObserved, err := strconv.ParseFloat(record[10], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Parse the bmi value.
		bmiVal, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Parse the ltg value.
		ltgVal, err := strconv.ParseFloat(record[8], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Predict y with our trained model.
		yPredicted := predict(bmiVal, ltgVal)

		// Add the to the mean absolute error.
		mAE += math.Abs(yObserved-yPredicted) / float64(len(holdoutData))
	}

	// Output the MAE to standard out.
	fmt.Printf("\nMAE = %0.2f\n\n", mAE)
}

// predict uses our trained regression model to made a prediction based on a
// bmi and ltg value.
func predict(bmi, ltg float64) float64 {
	return 151.50 + bmi*623.59 + ltg*644.50
}

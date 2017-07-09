// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./template1c

// Sample program to validate a trained multiple regression model on a holdout data set.
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	// Open the holdout dataset file.
	f, err := os.Open("../../data/holdout.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)

	// Read in all of the CSV records
	reader.FieldsPerRecord = 11
	holdoutData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Loop over the holdout data predicting y and evaluating the prediction
	// with the mean absolute error.

	// Output the MAE to standard out.
	fmt.Printf("\nMAE = %0.2f\n\n", mAE)
}

// Create a "predict" function that uses our trained regression model
// to made a prediction based on a bmi and ltg value.

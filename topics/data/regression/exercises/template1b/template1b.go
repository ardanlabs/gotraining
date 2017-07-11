// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./template1b

// Sample program to train and test a multiple regression model.
package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {

	// Open the training dataset file.
	f, err := os.Open("../../data/training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)

	// Read in all of the CSV records
	reader.FieldsPerRecord = 11
	trainingData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// In this case we are going to try and model our disease measure
	// y by the bmi feature, another feature of your choice, plus an
	// intercept.  As such, let's create the struct needed to train
	// a model using github.com/sajari/regression.

	// Loop over the CSV records adding the training data.

	// Train/fit the regression model.

	// Output the trained model parameters.

	// Open the test dataset file.
	f, err := os.Open("../../data/test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a CSV reader reading from the opened file.
	reader = csv.NewReader(f)

	// Read in all of the CSV records
	reader.FieldsPerRecord = 11
	testData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Loop over the test data predicting y and evaluating the prediction
	// with the mean absolute error.

	// Output the MAE to standard out.
}

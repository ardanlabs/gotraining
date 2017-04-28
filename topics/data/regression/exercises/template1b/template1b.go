// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./template1b

// Sample program to train and test a multiple regression model.
package main

import (
	"bytes"
	"encoding/csv"
	"log"

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

	// Get the training dataset from Pachyderm's data
	// versioning at the latest commit.
	var b bytes.Buffer
	if err := c.GetFile("regression_split", "master", "training.csv", 0, 0, &b); err != nil {
		log.Fatal()
	}

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(bytes.NewReader(b.Bytes()))

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

	// Get the test dataset from Pachyderm's data
	// versioning at the latest commit.
	b.Reset()
	if err := c.GetFile("regression_split", "master", "test.csv", 0, 0, &b); err != nil {
		log.Fatal()
	}

	// Create a CSV reader reading from the opened file.
	reader = csv.NewReader(bytes.NewReader(b.Bytes()))

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

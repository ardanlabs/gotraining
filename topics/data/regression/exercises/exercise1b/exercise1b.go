// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./exercise1b

// Sample program to train and test a multiple regression model.
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/sajari/regression"
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
	// y by the bmi feature plust an intercept.  As such, let's create
	// the struct needed to train a model using github.com/sajari/regression.
	var r regression.Regression
	r.SetObserved("diabetes progression")
	r.SetVar(0, "bmi")
	r.SetVar(1, "ltg")

	// Loop over the CSV records adding the training data.
	for i, record := range trainingData {

		// Skip the header.
		if i == 0 {
			continue
		}

		// Parse the diabetes progression measure, or "y".
		yVal, err := strconv.ParseFloat(record[10], 64)
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

		// Add these points to the regression value.
		r.Train(regression.DataPoint(yVal, []float64{bmiVal, ltgVal}))
	}

	// Train/fit the regression model.
	r.Run()

	// Output the trained model parameters.
	fmt.Printf("\nRegression Formula:\n%v\n\n", r.Formula)

	// Open the test dataset file.
	f, err = os.Open("../../data/test.csv")
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
	var mAE float64
	for i, record := range testData {

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
		yPredicted, err := r.Predict([]float64{bmiVal, ltgVal})

		// Add the to the mean absolute error.
		mAE += math.Abs(yObserved-yPredicted) / float64(len(testData))
	}

	// Output the MAE to standard out.
	fmt.Printf("MAE = %0.2f\n\n", mAE)
}

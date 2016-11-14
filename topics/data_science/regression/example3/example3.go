// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example3

// Sample program to create training, test, and holdout data sets.
package main

import (
	"io/ioutil"
	"log"

	"github.com/kniren/gota/data-frame"
)

func main() {

	// Pull in the CSV data.
	diabetesData, err := ioutil.ReadFile("../data/diabetes.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Create a dataframe from the CSV string.
	// The types of the columns will be inferred.
	diabetesDF := df.ReadCSV(string(diabetesData))

	// Calculate the number of elements in each set.
	trainingNum := diabetesDF.Nrow() / 2
	testNum := diabetesDF.Nrow() / 4
	holdoutNum := diabetesDF.Nrow() / 4
	if trainingNum+testNum+holdoutNum < diabetesDF.Nrow() {
		trainingNum++
	}

	// Create the subset indices.
	trainingIdx := make([]int, trainingNum)
	testIdx := make([]int, testNum)
	holdoutIdx := make([]int, holdoutNum)

	// Enumerate the training indices.
	for i := 0; i < trainingNum; i++ {
		trainingIdx[i] = i
	}

	// Enumerate the test indices.
	for i := 0; i < testNum; i++ {
		testIdx[i] = trainingNum + i
	}

	// Enumerate the holdout indices.
	for i := 0; i < holdoutNum; i++ {
		holdoutIdx[i] = trainingNum + testNum + i
	}

	// Create the subset dataframes.
	trainingDF := diabetesDF.Subset(trainingIdx)
	testDF := diabetesDF.Subset(testIdx)
	holdoutDF := diabetesDF.Subset(holdoutIdx)

	// Save the training data.
	b, err := trainingDF.SaveCSV()
	if err != nil {
		log.Fatal(err)
	}
	if err = ioutil.WriteFile("../data/training.csv", b, 0644); err != nil {
		log.Fatal(err)
	}

	// Save the test data.
	b, err = testDF.SaveCSV()
	if err != nil {
		log.Fatal(err)
	}
	if err = ioutil.WriteFile("../data/test.csv", b, 0644); err != nil {
		log.Fatal(err)
	}

	// Save the holdout data.
	b, err = holdoutDF.SaveCSV()
	if err != nil {
		log.Fatal(err)
	}
	if err = ioutil.WriteFile("../data/holdout.csv", b, 0644); err != nil {
		log.Fatal(err)
	}

}

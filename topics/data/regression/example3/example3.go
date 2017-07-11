// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example3

// Sample program to create training, test, and holdout data sets.
package main

import (
	"bufio"
	"log"
	"os"
	"path"

	"github.com/kniren/gota/dataframe"
)

func main() {

	// Open the diabetes dataset file.
	f, err := os.Open("../data/diabetes.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	diabetesDF := dataframe.ReadCSV(f)

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

	// Create a map that will be used in writing the data
	// to files.
	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testDF,
		2: holdoutDF,
	}

	// Create the respective files.
	for idx, setName := range []string{"training.csv", "test.csv", "holdout.csv"} {

		// Save the filtered dataset file.
		f, err := os.Create(path.Join("../data/", setName))
		if err != nil {
			log.Fatal(err)
		}

		// Create a buffered writer.
		w := bufio.NewWriter(f)

		// Write the dataframe out as a CSV.
		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}
	}
}

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example2

// Sample program to determine an optimal value of the decision tree pruning parameter.
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"os"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/trees"
)

func main() {

	// Open the iris dataset file.
	f, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create golearn instances from the iris data.
	irisData, err := CreateInstancesFromReader(f)
	if err != nil {
		log.Fatal(err)
	}

	// Print a header for the output.
	fmt.Printf("Parameter\tAccuracy\n")

	// Define the parameters we will test.
	params := []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8}
	for _, param := range params {

		// Seed the random number generator.
		rand.Seed(44111342)

		// Define the decision tree model.
		tree := trees.NewID3DecisionTree(param)

		// Perform the cross validation.
		cfs, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(irisData, tree, 5)
		if err != nil {
			panic(err)
		}

		// Calculate the metrics.
		mean, variance := evaluation.GetCrossValidatedMetric(cfs, evaluation.GetAccuracy)
		stdev := math.Sqrt(variance)

		// Output the results to standard out.
		fmt.Printf("%0.2f\t\t%.2f (+/- %.2f)\n", param, mean, stdev*2)
	}
}

// CreateInstancesFromReader creates golearn instances from
// an io.Reader.
func CreateInstancesFromReader(r io.Reader) (*base.DenseInstances, error) {

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(r)

	// Read in all of the CSV records
	reader.FieldsPerRecord = 5
	tmpData, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Create the output file.
	f, err := os.Create("/tmp/iris.csv")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Create a CSV writer.
	w := csv.NewWriter(f)

	// Write all the records out to the file. Note, this can
	// also we done record by record, with a final call to
	// Flush().
	w.WriteAll(tmpData)
	if err := w.Error(); err != nil {
		return nil, err
	}

	// Read in the iris data set into golearn "instances".
	instances, err := base.ParseCSVToInstances("/tmp/iris.csv", true)
	if err != nil {
		return nil, err
	}

	return instances, nil
}

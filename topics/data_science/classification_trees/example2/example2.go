// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example2

// Sample program to determine an optimal value of the decision tree pruning parameter.
package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/trees"
)

func main() {

	// Read in the iris data set into golearn "instances".
	irisData, err := base.ParseCSVToInstances("../data/iris.csv", true)
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

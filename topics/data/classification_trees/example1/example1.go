// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example1

// Sample program to train and validate a decision tree model with cross validation.
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

	// This is to seed the random processes involved in building the
	// decision tree.
	rand.Seed(44111342)

	// We will use the ID3 algorithm to build our decision tree.  Also, we
	// will start with a parameter of 0.6 that controls the train-prune split.
	tree := trees.NewID3DecisionTree(0.6)

	// Use cross-fold validation to successively train and evalute the model
	// on 5 folds of the data set.
	cfs, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(irisData, tree, 5)
	if err != nil {
		log.Fatal(err)
	}

	// Get the mean, variance and standard deviation of the accuracy for the
	// cross validation.
	mean, variance := evaluation.GetCrossValidatedMetric(cfs, evaluation.GetAccuracy)
	stdev := math.Sqrt(variance)

	// Output the cross metrics to standard out.
	fmt.Printf("\nAccuracy\n%.2f (+/- %.2f)\n\n", mean, stdev*2)
}

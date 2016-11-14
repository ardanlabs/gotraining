// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example2

// Sample program to train and validate a kNN model with cross validation.
package main

import (
	"fmt"
	"log"
	"math"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func main() {

	// Read in the iris data set into golearn "instances".
	irisData, err := base.ParseCSVToInstances("../data/iris.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize a new KNN classifier.
	cls := knn.NewKnnClassifier("euclidean", 2)

	// Use cross-fold validation to successively train and evalute the model
	// on 5 folds of the data set.
	cfs, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(irisData, cls, 5)
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

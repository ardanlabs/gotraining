// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./exercise1

// Program for finding an optimal k value for a k nearest neighbors model.
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
	irisData, err := base.ParseCSVToInstances("../../data/iris.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	// Enumerate some possible k values.
	var ks [30]int
	for i := 0; i < len(ks); i++ {
		ks[i] = i + 2
	}

	// results will hold the results for each k value.
	var results [30]string

	// Loop over k values, evaluting the resulting predictions.
	for i, k := range ks {

		// Initialize a new KNN classifier.
		cls := knn.NewKnnClassifier("euclidean", k)

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

		// Format the results.
		results[i] = fmt.Sprintf("%d\t\t%.2f (+/- %.2f)", k, mean, stdev*2)
	}

	// Output the results to standard out.
	fmt.Printf("\nk value\tAccuracy\n")
	for _, result := range results {
		fmt.Printf("%s\n", result)
	}
}

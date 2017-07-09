// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./exercise1

// Program for finding an optimal k value for a k nearest neighbors model.
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func main() {

	// Open the iris dataset file.
	f, err := os.Open("../../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create golearn instances from the iris data.
	irisData, err := CreateInstancesFromReader(f)
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
		cls := knn.NewKnnClassifier("euclidean", "linear", k)

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

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./template1.go

// Sample program to visualize the accuracy of models with various
// decision tree pruning parameters.
package main

import (
	"log"

	"github.com/sjwhitworth/golearn/base"
)

func main() {

	// Read in the iris data set into golearn "instances".
	irisData, err := base.ParseCSVToInstances("../../data/iris.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	// Define the parameters we will test.
	var params [90]float64
	params[0] = 0.1
	for i := 1; i < len(params); i++ {
		params[i] = params[i-1] + 0.01
	}

	// Define a slice to hold the accuracy results for each pruning parameter.

	// Loop over the parameter choices.

	// Define the decision tree model.

	// Perform the cross validation.

	// Calculate the accuracy.

	// Create a new plot.

	// Add our data to the plot.

	// Save the plot to a PNG file.
}

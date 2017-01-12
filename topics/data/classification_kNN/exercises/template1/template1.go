// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./template1

// Template programe for finding an optimal k value for a k nearest neighbors model.
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

	// Enumerate some possible k values.

	// Loop over k values, evaluting the resulting predictions with cross validation.

	// Output the results to standard out.
}

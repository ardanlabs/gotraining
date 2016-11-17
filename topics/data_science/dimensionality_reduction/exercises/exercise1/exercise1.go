// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./exercise1

// Sample program to project iris data on to 3 principal components.
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gonum/matrix/mat64"
	"github.com/gonum/stat"
)

func main() {

	// Open the iris dataset file.
	f, err := os.Open("../../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 5

	// Read in all of the CSV records
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// floatData will hold all the float values that will eventually be
	// used to form out matrix.
	floatData := make([]float64, 4*len(rawCSVData))

	// dataIndex will track the current index of the matrix values.
	var dataIndex int

	// Sequentially move the rows into a slice of floats.
	for _, record := range rawCSVData {

		// Loop over the float columns.
		for i := 0; i < 4; i++ {

			// Convert the value to a float.
			val, err := strconv.ParseFloat(record[i], 64)
			if err != nil {
				log.Fatal(fmt.Errorf("Could not parse float value"))
			}

			// Add the float value to the slice of floats.
			floatData[dataIndex] = val
			dataIndex++
		}
	}

	// Form the matrix.
	mat := mat64.NewDense(len(rawCSVData), 4, floatData)

	// Calculate the principal component direction vectors
	// and variances.
	vecs, _, ok := stat.PrincipalComponents(mat, nil)
	if !ok {
		log.Fatal("Could not calculate prinicple components")
	}

	// Project the data onto the first 3 principal components.
	var proj mat64.Dense
	proj.Mul(mat, vecs.View(0, 0, 4, 3))

	// Output the resulting projected features to standard out.
	fmt.Printf("proj = %.4f\n", mat64.Formatted(&proj, mat64.Prefix("       ")))
}

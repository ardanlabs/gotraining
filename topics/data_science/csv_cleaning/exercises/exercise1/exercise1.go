// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./exercise1

// Sample program to read in a CSV, create three filtered datasets, and
// save those datasets to three separate files.
package main

import (
	"io/ioutil"
	"log"

	"github.com/kniren/gota/data-frame"
)

func main() {

	// Pull in the CSV data.
	irisData, err := ioutil.ReadFile("../../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Create a dataframe from the CSV string.
	// The types of the columns will be inferred.
	irisDF := df.ReadCSV(string(irisData))

	// Filter the dataset into three separate dataframes,
	// each corresponding to one of the Iris species.
	speciesNames := []string{
		"Iris-setosa",
		"Iris-versicolor",
		"Iris-virginica",
	}
	for _, species := range speciesNames {

		// Filer the original dataset.
		filter := df.F{
			Colname:    "species",
			Comparator: "==",
			Comparando: species,
		}
		filtered := irisDF.Filter(filter)

		// Save the filtered dataset.
		b, err := filtered.SaveCSV()
		if err != nil {
			log.Fatal(err)
		}

		if err := ioutil.WriteFile(species+".csv", b, 0644); err != nil {
			log.Fatal(err)
		}
	}
}

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./exercise1

// Sample program to read in a CSV, create three filtered datasets, and
// save those datasets to three separate files.
package main

import (
	"bufio"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {

	// Pull in the CSV file.
	irisFile, err := os.Open("../../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	irisDF := dataframe.ReadCSV(irisFile)

	// Define the names of the three separate species contained in the CSV file.
	speciesNames := []string{
		"Iris-setosa",
		"Iris-versicolor",
		"Iris-virginica",
	}

	// Filter the dataset into three separate dataframes,
	// each corresponding to one of the Iris species.
	for _, species := range speciesNames {

		// Filer the original dataset.
		filter := dataframe.F{
			Colname:    "species",
			Comparator: "==",
			Comparando: species,
		}
		filtered := irisDF.Filter(filter)

		// Save the filtered dataset file.
		f, err := os.Create(species + ".csv")
		if err != nil {
			log.Fatal(err)
		}

		// Create a buffered writer.
		w := bufio.NewWriter(f)

		// Write the contents of the dataframe
		if err := filtered.WriteCSV(w); err != nil {
			log.Fatal(err)
		}
	}
}

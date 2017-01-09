// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./template1

// Sample program to generate a box plot of diabetes bmi values.
package main

import (
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {

	// Pull in the CSV file.
	diabetesFile, err := os.Open("../../data/diabetes.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer diabetesFile.Close()

	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	diabetesDF := dataframe.ReadCSV(diabetesFile)

	// Create the plot and set its title and axis label.

	// Create the box for our data.

	// Extract the bmi col as a slice of floats.

	// Create a plotter.Values value and fill it with the
	// values from the respective column of the dataframe.

	// Add the data to the plot.

	// Set the X axis of the plot to nominal with
	// the given names for x=0, x=1, etc.

	// Save the plot.
}

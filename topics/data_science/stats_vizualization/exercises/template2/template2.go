// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./template2

// Sample program to generate a histogram of diabetes bmi values.
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

	// Create a plotter.Values value and fill it with the
	// values from the respective column of the dataframe.

	// Make a plot and set its title.

	// Create a histogram of our values drawn
	// from the standard normal.

	// Normalize the histogram.

	// Add the histogram to the plot.

	// Save the plot to a PNG file.
}

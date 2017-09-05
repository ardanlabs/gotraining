// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./exercise1

// Sample program to calculate both central tendency and statistical dispersion
// measures for the iris dataset.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/stat"
)

func main() {

	// Open the iris dataset file.
	f, err := os.Open("../../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	irisDF := dataframe.ReadCSV(f)

	// Loop over the float columns.
	for _, colName := range irisDF.Names() {

		// Only look at the numeric columns.
		if colName != "species" {

			// Get the float values from the column.
			col := irisDF.Col(colName).Float()

			// Calculate the Mean of the variable.
			meanVal := stat.Mean(col, nil)

			// Calculate the Mode of the variable.
			modeVal, modeCount := stat.Mode(col, nil)

			// Calculate the Median of the variable.
			medianVal, err := stats.Median(col)
			if err != nil {
				log.Fatal(err)
			}

			// Calculate the variance of the variable.
			varianceVal := stat.Variance(col, nil)

			// Calculate the standard deviation of the variable.
			stdDevVal := stat.StdDev(col, nil)

			// Output the results to standard out.
			fmt.Printf("\n%s Summary Statistics:\n", colName)
			fmt.Printf("Mean value: %0.2f\n", meanVal)
			fmt.Printf("Mode value: %0.2f\n", modeVal)
			fmt.Printf("Mode count: %d\n", int(modeCount))
			fmt.Printf("Median value: %0.2f\n", medianVal)
			fmt.Printf("Variance value: %0.2f\n", varianceVal)
			fmt.Printf("Std Dev value: %0.2f\n", stdDevVal)
		}
	}
}

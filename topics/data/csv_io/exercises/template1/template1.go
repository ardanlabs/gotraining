// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./template1

// Sample program to read in records from an example CSV file,
// and catch an unexpected types in any of the columns.
package main

import (
	"encoding/csv"
	"log"
	"os"
)

// Define a struct that will contain a sucessfully parsed row of the CSV file.

func main() {

	// Open the iris dataset file.
	f, err := os.Open("../../data/iris_multiple_mixed_types.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)

	// Create a slice value that will hold all of the successfully parsed
	// records from the CSV.

	// Read in the records looking for unexpected types.

	// Read in a row. Check if we are at the end of the file.

	// Create a CSVRecord value for the row.

	// Parse each of the values in the record based on an expected type.

	// Parse the value in the record as a string for the string column.

	// Otherwise, parse the value in the record as a float64.

	// Append successfully parsed records to the slice defined above.

	// Output the number of records sucessfully parsed to stdout.
}

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./template1

// Sample program to read in records from an example CSV file,
// and catch an unexpected types in any of the columns.
package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

// Define a struct here that will hold the successfully parsed
// records in each column of our CSV.

func main() {

	// Open the iris dataset file.
	csvFile, err := os.Open("../data/iris_mixed_types.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(csvFile)

	// Create a value of the type you defined above that will hold
	// all of your successfully parsed records from the CSV.

	// Read in the records looking for unexpected types.
	line := 1
	for {

		// Read in a row. Check if we are at the end of the file.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// Parse each of the columns based on an expected type.

		// Handle any errors such that you only append successfully
		// parsed records to the value you defined above.  Also,
		// log any parsing errors.

		line++
	}

	// Output the number of records sucessfully parsed to stdout.
	log.Printf("Sucessfully parsed %d rows\n", len(secondColumn))
}

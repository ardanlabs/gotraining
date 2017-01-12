// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example4

// Sample program to illustrate maintaining integrity with Go
// in the presence of messy data.
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	// Open the CSV.
	f, err := os.Open("../data/example_messy.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Read in the CSV records.
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Get the maximum value in the integer column.
	var intMax int
	for _, record := range records {

		// Parse the integer value.
		intVal, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}

		// Replace the maximum value if appropriate.
		if intVal > intMax {
			intMax = intVal
		}
	}

	// Print the maxium value.
	fmt.Println(intMax)
}

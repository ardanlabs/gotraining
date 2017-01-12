// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example3

// Sample program to register of CSV file as an in-memory
// SQL database and execute SQL queries on the CSV.
package main

import (
	"fmt"
	"log"

	"github.com/go-hep/csvutil/csvdriver"
)

func main() {

	// Open the CSV file as a database table.
	db, err := csvdriver.Conn{
		File:    "../data/iris.csv",
		Comment: '#',
		Comma:   ',',
	}.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Start a database transaction.
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Commit()

	// Define a SQL query that we will execute against the CSV file.
	query := `SELECT var3, var4, var5 FROM csv WHERE var5 = "Iris-versicolor"`

	// Query the CSV via the SQL statement.  Here we will only get
	// the petal length, petal width, and species for all rows
	// where the species is "Iris-versicolor".
	rows, err := tx.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Output the results of the query to standard out.
	for rows.Next() {

		// Define variables for the result values.
		var (
			petalLength float64
			petalWidth  float64
			species     string
		)

		// Scan for the results.
		if err = rows.Scan(&petalLength, &petalWidth, &species); err != nil {
			log.Fatal(err)
		}

		// Output the results to standard out.
		fmt.Printf("petal length: %.2f, petal width: %.2f, species: %s\n",
			petalLength, petalWidth, species)
	}

	// Handle any errors from rows.Next().
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

}

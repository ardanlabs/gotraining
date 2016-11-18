// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./exercise2

// Sample program to register of CSV file as an in-memory
// SQL database, sum float columns, and output a process CSV.
package main

import (
	"log"

	"github.com/go-hep/csvutil"
	"github.com/go-hep/csvutil/csvdriver"
)

func main() {

	// Open the CSV file as a database table.
	db, err := csvdriver.Conn{
		File:    "../../data/iris.csv",
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

	// Define a query we will execute against our CSV file.
	query := `SELECT var1, var2, var3, var4, var5 FROM csv WHERE var1 != "sepal_length";`

	// Query the CSV via the SQL statement.
	rows, err := tx.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Register the output file as a table.
	fname := "processed.csv"
	tbl, err := csvutil.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	tbl.Writer.Comma = ';'
	defer tbl.Close()

	// Output the results of the query to a processed CSV file.
	for rows.Next() {

		// Scan the row for the values.
		var (
			f1      float64
			f2      float64
			f3      float64
			f4      float64
			species string
		)
		if err = rows.Scan(&f1, &f2, &f3, &f4, &species); err != nil {
			log.Fatal(err)
		}

		// Write the row.
		if err = tbl.WriteRow(f1+f2+f3+f4, species); err != nil {
			log.Fatal(err)
		}
	}

	// Make sure the output file is properly saved.
	if err := tbl.Close(); err != nil {
		log.Fatal(err)
	}
}

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example3

// Sample program to retrieve results from a database.
package main

import (
	"database/sql"
	"fmt"
	"log"

	// go-sqlite3 is the libary that allows us to connect
	// to sqlite with databases/sql.
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Open a database value.  Specify the sqlite3 driver
	// for databases/sql.
	db, err := sql.Open("sqlite3", "../data/iris.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query the database.
	rows, err := db.Query(`
		SELECT 
			sepal_length as sLength, 
			sepal_width as sWidth, 
			petal_length as pLength, 
			petal_width as pWidth 
		FROM iris
		WHERE species = ?`, "Iris-setosa")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate over the rows, sending the results to
	// standard out.
	for rows.Next() {
		var (
			sLength float64
			sWidth  float64
			pLength float64
			pWidth  float64
		)
		if err := rows.Scan(&sLength, &sWidth, &pLength, &pWidth); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%.2f, %.2f, %.2f, %.2f\n", sLength, sWidth, pLength, pWidth)
	}

	// Check for errors after we are done iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

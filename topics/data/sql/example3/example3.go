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
	"os"

	// pq is the libary that allows us to connect
	// to postgres with databases/sql.
	_ "github.com/lib/pq"
)

func main() {

	// Get my ElephantSQL postgres URL. I have it stored in
	// an environmental variable.
	pgURL := os.Getenv("PGURL")
	if pgURL == "" {
		log.Fatal("PGURL empty")
	}

	// Open a database value.  Specify the postgres driver
	// for databases/sql.
	db, err := sql.Open("postgres", pgURL)
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
		WHERE species = $1`, "Iris-setosa")
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

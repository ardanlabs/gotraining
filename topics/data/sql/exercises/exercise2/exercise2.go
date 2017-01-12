// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./exercise2

// Sample program to delete rows in a database table.
package main

import (
	"database/sql"
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

	// Delete the rows.
	res, err := db.Exec("DELETE FROM iris WHERE sepal_length > 6.0")
	if err != nil {
		log.Fatal(err)
	}

	// See how many rows where affected.
	rowCount, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	// Output the number of rows deleted to standard out.
	log.Printf("Deleted %d rows!\n", rowCount)
}

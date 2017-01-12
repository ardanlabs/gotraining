// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example1

// Sample program to connect to and ping a database connection.
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

	// sql.Open() does not establish any connections to the
	// database.  It just prepares the database connection value
	// for later use.  To make sure the database is available and
	// accessible, we will use db.Ping().
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Able to connect to the postgres database!")
}

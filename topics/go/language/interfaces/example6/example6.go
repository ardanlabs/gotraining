// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show type assertions using the comma-ok idiom.
package main

import (
	"fmt"
	"log"
)

type user struct {
	id   int
	name string
}

type userDB struct {
	host string
}

func (db userDB) findUser(id int) (*user, error) {

	// Pretend this comes from a database.
	return &user{id: id, name: "Anna Walker"}, nil
}

type mockDB struct{}

func (db mockDB) findUser(id int) (*user, error) {

	// This is hard coded so code can test against it.
	return &user{id: id, name: "Jacob Walker"}, nil
}

func main() {

	db := mockDB{}

	if err := run(db); err != nil {
		log.Fatal(err)
	}
}

type userFinder interface {
	findUser(id int) (*user, error)
}

func run(f userFinder) error {

	u, err := f.findUser(1234)
	if err != nil {
		return err
	}

	fmt.Printf("Found user %+v\n", u)

	// If the concrete type inside the interface value is of the
	// type userDB then "ok" will be true and "db" can be used.
	if db, ok := f.(userDB); ok {
		log.Println("queried", db.host)
	}

	return nil
}

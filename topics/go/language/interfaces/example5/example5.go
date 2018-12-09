// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show the syntax of type assertions.
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

func main() {

	db := userDB{host: "localhost:3434"}

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

	// Ideally the userFinder abstraction would encompass all of
	// the behavior you care about. But what if, for some reason,
	// you really need to get to the concrete value behind the
	// interface?

	// Can you get to the "host" field of the db struct inside
	// this interface variable? No, not directly. All you know is
	// the value has a method "findUser".
	// ./example5.go:58:26: f.host undefined (type userFinder has no field or method host)
	log.Println("queried", f.host)

	// You can use a type assertion if you know the concrete type.
	db := f.(userDB)
	log.Println("queried", db.host)

	return nil
}

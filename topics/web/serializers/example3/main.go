// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to use the XML encoder.
package main

import (
	"encoding/xml"
	"log"
	"os"
	"time"
)

// User represents a user in the system. Omit the LastName
// from Marshaling unless we have a value. Omit Age always.
type User struct {
	FirstName string `xml:"first_name"`
	LastName  string `xml:",omitempty"`
	Age       int    `xml:"-"`
	CreatedAt time.Time
	Admin     bool
	Bio       *string
}

func main() {

	// Encode a zero valued version of a user and write to stdout.
	err := xml.NewEncoder(os.Stdout).Encode(&User{})
	if err != nil {
		log.Fatal(err)
	}

	// Create a user value for Mary Jane.
	u := User{
		FirstName: "Mary",
		LastName:  "Jane",
	}

	// Encode the user value and write to stdout.
	err = xml.NewEncoder(os.Stdout).Encode(&u)
	if err != nil {
		log.Fatal(err)
	}
}

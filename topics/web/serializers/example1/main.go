// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to use the JSON encoder.
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

// User represents a user in the system. Omit the LastName
// from Marshaling unless we have a value. Omit Age always.
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:",omitempty"`
	Age       int    `json:"-"`
	CreatedAt time.Time
	Admin     bool
	Bio       *string
}

func main() {

	// Encode a zero valued version of a user and write to stdout.
	err := json.NewEncoder(os.Stdout).Encode(&User{})
	if err != nil {
		log.Fatal(err)
	}

	// Create a user value for Mary Jane.
	u := User{
		FirstName: "Mary",
		LastName:  "Jane",
	}

	// Encode the user value and write to stdout.
	err = json.NewEncoder(os.Stdout).Encode(&u)
	if err != nil {
		log.Fatal(err)
	}
}

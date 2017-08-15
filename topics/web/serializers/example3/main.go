// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to implement the json.Marshaler interface
// to dictate the marshaling.
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

// User represents a user in the system.
type User struct {
	FirstName string
	LastName  string
	Age       int
	CreatedAt time.Time
	Admin     bool
	Bio       *string
}

// MarshalJSON implements the json.Marshaler interface so we
// can dictate how the user is marshaled.
func (u *User) MarshalJSON() ([]byte, error) {

	// Create a document of key/value pairs for each field.
	m := map[string]interface{}{
		"first_name": u.FirstName,
		"CreatedAt":  u.CreatedAt,
		"Admin":      u.Admin,
		"Bio":        u.Bio,
	}

	// Omit the last name from the document unless
	// we have a value.
	if u.LastName != "" {
		m["LastName"] = u.LastName
	}

	// We always omit Age so nothing to do.

	return json.Marshal(m)
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

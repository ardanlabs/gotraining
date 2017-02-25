// All material is licensed under the Apache License Version 2.0, June 2016
// http://www.apache.org/licenses/LICENSE-2.0

// Example shows how to reflect over a struct type value that is stored
// inside an interface value.
package main

import (
	"fmt"
	"reflect"
)

// user represents a basic user in the system.
type user struct {
	name     string
	age      int
	building float32
	secure   bool
	roles    []string
}

func main() {

	// Create a struct type user value.
	u := user{
		name:     "Cindy",
		age:      27,
		building: 321.45,
		secure:   true,
		roles:    []string{"admin", "developer"},
	}

	// Store a copy of the user value inside an empty interface value.
	var i interface{} = u

	// Display information about the user value that was stored.
	v := reflect.ValueOf(i)
	fmt.Printf("Kind: %v\tType: %v\t\tNumFields: %v\n", v.Kind(), v.Type(), v.NumField())
}

// All material is licensed under the Apache License Version 2.0, June 2016
// http://www.apache.org/licenses/LICENSE-2.0

// Example shows how to reflect over a map of struct type values that
// are stored inside an interface value.
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

	// Create a map of struct type user values.
	um := map[string]user{
		"Cindy": {
			name:     "Cindy",
			age:      27,
			building: 321.45,
			secure:   true,
			roles:    []string{"admin", "developer"},
		},
		"Bill": {
			name:     "Bill",
			age:      40,
			building: 456.21,
			secure:   false,
			roles:    []string{"developer"},
		},
	}

	// Store a value of the map inside an empty interface value.
	var i interface{} = um

	// Display the information about the map of users values.
	v := reflect.ValueOf(i)
	fmt.Printf("Kind: %v\tType: %v\n", v.Kind(), v.Type())

	// Iterate over the map via reflection.
	for i, key := range v.MapKeys() {
		fmt.Println(i, ":", v.MapIndex(key))
	}
}

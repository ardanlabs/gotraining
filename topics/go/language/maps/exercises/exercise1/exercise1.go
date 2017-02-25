// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs.
package main

import "fmt"

func main() {

	// Declare and make a map of integer type values.
	departments := make(map[string]int)

	// Initialize some data into the map.
	departments["IT"] = 20
	departments["Marketing"] = 15
	departments["Executives"] = 5
	departments["Sales"] = 50
	departments["Security"] = 8

	// Display each key/value pair.
	for key, value := range departments {
		fmt.Printf("Dept: %s People: %d\n", key, value)
	}
}

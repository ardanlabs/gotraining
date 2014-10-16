// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/wtQ54cnEpk

// Declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs.
package main

import (
	"fmt"
)

// main is the entry point for the application.
func main() {
	// Declare and make a map of integer type values.
	departments := make(map[string]int)

	// Intialize some data into our map of maps.
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

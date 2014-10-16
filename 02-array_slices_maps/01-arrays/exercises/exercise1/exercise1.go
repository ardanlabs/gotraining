// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/PvjKR2uU2H

// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
package main

import "fmt"

// main is the entry point for the application.
func main() {
	// Declare string arrays.
	var names [5]string
	friends := [5]string{"Joe", "Ed", "Jim", "Erick", "Bill"}

	// Asssign my friends to the names array.
	names = friends

	// Display each name in names.
	for _, name := range names {
		fmt.Println(name)
	}
}

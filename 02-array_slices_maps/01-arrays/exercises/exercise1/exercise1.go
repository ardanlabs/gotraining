// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/Pa3mrTCcpB

// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
package main

import "fmt"

// main is the entry point for the application.
func main() {
	// Declare string arrays to hold names.
	var names [5]string

	// Declare an array pre-populated with friend's names.
	friends := [5]string{"Joe", "Ed", "Jim", "Erick", "Bill"}

	// Assign the array of friends to the names array.
	names = friends

	// Display each name in names.
	for _, name := range names {
		fmt.Println(name)
	}
}

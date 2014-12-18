// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/0K0dlgG9yq

// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
package main

import "fmt"

// main is the entry point for the application.
func main() {
	// Declare string arrays to hold names.
	var array_name1 [N]type

	// Declare an array pre-populated with friend's names.
	array_name2 := [N]type{"name", "name", "name", "name", "name"}

	// Asssign the array of friends to the names array.
	array_name1 = array_name2

	// Display each name in names.
	for _, value_name := range array_name1 {
		fmt.Println(value_name)
	}
}

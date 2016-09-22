// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
// Display the string value and address of each element.
package main

import "fmt"

func main() {

	// Declare string arrays to hold names.
	var names [5]string

	// Declare an array pre-populated with friend's names.
	friends := [5]string{"Joe", "Ed", "Jim", "Erick", "Bill"}

	// Assign the array of friends to the names array.
	names = friends

	// Display each string value and address index in names.
	for i, name := range names {
		fmt.Println(name, &names[i])
	}
}

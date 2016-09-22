// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how arrays of different sizes are
// not of the same type.
package main

import "fmt"

func main() {

	// Declare an array of 5 integers that is initialized
	// to its zero value.
	var five [5]int

	// Declare an array of 4 integers that is initialized
	// with some values.
	four := [4]int{10, 20, 30, 40}

	// Assign one array to the other
	five = four

	// ./example2.go:21: cannot use four (type [4]int) as type [5]int in assignment

	fmt.Println(four)
	fmt.Println(five)
}

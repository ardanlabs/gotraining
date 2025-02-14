// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how the behavior of the for range and
// how memory for an array is contiguous.
package main

import "fmt"

func main() {

	// Declare an array of 5 strings initialized with values.
	friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	// ----- ruben
	for i := range friends {
		fmt.Println(friends[i], &friends[i])
	}

	fmt.Printf("\n\n")
	// ----- ruben

	// Iterate over the array displaying the value and
	// address of each element.
	for i, v := range friends {
		fmt.Printf("Value[%s]\tAddress[%p] IndexAddr[%p]\n", v, &v, &friends[i])
	}
}

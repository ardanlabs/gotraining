// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/irrA08aCkm

// Sample program to show how to declare and iterate over
// arrays of different types.
package main

import (
	"fmt"
)

// main is the entry point for the application.
func main() {
	// Declare an array of five strings that is initialized
	// to its zero value.
	var strings [5]string
	strings[0] = "Apple"
	strings[1] = "Orange"
	strings[2] = "Banana"
	strings[3] = "Grape"
	strings[4] = "Plum"

	// Iterate over the array of strings.
	for index, fruit := range strings {
		fmt.Println(index, fruit)
	}

	// Declare an array of 4 integers that is initalized
	// with some values.
	numbers := [4]int{10, 20, 30, 40}

	// Iterate over the array of numbers.
	for index := 0; index < len(numbers); index++ {
		fmt.Println(index, numbers[index])
	}
}

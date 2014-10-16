// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/LVD43cYBNS

// Sample program to show how arrays of different sizes are
// not of the same type.
package main

import (
	"fmt"
)

// main is the entry point for the application.
func main() {
	// Declare an array of 5 integers that is initalized
	// to its zero value.
	var five [5]int

	// Declare an array of 4 integers that is initalized
	// with some values.
	four := [4]int{10, 20, 30, 40}

	// Assign one array to the other
	five = four

	// ./example2.go:20: cannot use four (type [4]int) as type [5]int in assignment

	fmt.Println(four)
	fmt.Println(five)
}

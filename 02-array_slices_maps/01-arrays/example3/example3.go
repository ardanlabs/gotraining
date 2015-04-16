// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// https://play.golang.org/p/s4BSgxz0Y3

// Sample program to show how the behavior of the for range and
// how memory for an array is contiguous.
package main

import (
	"fmt"
)

// main is the entry point for the application.
func main() {
	// Declare an array of 5 strings initalized with values.
	five := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	// Iterate over the array displaying the value and
	// address of each element.
	for i, v := range five {
		fmt.Printf("Value[%s] Address[%p] IndexAddr[%p]\n", v, &v, &five[i])
	}
}

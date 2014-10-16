// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/HV5t0VrRie

// Sample program to use a composite literal to initialize
// a slice to a length and capacity. Iterate over a slice using
// a for range.
package main

import (
	"fmt"
)

// main is the entry point for the application.
func main() {
	// Get the slice of Dog values.
	names := []string{"bill", "jack", "sammy", "jill", "choley", "harley", "jamie", "Ed", "Lisa", "Missy"}

	// Iterate through the slice and displays the values.
	for index, name := range names {
		fmt.Printf("Index[%d] Name[%s]\t- Addr Name[%p] Addr Elem[%p]\n", index, name, &name, &names[index])
	}
}

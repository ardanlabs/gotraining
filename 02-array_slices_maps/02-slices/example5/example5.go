// http://play.golang.org/p/HV5t0VrRie

// Sample program to show how to iterate over a slice.
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

// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/Sl-vYp7pp_

// Sample program to show how to declare and initalize struct types.
package main

import (
	"fmt"
)

// example represents a type with different fields.
type example struct {
	flag    bool
	counter int16
	pi      float32
}

// main is the entry point for the application.
func main() {
	// Declare a variable of type example set to its
	// zero value.
	var e1 example

	// Display the value.
	fmt.Printf("%+v\n", e1)

	// Declare a variable of type example and init using
	// a composite literal.
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Display the field values.
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
}

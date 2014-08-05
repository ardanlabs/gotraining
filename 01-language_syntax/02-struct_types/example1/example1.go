// http://play.golang.org/p/lgXBEs4nx2

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
	// Declare variable of type example and init using
	// a composite literal.
	e := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Display the values.
	fmt.Printf("%+v\n", e)
	fmt.Println("Flag", e.flag)
	fmt.Println("Counter", e.counter)
	fmt.Println("Pi", e.pi)
}

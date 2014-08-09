// http://play.golang.org/p/ncYlNX7pcN

// Sample program to show how to declare and initalize anonymous
// struct types.
package main

import (
	"fmt"
)

// main is the entry point for the application.
func main() {
	// Declare a variable of an anonymous type and init
	// using a composite literal.
	e := struct {
		flag    bool
		counter int16
		pi      float32
	}{
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

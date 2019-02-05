// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show the syntax and mechanics of type
// switches and the empty interface.
package main

import "fmt"

func main() {

	// fmt.Println can be called with values of any type.
	fmt.Println("Hello, world")
	fmt.Println(12345)
	fmt.Println(3.14159)
	fmt.Println(true)

	// How can we do the same?
	myPrintln("Hello, world")
	myPrintln(12345)
	myPrintln(3.14159)
	myPrintln(true)

	// - An interface is satisfied by any piece of data when the data exhibits
	// the full method set of behavior defined by the interface.
	// - The empty interface defines no method set of behavior and therefore
	// requires no method by the data being stored.

	// - The empty interface says nothing about the data stored inside
	// the interface.
	// - Checks would need to be performed at runtime to know anything about
	// the data stored in the empty interface.
	// - Decouple around well defined behavior and only use the empty
	// interface as an exception when it is reasonable and practical to do so.
}

func myPrintln(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Printf("Is string  : type(%T) : value(%s)\n", v, v)
	case int:
		fmt.Printf("Is int     : type(%T) : value(%d)\n", v, v)
	case float64:
		fmt.Printf("Is float64 : type(%T) : value(%f)\n", v, v)
	default:
		fmt.Printf("Is unknown : type(%T) : value(%v)\n", v, v)
	}
}

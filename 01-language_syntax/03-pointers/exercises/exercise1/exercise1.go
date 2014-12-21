// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/pUtHSTN0Ef

// Declare and initalize a variable of type int with the value of 20. Display
// the _address of_ and _value of_ the variable.
//
// Declare and initialize a pointer variable of type int that points to the last
// variable you just created. Display the _address of_ , _value of_ and the
// _value that the pointer points to_.
package main

import "fmt"

// main is the entry point for the application.
func main() {
	// Declare an integer variable with the value of 20.
	value := 20

	// Display the address of and value of the variable.
	fmt.Println("Address Of:", &value, "Value Of:", value)

	// Declare a pointer variable of type int. Assign the
	// address of the integer variable above.
	p := &value

	// Display the address of, value of and the value the pointer
	// points to.
	fmt.Println("Address Of:", &p, "Value Of:", p, "Points To:", *p)
}

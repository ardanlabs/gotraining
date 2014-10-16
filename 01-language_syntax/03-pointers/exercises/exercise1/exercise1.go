// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/9XIYSnptUF

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
	value := 20
	fmt.Println("Address Of:", &value, "Value Of:", value)

	p := &value
	fmt.Println("Address Of:", &p, "Value Of:", p, "Points To:", *p)
}

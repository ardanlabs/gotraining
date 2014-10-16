// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/-SBwG9FnfJ

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initalize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

import (
	"fmt"
)

// user represents a user in the system.
type user struct {
	name  string
	email string
	age   int
}

// main is the entry point for the application.
func main() {
	// Declare variable of type example and init using
	// a composite literal.
	bill := user{
		name:  "Bill",
		email: "bill@ardanstudios.com",
		age:   45,
	}

	// Display the values.
	fmt.Println("Name", bill.name)
	fmt.Println("Email", bill.email)
	fmt.Println("Age", bill.age)

	// Declare a variable using an anonymous struct.
	ed := struct {
		name  string
		email string
		age   int
	}{
		name:  "Ed",
		email: "ed@ardanstudios.com",
		age:   46,
	}

	// Display the values.
	fmt.Println("Name", ed.name)
	fmt.Println("Email", ed.email)
	fmt.Println("Age", ed.age)
}

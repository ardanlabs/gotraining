// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/5uDVuormwB

// Sample program to show how to declare and use variadic functions.
package main

import "fmt"

// user is a struct type that declares user information.
type user struct {
	id   int
	name string
}

// main is the entry point for the application.
func main() {
	// Declare and initialize a value of type user.
	u1 := user{
		id:   1432,
		name: "Betty",
	}

	// Declare and initialize a value of type user.
	u2 := user{
		id:   4367,
		name: "Janet",
	}

	// Display both user values.
	display(u1, u2)
}

// display can accept and display multiple values of user types.
func display(users ...user) {
	for i := 0; i < len(users); i++ {
		fmt.Printf("%+v\n", users[i])
	}
}

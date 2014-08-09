// NEED PLAYGROUND

// Sample program to show how to declare and use variadic functions.
package main

import "fmt"

// user is a struct type that declares user information.
type user struct {
	ID   int
	Name string
}

// main is the entry point for the application.
func main() {
	// Declare and initalize a value of type user.
	u1 := user{
		ID:   1432,
		Name: "Betty",
	}

	// Declare and initalize a value of type user.
	u2 := user{
		ID:   4367,
		Name: "Janet",
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

// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/cK3y_qYUgd

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
	// Declare and initalize a value of type user.
	u1 := user{
		id:   1432,
		name: "Betty",
	}

	// Declare and initalize a value of type user.
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

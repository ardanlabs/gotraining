// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/gqsDjMd5bG

// Sample program to show how to embed a type into another type and
// the relationship between the inner and outer type.
package main

import (
	"fmt"
)

type (
	// user defines a user in the program.
	user struct {
		name  string
		email string
	}

	// admin represents an admin user with privileges.
	admin struct {
		user  // Embedded Type
		level string
	}
)

// notify implements a method that can be called via
// a value of type user.
func (u *user) notify() {
	fmt.Printf("user: Sending user email To %s<%s>\n",
		u.name,
		u.email)
}

// main is the entry point for the application.
func main() {
	// Create an admin user.
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// We can acces the inner type's method direectly.
	ad.user.notify()

	// The inner type's method is promoted.
	ad.notify()
}

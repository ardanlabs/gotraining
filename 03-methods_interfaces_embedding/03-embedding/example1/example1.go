// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/5LlI_KJ2ZT

// Sample program to show how what we are doing is NOT embedding
// a type but just using a type as a field.
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
		person user // NOT Embedding
		level  string
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
	admin := admin{
		person: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// We can acces fields methods.
	admin.person.notify()
}

// http://play.golang.org/p/AQlYR3zQqw

// Sample program to show how to embed a type into another type and
// the relationship between the inner and outer type.
package main

import (
	"fmt"
)

type (
	// User defines a user in the program.
	User struct {
		Name  string
		Email string
	}

	// Security defines security rights.
	Security struct {
		AccessLevel int
		Auth        bool
	}

	// Admin represents an admin user with privileges.
	Admin struct {
		User  // Embedded Type
		Sec   Security
		Level string
	}
)

// Notify implements a method that can be called via
// a value of type User.
func (u *User) Notify() {
	fmt.Printf("User: Sending User Email To %s<%s>\n",
		u.Name,
		u.Email)
}

// main is the entry point for the application.
func main() {
	// Create an admin user.
	admin := Admin{
		User: User{
			Name:  "john smith",
			Email: "john@email.com",
		},
		Sec: Security{
			AccessLevel: 10,
			Auth:        false,
		},
		Level: "super",
	}

	// We can acces the inner type's method direectly.
	admin.User.Notify()

	// The inner type's method is promoted.
	admin.Notify()
}

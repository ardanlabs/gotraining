// http://play.golang.org/p/-jGSPA8q1u

// Sample program to show how what we are doing is NOT embedding
// a type but just using a type as a field.
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

	// Admin represents an admin user with privileges.
	Admin struct {
		Person User // NOT Embedding
		Level  string
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
		Person: User{
			Name:  "john smith",
			Email: "john@yahoo.com",
		},
		Level: "super",
	}

	// We can acces fields methods.
	admin.Person.Notify()
}

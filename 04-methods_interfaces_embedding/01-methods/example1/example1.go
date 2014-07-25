// http://play.golang.org/p/jpf5IrVIxf

// Sample program to show how to declare methods and how the Go
// compiler supports them.
package main

import (
	"fmt"
)

// User defines a user in the program.
type User struct {
	Name  string
	Email string
}

// Notify implements a method that can be called via
// a value of type User.
func (u User) Notify() {
	fmt.Printf("User: Sending User Email To %s<%s>\n",
		u.Name,
		u.Email)
}

// ChangeEmail implements a method that can be called via
// a pointer of type User.
func (u *User) ChangeEmail(email string) {
	u.Email = email
}

// main is the entry point for the application.
func main() {
	// Value of type User can be used to call the method
	// with a value receiver.
	user1 := User{"Bill", "bill@email.com"}
	user1.Notify()

	// Pointer of type User can also be used to call a method
	// with a value receiver.
	user2 := &User{"Jill", "jill@email.com"}
	user2.Notify()

	// Value of type User can be used to call the method
	// with a pointer receiver.
	user1.ChangeEmail("bill@gmail.com")
	user1.Notify()

	// Pointer of type User can be used to call the method
	// with a pointer receiver.
	user2.ChangeEmail("jill@gmail.com")
	user2.Notify()
}

// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/mF2Z5ZPQFi

// Sample program to show how to declare methods and how the Go
// compiler supports them.
package main

import (
	"fmt"
)

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// Notify implements a method that can be called via
// a value of type user.
func (u user) Notify() {
	fmt.Printf("User: Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// ChangeEmail implements a method that can be called via
// a pointer of type user.
func (u *user) ChangeEmail(email string) {
	u.email = email
}

// main is the entry point for the application.
func main() {
	// Value of type user can be used to call the method
	// with a value receiver.
	user1 := user{"Bill", "bill@email.com"}
	user1.Notify()

	// Pointer of type user can also be used to call a method
	// with a value receiver.
	user2 := &user{"Jill", "jill@email.com"}
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

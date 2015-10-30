// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/olva991YF4

// Sample program to show how to declare methods and how the Go
// compiler supports them.
package main

import "fmt"

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method with a value receiver.
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// changeEmail implements a method with a pointer receiver.
func (u *user) changeEmail(email string) {
	u.email = email
}

// main is the entry point for the application.
func main() {
	// Values of type user can be used to call methods
	// declared with a value receiver.
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	// Pointers of type user can also be used to call methods
	// declared with a value receiver.
	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	// Values of type user can be used to call methods
	// declared with a pointer receiver.
	bill.changeEmail("bill@gmail.com")
	bill.notify()

	// Pointers of type user can be used to call methods
	// declared with a pointer receiver.
	lisa.changeEmail("lisa@gmail.com")
	lisa.notify()
}

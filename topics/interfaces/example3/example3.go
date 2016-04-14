// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/_jaCo94x66

// Sample program to show how method sets can affect behavior.
package main

import "fmt"

// user defines a user in the system.
type user struct {
	name  string
	email string
}

// String implements the fmt.Stringer interface.
func (u user) String() string {
	return fmt.Sprintf("My name is %q and name email is %q", u.name, u.email)
}

// main is the entry point for the application.
func main() {

	// Create a value of type user.
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	// Display the value.
	fmt.Println(u)
}

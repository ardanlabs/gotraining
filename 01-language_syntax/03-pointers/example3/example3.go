// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/cK1_GFyDOo

// Sample program to show the basic concept of using a pointer
// to share data.
package main

import "fmt"

// user represents a user in the system.
type user struct {
	name   string
	email  string
	logins int
}

// main is the entry point for the application.
func main() {
	// Declare and initialize a variable named bill of type user.
	bill := user{
		name:  "Bill",
		email: "bill@ardanstudios.com",
	}

	//** We don't need to include all the fields when specifying field
	// names with a composite literal.

	// Pass the "address of" the bill value.
	display(&bill)

	// Pass the "address of" the logins field from within the bill value.
	increment(&bill.logins)

	// Pass the "address of" the bill value.
	display(&bill)
}

// increment declares logins as a pointer variable whose value is
// always an address and points to values of type int.
func increment(logins *int) {
	*logins++
	fmt.Printf("&logins[%p] logins[%p] *logins[%d]\n", &logins, logins, *logins)
}

// display declares u as user pointer variable whose value is always an address
// and points to values of type dog.
func display(u *user) {
	fmt.Printf("%p\t%+v\n", u, *u)
}

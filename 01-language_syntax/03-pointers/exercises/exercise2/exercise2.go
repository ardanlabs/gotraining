// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/GJZXstEkBY

// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

import "fmt"

// user represents a user in the system.
type user struct {
	name        string
	email       string
	accessLevel int
}

// main is the entry point for the application.
func main() {
	bill := user{
		name:        "Bill",
		email:       "bill@ardanstudios.com",
		accessLevel: 1,
	}

	fmt.Println("access:", bill.accessLevel)
	accessLevel(&bill, 10)
	fmt.Println("access:", bill.accessLevel)
}

// accessLevel changes the value of the users access
// level.
func accessLevel(u *user, accessLevel int) {
	u.accessLevel = accessLevel
}

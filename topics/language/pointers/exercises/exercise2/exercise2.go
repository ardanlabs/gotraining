// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

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

func main() {

	// Create a variable of type user and initialize each field.
	bill := user{
		name:        "Bill",
		email:       "bill@ardanlabs.com",
		accessLevel: 1,
	}

	// Display the value of the accessLevel field.
	fmt.Println("access:", bill.accessLevel)

	// Share the bill variable with the accessLevel function
	// along with a value to update the accessLevel field with.
	accessLevel(&bill, 10)

	// Display the value of the accessLevel field again.
	fmt.Println("access:", bill.accessLevel)
}

// accessLevel changes the value of the users access level.
func accessLevel(u *user, accessLevel int) {

	// Set of value of the accessLevel field to the value
	// that is passed in.
	u.accessLevel = accessLevel
}

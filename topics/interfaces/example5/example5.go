// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how the concrete value assigned to
// the interface value is what is stored.
package main

import "fmt"

// printer displays information.
type printer interface {
	print()
}

// user defines a user in the program.
type user struct {
	name string
}

// print displays the user's name.
func (u user) print() {
	fmt.Printf("User Name: %s\n", u.name)
}

// admin defines an admin in the program.
type admin struct {
	name string
}

// print displays the admin's name.
func (a admin) print() {
	fmt.Printf("Admin Name: %s\n", a.name)
}

func main() {

	// Create values of type user and admin.
	u := user{"Bill"}
	a := admin{"Lisa"}

	// Add the values and pointers to the slice of
	// printer interface values.
	entites := []printer{
		// Store copies of the user and admin
		// values in the interface value.
		u,
		a,

		// Store the address of each user and
		// admin value in the interface value.
		&u,
		&a,
	}

	// Change the name field on both values.
	u.name = "Bill_CHG"
	a.name = "Lisa_CHG"

	// Iterate over the slice of entities and see
	// when the change is seen.
	for _, e := range entites {
		e.print()
	}
}

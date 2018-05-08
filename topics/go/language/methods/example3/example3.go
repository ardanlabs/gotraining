// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to declare function variables.
package main

import "fmt"

// data is a struct to bind methods to.
type data struct {
	name string
	age  int
}

// displayName provides a pretty print view of the name.
func (d data) displayName() {
	fmt.Printf("My name is: %s\n", d.name)
}

// setAge sets the age and displays the value.
func (d *data) setAge(age int) {
	d.age = age
	fmt.Printf("%s's is age: %d\n", d.name, d.age)
}

func main() {

	// Declare a variable of type data.
	d := data{
		name: "Bill",
	}

	fmt.Println("\nProper calls to methods")
	fmt.Println("*************************")

	// How we actually call methods in Go.
	d.displayName()
	d.setAge(45)

	fmt.Println("\nWhat the compiler is doing")
	fmt.Println("*************************")

	// This is what Go is doing underneath.
	data.displayName(d)
	(*data).setAge(&d, 45)

	// =========================================================================

	fmt.Println("\nCall Value Receiver Methods with Variable")
	fmt.Println("*************************")

	// Declare a function variable for the method bound to the d variable.
	// The function variable will get its own copy of d because the method
	// is using a value receiver.
	f1 := d.displayName

	// Call the method via the variable.
	f1()

	// Change the value of d.
	d.name = "Joan"

	// Call the method via the variable. We don't see the change.
	f1()

	// =========================================================================

	fmt.Println("\nCall Pointer Receiver Method with Variable")
	fmt.Println("*************************")

	// Declare a function variable for the method bound to the d variable.
	// The function variable will get the address of d because the method
	// is using a pointer receiver.
	f2 := d.setAge

	// Call the method via the variable.
	f2(45)

	// Change the value of d.
	d.name = "Sammy"

	// Call the method via the variable. We see the change.
	f2(45)
}

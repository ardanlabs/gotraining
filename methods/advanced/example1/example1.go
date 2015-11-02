// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/MNI1jR8Ets

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
	fmt.Println("My Name Is: ", d.name)
}

// setAge sets the age and displays the value.
func (d *data) setAge(age int) {
	d.age = age
	fmt.Println("Set Age: ", d.age)
}

// main is the entry point for the application.
func main() {
	// Declare a variable of type data.
	d := data{
		name: "Bill",
	}

	// Declare a function variable for the method
	// bound to the d variable.
	f1 := d.displayName

	// Call the method via the variable.
	f1()

	// Declare a function variable for the function
	// bound to the package.
	f2 := data.displayName

	// Call the function passing the receiver.
	f2(d)

	// Declare a function variable for the method
	// bound to the d variable.
	f3 := d.setAge

	// Call the method via the variable passing the parameter.
	f3(45)

	// Declare a function variable for the function
	// bound to the package. The receiver is a pointer.
	f4 := (*data).setAge

	// Call the function passing the receiver and the parameter.
	f4(&d, 55)
}

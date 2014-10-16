// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/RPyAn3y2OS

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
	// bound to the d varaible.
	f1 := d.displayName

	// Call the method via the variable.
	f1()

	// Declare a function variable for the function
	// bound to the package.
	f2 := data.displayName

	// Call the function passing the reciever.
	f2(d)

	// Declare a function variable for the method
	// bound to the d varaible.
	f3 := d.setAge

	// Call the method via the variable passing the parameter.
	f3(45)

	// Declare a function variable for the function
	// bound to the package. The receiver is a pointer.
	f4 := (*data).setAge

	// Call the function passing the reciever and the parameter.
	f4(&d, 55)
}

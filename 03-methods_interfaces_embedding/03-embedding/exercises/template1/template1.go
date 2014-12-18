// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/C-Wd9jntwb

// Declare a struct type named animal with two fields name and age. Declare a struct
// type named dog with the field bark. Embed the animal type into the dog type. Declare
// and initalize a value of type dog. Display the value of the variable.
//
// Declare a method named yelp to the animal type using a pointer reciever which displays the
// literal string "Not Implemented". Call the method from the value of type dog.
//
// Declare an interface named yelper with a single method called yelp. Declare a value of
// type yelper and assign the address of the value of type dog. Call the method yelp.
//
// Implement the yelper interface for the dog type. Be creative with the
// bark field. Call the method yelp again from the value of type yelper.
package main

import "fmt"

// yelper represents talking animals.
type interface_type_name interface {
	method_name()
}

// animal represents all animals.
type type_name struct {
	field_name type
	field_name type
}

// yelp represents how an animal can speak.
func (receiver_name *type_name) method_name() {
	fmt.Println("Not Implemented")
}

// dog represents a dog.
type type_name struct {
	embedded_type_name
	field_name type_name
}

// yelp represents how an animal can speak.
func (receiver_name *type_name) method_name() {
	for variable_name := 0; variable_name < receiver_name.field_name; variable_name++ {
		fmt.Print("BARK ")
	}
	fmt.Println()
}

// main is the entry point for the application.
func main() {
	// Create a value of type dog.
	variable_name := type_name{
		embedded_type_name: embedded_type_name{
			field_name: value,
			field_name: 1,
		},
		field_name: value,
	}

	// Display the value.
	fmt.Println(variable_name)

	// Use the interface.
	var variable_name2 interface_type_name
	variable_name2 = &variable_name
	variable_name2.method_name()
}

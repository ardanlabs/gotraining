// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/H_9Vm6fkuF

// Declare a named type called counter with a base type of int. Declare and initalize
// a variable of this named type to its zero value. Display the value of this variable
// and the variables type.
//
// Declare a new variable of the named type assign it the value of 10. Display the value.
//
// Declare a variable of the same base type as your named typed. Attempt to assign the
// value of your named type variable to your new base type variable. Does the compiler
// allow the assignment?
package main

import "fmt"

// Counter is a named type for counting.
type type_name type

func main() {
	// Declare a variable of type Counter.
	var variable_name type_name
	fmt.Println(variable_name)

	// Initalize a new variable.
	variable_name := type_name(value)
	fmt.Println(variable_name)

	// Will not compile
	variable_name := value
	variable_name = variable_name
}

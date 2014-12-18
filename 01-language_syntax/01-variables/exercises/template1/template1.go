// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/LL_-2T-6wa

// Declare three variables that are initalized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initalize the variable by
// converting the literal value of Pi (3.14).
package main

import "fmt"

// main is the entry point for the application.
func main() {
	// Declare variables that are set to their zero value.
	var variable_name type

	// Display the value of those variables.
	fmt.Println(variable_name)

	// Declare variables and initalize.
	// Using the short variable declaration operator.
	variable_name := value

	// Display the value of those variables.
	fmt.Println(variable_name)

	// Specify type and perform a conversion.
	variable := type(value)

	// Display the value of the variable.
	fmt.Printf("%T [%v]\n", variable_name, variable_name)
}

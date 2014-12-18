// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/7BOjiq6U2z

// Declare a struct type to maintain information about a user. Declare a function
// that creates value of and returns pointers of this type and an error value. Call
// this function from main and display the value.
//
// Make a second call to your function but this time ignore the value and just test
// the error value.
package main

import "fmt"

// user represents a user in the system.
type type_name struct {
	field_name type
	field_name type
}

// main is the entry point for the application.
func main() {
	// Create a value of type user.
	variable_name, variable_name := function_name()
	if variable_name != nil {
		fmt.Println(variable_name)
		return
	}

	// Display the value.
	fmt.Println(variable_name)

	// Create a value of type user and ignore the second value being returned.
	variable_name, _ := function_name()
	fmt.Println(function_name)
}

// newUser creates and returns pointers of user type values.
func function_name() ([operator]type, error_type) {
	return [operator]type{"value", "value"}, nil
}

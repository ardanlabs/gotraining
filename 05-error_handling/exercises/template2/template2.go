// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/GbWvjprxcc

// Create a custom error type called appError that contains three fields, Err error,
// Message string and Code int. Implement the error interface providing your own message
// using these three fields. Write a function called checkFlag that accepts a bool value.
// If the value is false, return a pointer of your custom error type initialized as you like.
// If the value is true, return a default error. Write a main function to call the
// checkFlag function and check the error for the concrete type.
package main

import (
	"errors"
	"fmt"
)

// appError is a custom error type for the program.
type type_name struct {
	field_name type
	field_name type
	field_name type
}

// Error implements the error interface for appError.
func (receiver_name [operator]type_name) method_name() string {
	return fmt.Sprintf("App Error[%s] Message[%s] Code[%d]", receiver_name.field_name, receiver_name.field_name, receiver_name.field_name)
}

// main is the entry point for the application.
func main() {
	// Call the function to simulate an error of the
	// concrete type.
	variable_name1 := function_name(false)

	// Check the concrete type and handle appropriately.
	switch variable_name2 := variable_name1.(type) {
	case [operator]type_name:
		fmt.Printf("App Error: Code[%d] Message[%s] Err[%s]", variable_name2.field_name, variable_name2.field_name, variable_name2.field_name)
		return
	default:
		fmt.Println(variable_name2)
		return
	}
}

// checkFlag returns one of two errors based on the value of the parameter.
func function_name(parameter_name type) error_type {
	// If the parameter is false return an appError.
	if !parameter_name {
		return [operator]type_name{errors.New("error message"), "error message", N}
	}

	// Return a default error.
	return errors.New("error message")
}

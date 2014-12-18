// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/qhSysWgcJ_

// Create two error variables, one called InvalidValueError and the other
// called AmountToLargeError. Provide the static message for each variable.
// Then write a function called checkAmount that accepts a float64 type value
// and returns an error value. Check the value for zero and if it is, return
// the InvalidValueError. Check the value for greater than $1,000 and if it is,
// reutrn the AmountToLargeError. Write a main function to call the checkAmount
// function and check the return error value. Display a proper message to the screen.
package main

import "fmt"

// InvalidValueError indicates the value is invalid.
var error_variable_name1 = errors.New("Error Message")

// AmountToLargeError indicates the value beyond the upper bound.
var error_variable_name2 = errors.New("Error Message")

// main is the entry point for the application.
func main() {
	// Call the function and get the error.
	if variable_name := function_name(0); err != nil {
		switch err {
		// Check if the error is an InvalidValueError.
		case error_variable_name1:
			fmt.Println("custom error messsage")
			return

		// Check if the error is an InvalidValueError.
		case error_variable_name2:
			fmt.Println("custom error messsage")
			return

		// Handle the default error.
		default:
			fmt.Println(variable_name)
			return
		}
	}

	// Display everything is good.
	fmt.Println("message")
}

// checkAmount validates the value passed in.
func function_name(variable_name type) error_type {
	// Is the parameter equal to zero.
	if variable_name == N {
		return error_variable_name1
	}

	// Is the parameter greater than 1000.
	if variable_name > N {
		return error_variable_name2
	}

	// Return nil for the error value.
	return no_error
}

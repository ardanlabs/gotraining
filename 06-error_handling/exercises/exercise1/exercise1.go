// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/cIVJqLzm4d

// Create two error variables, one called InvalidValueError and the other
// called AmountToLargeError. Provide the static message for each variable.
// Then write a function called checkAmount that accepts a float64 type value
// and returns an error value. Check the value for zero and if it is, return
// the InvalidValueError. Check the value for greater than $1,000 and if it is,
// reutrn the AmountToLargeError. Write a main function to call the checkAmount
// function and check the return error value. Display a proper message to the screen.
package main

import (
	"errors"
	"fmt"
)

// InvalidValueError indicates the value is invalid.
var InvalidValueError = errors.New("Invalid Value")

// AmountToLargeError indicates the value beyond the upper bound.
var AmountTooLargeError = errors.New("Amount To Large")

// main is the entry point for the application.
func main() {
	// Call the function and get the error.
	if err := checkAmount(0); err != nil {
		switch err {
		// Check if the error is an InvalidValueError.
		case InvalidValueError:
			fmt.Println("Value provided is not valid.")
			return

		// Check if the error is an InvalidValueError.
		case AmountTooLargeError:

			fmt.Println("Value provided is too large.")
			return

		// Handle the default error.
		default:
			fmt.Println(err)
			return
		}
	}

	// Display everything is good.
	fmt.Println("Everything checks out.")
}

// checkAmount validates the value passed in.
func checkAmount(f float64) error {

	switch {
	// Is the parameter equal to zero.
	case f == 0:
		return InvalidValueError
	// Is the parameter greater than 1000.
	case f > 1000:
		return AmountTooLargeError
	// Return nil for the error value.
	default:
		return nil
	}

}

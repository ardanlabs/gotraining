// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/TczNj28oNf

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
var AmountToLargeError = errors.New("Amount To Large")

// main is the entry point for the application.
func main() {
	if err := checkAmount(0); err != nil {
		switch err {
		case InvalidValueError:
			fmt.Println("Value provided is not valid.")
			return

		case AmountToLargeError:
			fmt.Println("Value provided is too large.")
			return

		default:
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Everything checks out.")
}

// checkAmount validates the value passed in.
func checkAmount(f float64) error {
	if f == 0 {
		return InvalidValueError
	}

	if f > 1000 {
		return AmountToLargeError
	}

	return nil
}

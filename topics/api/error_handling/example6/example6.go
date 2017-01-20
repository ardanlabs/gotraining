// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how wrapping errors work.
package main

import (
	"fmt"

	"github.com/pkg/errors"
)

// AppError represents a custom error type.
type AppError struct {
	State int
}

// Error implements the error interface.
func (c *AppError) Error() string {
	return fmt.Sprintf("App Error, State: %d", c.State)
}

func main() {

	// Make the function call and validate the error.
	if err := firstCall(); err != nil {

		// Use type as context to determine cause.
		switch v := errors.Cause(err).(type) {
		case *AppError:

			// We got our custom error type.
			fmt.Println("Custom App Error:", v.State)

		default:

			// We did not get any specific error type.
			fmt.Println("Default Error")
		}

		// Display the stack trace for the error.
		fmt.Println("\nStack Trace\n********************************")
		fmt.Printf("%+v", err)
	}
}

// firstCall makes a call to a second function and wraps any error.
func firstCall() error {
	return errors.Wrap(secondCall(), "attempted firstCall")
}

// secondCall makes a call to a third function and wraps any error.
func secondCall() error {
	return errors.Wrap(thirdCall(), "attempted doSomething")
}

// thirdCall create an error value we will validate.
func thirdCall() error {
	return &AppError{99}
}

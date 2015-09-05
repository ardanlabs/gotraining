// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/rLCuGVzwy4

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
type appError struct {
	Err     error
	Message string
	Code    int
}

// Error implements the error interface for appError.
func (a *appError) Error() string {
	return fmt.Sprintf("App Error[%s] Message[%s] Code[%d]", a.Err, a.Message, a.Code)
}

// main is the entry point for the application.
func main() {
	// Call the function to simulate an error of the
	// concrete type.
	err := checkFlag(false)

	// Check the concrete type and handle appropriately.
	switch e := err.(type) {
	case *appError:
		fmt.Printf("App Error: Code[%d] Message[%s] Err[%s]", e.Code, e.Message, e.Err)
		return
	default:
		fmt.Println(e)
		return
	}
}

// checkFlag returns one of two errors based on the value of the parameter.
func checkFlag(t bool) error {
	// If the parameter is false return an appError.
	if !t {
		return &appError{errors.New("Flag False"), "The Flag was false", 9}
	}

	// Return a default error.
	return errors.New("Flag True")
}

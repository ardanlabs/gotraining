// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Create a custom error type called appError that contains three fields, err error,
// message string and code int. Implement the error interface providing your own message
// using these three fields. Implement a second method named temporary that returns false
// when the value of the code field is 9. Write a function called checkFlag that accepts
// a bool value. If the value is false, return a pointer of your custom error type
// initialized as you like. If the value is true, return a default error. Write a main
// function to call the checkFlag function and check the error using the temporary
// interface.
package main

import (
	"errors"
	"fmt"
)

// appError is a custom error type for the program.
type appError struct {
	err     error
	message string
	code    int
}

// Error implements the error interface for appError.
func (a *appError) Error() string {
	return fmt.Sprintf("App Error[%s] Message[%s] Code[%d]", a.err, a.message, a.code)
}

// Temporary implements behavior about the error.
func (a *appError) Temporary() bool {
	return (a.code != 9)
}

// temporary is used to test the error we receive.
type temporary interface {
	Temporary() bool
}

func main() {
	if err := checkFlag(false); err != nil {
		switch e := err.(type) {
		case temporary:
			fmt.Println(err)
			if !e.Temporary() {
				fmt.Println("Critical Error!")
			}
		default:
			fmt.Println(err)
		}
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

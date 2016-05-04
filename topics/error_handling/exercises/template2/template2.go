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

// Add imports.

// Declare a struct type named appError with three fields, err of type error,
// message of type string and code of type int.

// Declare a method for the appError struct type that implements the
// error interface.

// Declare a method for the appError struct type named Temporary that returns
// true when the value of the code field is not 9.

// Declare the temporary interface type with a method named Temporary that
// takes no parameters and returns a bool.

// Declare a function named checkFlag that accepts a boolean value and
// returns an error interface value.
func checkFlag( /* parameter */ ) /* return arg */ {

	// If the parameter is false return an appError.

	// Return a default error.
}

func main() {

	// Call the checkFlag function to simulate an error of the
	// concrete type.

	// Check the concrete type and handle appropriately.
	switch e := err.(type) {

	// Apply the case for the existence of the Temporary behavior.
	// Log the error and write a second message only if the
	// error is not temporary.

	// Apply the default case and just log the error.
	}
}

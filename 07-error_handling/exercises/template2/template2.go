// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/x6UimVQMMQ

// Create a custom error type called appError that contains three fields, Err error,
// Message string and Code int. Implement the error interface providing your own message
// using these three fields. Write a function called checkFlag that accepts a bool value.
// If the value is false, return a pointer of your custom error type initialized as you like.
// If the value is true, return a default error. Write a main function to call the
// checkFlag function and check the error for the concrete type.
package main

// Add imports.

// Declare a struct type named appError with three fields, Err of type error,
// Message of type string and Code of type int.

// Declare a method for the appError struct type that implements the
// error interface.

// Declare a function named checkFlag that accepts a boolean value and
// returns an error interface value.
func checkFlag( /* parameter */ ) /* return arg */ {
	// If the parameter is false return an appError.

	// Return a default error.
}

// main is the entry point for the application.
func main() {
	// Call the checkFlag function to simulate an error of the
	// concrete type.

	// Check the concrete type and handle appropriately.
	switch e := err.(type) {
	// Apply the case for the default error type.

	// Apply the default case.
	}
}

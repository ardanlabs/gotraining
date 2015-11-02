// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/QiGOu7b32S

// Sample program to show how to trap panics that can occur from the Go runtime.
package main

import (
	"fmt"
	"runtime"
)

// main is the entry point for the application.
func main() {
	// Call the testPanic function to run the test.
	if err := testPanic(); err != nil {
		fmt.Println("Error:", err)
	}
}

// testPanic simulates a function that encounters a panic to
// test our catchPanic function.
func testPanic() (err error) {
	// Schedule the catchPanic function to be called when
	// the testPanic function returns.
	defer catchPanic(&err)

	fmt.Println("Start Test")

	// Mimic a traditional error from a function.
	err = mimicError("1")

	// Trying to dereference a nil pointer will cause the
	// runtime to panic.
	var p *int
	*p = 10

	fmt.Println("End Test")
	return err
}

// catchPanic catches panics and processes the error.
func catchPanic(err *error) {
	// Check if a panic occurred.
	if r := recover(); r != nil {
		fmt.Println("PANIC Deferred")

		// Capture the stack trace
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)
		fmt.Println("Stack Trace:", string(buf))

		// If the caller wants the error back provide it.
		if err != nil {
			*err = fmt.Errorf("%v", r)
		}
	}
}

// mimicError is a function that simulates an error for
// testing the code.
func mimicError(key string) error {
	return fmt.Errorf("Mimic Error : %s", key)
}

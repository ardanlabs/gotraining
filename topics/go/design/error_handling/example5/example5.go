// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show see if the class can find the bug.
package main

import "log"

// customError is just an empty struct.
type customError struct{}

// Error implements the error interface.
func (c *customError) Error() string {
	return "Find the bug."
}

// fail returns nil values for both return types.
func fail() ([]byte, *customError) {
	return nil, nil
}

func main() {
	var err error
	if _, err = fail(); err != nil {
		log.Fatal("Why did this fail?")
	}

	log.Println("No Error")
}

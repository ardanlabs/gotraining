// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/czXpjvWWTT

// Sample program to show see if the class can find the bug.
package main

import (
	"log"
)

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

// main is the entry point for the application.
func main() {
	var err error
	if _, err = fail(); err != nil {
		log.Fatal("Why did this fail?")
	}

	log.Println("No Error")
}

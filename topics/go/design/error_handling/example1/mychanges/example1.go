// Ruben: These changes were made in the video of safaribooks but not in this code, so I added it.

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how the default error type is implemented.
package main

import "fmt"

// http://golang.org/pkg/builtin/#error
type error interface {
	Error() string
}

// http://golang.org/src/pkg/errors/errors.go
type errorString struct {
	s string
}

// http://golang.org/src/pkg/errors/errors.go
func (e errorString) Error() string {
	return e.s
}

// http://golang.org/src/pkg/errors/errors.go
// New returns an error that formats as the given text.
func New(text string) error {
	return errorString{text}
}

func main() {
	// err == errMine
	// In this example we updated it to use value semantics
	// and if you use value semantic then what is compared are the values,
	// so this is making a basic string comparison "Bad Request" == "Bad Request",
	// if you use pointer semantics as we have it originally,
	// it will NOT be the same because '==' will compare addresses.

	if err := webCall(); err == errMine {
		fmt.Println(err, "SAME")
		return
	}

	fmt.Println("Life is good")
}

var errMine = New("Bad Request")

// webCall performs a web operation.
func webCall() error {
	return New("Bad Request")
}

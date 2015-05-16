// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// https://play.golang.org/p/4YHAbpynl3

// Sample program to show how to use error variables to help the
// caller determine the exact error being returned.
package main

import (
	"errors"
	"fmt"
)

var (
	// ErrBadRequest is returned when there are problems with the request.
	ErrBadRequest = errors.New("Bad Request")

	// ErrMovedPermanently is returned when a 301/302 is returned.
	ErrMovedPermanently = errors.New("Moved Permanently")
)

// main is the entry point for the application.
func main() {
	if err := webCall(); err != nil {
		switch err {
		case ErrBadRequest:
			fmt.Println("Bad Request Occurred")
			return

		case ErrMovedPermanently:
			fmt.Println("The URL moved, check it again")
			return

		default:
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Life is good")
}

// webCall performs a web operation.
func webCall() error {
	return ErrBadRequest
}

// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/Zo1o0LFThp

// Sample program to show how to use error variables to help the
// caller determine the exact error being returned.
package main

import (
	"errors"
	"fmt"
)

var (
	// BadRequestError is returned when there are problems with the request.
	BadRequestError = errors.New("Bad Request")

	// MovedPermanentlyError is returned when a 301/302 is returned.
	MovedPermanentlyError = errors.New("Moved Permanently")
)

// main is the entry point for the application.
func main() {
	if err := webCall(); err != nil {
		switch err {
		case BadRequestError:
			fmt.Println("Bad Request Occurred")
			return

		case MovedPermanentlyError:
			fmt.Println("The URL moved, check it agaib")
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
	return BadRequestError
}

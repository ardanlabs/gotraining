// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/FRnwmQx_ZI

// Sample program to show how to use error variables to help the
// caller determine the exact error being returned.
package main

import (
	"errors"
	"fmt"
)

// ErrBadRequest is returned when there are problems with the request.
var ErrBadRequest = errors.New("Bad Request")

// ErrMovedPermanently is returned when a 301/302 is returned.
var ErrMovedPermanently = errors.New("Moved Permanently")

// main is the entry point for the application.
func main() {
	if err := webCall(true); err != nil {
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
func webCall(b bool) error {
	if b {
		return ErrBadRequest
	}

	return ErrMovedPermanently
}

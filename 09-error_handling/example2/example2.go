// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/kKVIMMDpjb

// Sample program to show how to use error variables to help the
// caller determine the exact error being returned.
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
func (e *errorString) Error() string {
	return e.s
}

// http://golang.org/src/pkg/errors/errors.go
// New returns an error that formats as the given text.
func New(text string) error {
	return &errorString{text}
}

var (
	// BadRequestError is returned when there are problems with the request.
	BadRequestError = New("Bad Request")

	// MovedPermanentlyError is returned when a 301/302 is returned.
	MovedPermanentlyError = New("Moved Permanently")
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

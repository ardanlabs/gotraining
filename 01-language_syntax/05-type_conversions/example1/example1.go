// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/mhKlxSyuxr

// Sample program to show how to declare and use a named type.
package main

import "fmt"

// Duration is a named type that represents a duration
// of time in Nanosecond.
type Duration int64

func main() {
	// Declare a variable of type Duration.
	var duration Duration
	fmt.Println(duration)

	// Declare a variable of type int64 and assign a value.
	nanosecond := int64(10)

	// Attemped to assign a variable of type int64 (base type of Duration) to
	// a variable of type Duration.
	// duration = nanosecond

	// ./example1.go:22: cannot use nanosecond (type int64) as type Duration in assignment

	// Convert a value of type int64 to type Duration.
	duration = Duration(nanosecond)
	fmt.Println(duration)
}

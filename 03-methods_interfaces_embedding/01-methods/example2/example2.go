// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/EMY2xb1csT

// Sample program to show how to declare methods against
// a named type.
package main

import (
	"fmt"
)

// duration is a named type that represents a duration
// of time in Nanosecond.
type duration int64

// SetSeconds can change the value of duration type variables.
func (d *duration) SetSeconds(seconds duration) {
	*d = 1e9 * seconds
}

// Seconds returns a formatted string of duration in seconds.
func (d duration) Seconds() string {
	seconds := d / 1e9
	return fmt.Sprintf("%d Seconds", seconds)
}

// main is the entry point for the application.
func main() {
	// Declare a variable of type duration set to
	// its zero value.
	var dur duration

	// Change the value of dur to equal
	// five seconds.
	dur.SetSeconds(5)

	// Display the new value of dur.
	fmt.Println(dur.Seconds())
}

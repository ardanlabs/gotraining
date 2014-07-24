// Sample program to show how to create methods against
// a named type.
package main

import (
	"fmt"
)

// Duration is a named type that represents a duration
// of time in Nanosecond.
type Duration int64

// SetSeconds can change the value of Duration type variables.
func (d *Duration) SetSeconds(seconds Duration) {
	*d = 1e9 * seconds
}

// Seconds returns a formatted string of duration in seconds.
func (d Duration) Seconds() string {
	seconds := d / 1e9
	return fmt.Sprintf("%d Seconds", seconds)
}

// main is the entry point for the application.
func main() {
	// Declare a variable of type Duration set to
	// its zero value.
	var duration Duration

	// Change the value of duration to equal
	// five seconds.
	duration.SetSeconds(5)

	// Display the new value of duration.
	fmt.Println(duration.Seconds())
}

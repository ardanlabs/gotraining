// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to declare methods against
// a named type.
package main

import "fmt"

// duration is a named type that represents a duration
// of time in Nanosecond.
type duration int64

const (
	nanosecond  duration = 1
	microsecond          = 1000 * nanosecond
	millisecond          = 1000 * microsecond
	second               = 1000 * millisecond
	minute               = 60 * second
	hour                 = 60 * minute
)

// setHours sets the specified number of hours.
func (d *duration) setHours(h float64) {
	*d = duration(h) * hour
}

// hours returns the duration as a floating point number of hours.
func (d duration) hours() float64 {
	hour := d / hour
	nsec := d % hour
	return float64(hour) + float64(nsec)*(1e-9/60/60)
}

func main() {

	// Declare a variable of type duration set to
	// its zero value.
	var dur duration

	// Change the value of dur to equal
	// five hours.
	dur.setHours(5)

	// Display the new value of dur.
	fmt.Println("Hours:", dur.hours())
}

// http://golang.org/pkg/time/

/*
// A Duration represents the elapsed time between two instants as
// an int64 nanosecond count. The representation limits the largest
// representable duration to approximately 290 years.

type Duration int64

// Common durations. There is no definition for units of Day or larger
// to avoid confusion across daylight savings time zone transitions.

const (
        Nanosecond  Duration = 1
        Microsecond          = 1000 * Nanosecond
        Millisecond          = 1000 * Microsecond
        Second               = 1000 * Millisecond
        Minute               = 60 * Second
        Hour                 = 60 * Minute
)

// Add returns the time t+d.
func (t Time) Add(d Duration) Time
*/

package main

import (
	"fmt"
	"time"
)

// time.Duration(5) * time.Second
const fiveSeconds = 5 * time.Second

func main() {
	now := time.Now()
	lessFiveNanoseconds := now.Add(-5)
	lessFiveSeconds := now.Add(-fiveSeconds)

	fmt.Printf("Now     : %v\n", now)
	fmt.Printf("Nano    : %v\n", lessFiveNanoseconds)
	fmt.Printf("Seconds : %v\n", lessFiveSeconds)
}

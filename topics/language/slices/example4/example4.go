// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to grow a slice using the built-in function append
// and how append grows the capacity of the underlying array.
package main

import "fmt"

func main() {

	// Declare a nil slice of strings.
	var data []string

	// Capture the capacity of the slice.
	lastCap := cap(data)

	// Append ~10k strings to the slice.
	for record := 1; record <= 102400; record++ {

		// Use the built-in function append to add to the slice.
		data = append(data, fmt.Sprintf("Rec: %d", record))

		// When the capacity of the slice changes, display the changes.
		if lastCap != cap(data) {

			// Calculate the percent of change.
			capChg := float64(cap(data)-lastCap) / float64(lastCap) * 100

			// Save the new values for capacity.
			lastCap = cap(data)

			// Display the results.
			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n",
				&data[0],
				record,
				cap(data),
				capChg)
		}
	}
}

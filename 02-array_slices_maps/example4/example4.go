// Sample program to show how to grow a slice using the built-in function append
// and how append grows the capacity of the underlying array.
package main

import (
	"fmt"
)

// main is the entry point for the application.
func main() {
	// Declare a nil slice of strings.
	var data []string

	// Capture the length and capacity of the slice.
	lastLen := len(data)
	lastCap := cap(data)

	// Append ~10k strings to the slice.
	for record := 1; record <= 10240; record++ {
		// Use the built-in function append to add to the slice.
		data = append(data, fmt.Sprintf("Rec: %d", record))

		// when the capacity of the slice changes, display the changes.
		if lastCap != cap(data) {
			// Calculate the percent of change.
			lenChg := float64(len(data)-lastLen) / float64(lastLen) * 100
			capChg := float64(cap(data)-lastCap) / float64(lastCap) * 100

			// Save the new values for length and capacity.
			lastLen = len(data)
			lastCap = cap(data)

			// Display the results.
			fmt.Printf("Addr[%p]\tIndex[%d]\t\tLen[%d - %2.f%%]\t\tCap[%d - %2.f%%]\n",
				&data[0],
				record,
				len(data),
				lenChg,
				cap(data),
				capChg)
		}
	}
}

/*
Addr[0x208178180]   Index[1]        Len[1 - +Inf%]      Cap[1 - +Inf%]
Addr[0x2081b2020]   Index[2]        Len[2 - 100%]       Cap[2 - 100%]
Addr[0x2081ae080]   Index[3]        Len[3 - 50%]        Cap[4 - 100%]
Addr[0x2081b8000]   Index[5]        Len[5 - 67%]        Cap[8 - 100%]
Addr[0x2081ba000]   Index[9]        Len[9 - 80%]        Cap[16 - 100%]
Addr[0x208186200]   Index[17]       Len[17 - 89%]       Cap[32 - 100%]
Addr[0x2081bc000]   Index[33]       Len[33 - 94%]       Cap[64 - 100%]
Addr[0x2081be000]   Index[65]       Len[65 - 97%]       Cap[128 - 100%]
Addr[0x20818d000]   Index[129]      Len[129 - 98%]      Cap[256 - 100%]
Addr[0x2081c0000]   Index[257]      Len[257 - 99%]      Cap[512 - 100%]
Addr[0x2081c2000]   Index[513]      Len[513 - 100%]     Cap[1024 - 100%]
Addr[0x2081c8000]   Index[1025]     Len[1025 - 100%]    Cap[1280 - 25%]
Addr[0x2081d2000]   Index[1281]     Len[1281 - 25%]     Cap[1776 - 39%]
Addr[0x2081e2000]   Index[1777]     Len[1777 - 39%]     Cap[2560 - 44%]
Addr[0x2081ec000]   Index[2561]     Len[2561 - 44%]     Cap[3584 - 40%]
Addr[0x2081c8000]   Index[3585]     Len[3585 - 40%]     Cap[4608 - 29%]
Addr[0x2081fa000]   Index[4609]     Len[4609 - 29%]     Cap[6144 - 33%]
Addr[0x208212000]   Index[6145]     Len[6145 - 33%]     Cap[7680 - 25%]
Addr[0x208230000]   Index[7681]     Len[7681 - 25%]     Cap[9728 - 27%]
Addr[0x2081e2000]   Index[9729]     Len[9729 - 27%]     Cap[12288 - 26%]
*/

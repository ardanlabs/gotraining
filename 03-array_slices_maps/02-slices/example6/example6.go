// http://play.golang.org/p/zN7btAZFV1

// Sample program to show how to use a third index slice.
package main

import (
	"fmt"
	"unsafe"
)

// main is the entry point for the application.
func main() {
	// Create a slice of strings with different types of fruit.
	slice := []string{"Apple", "Orange", "Banana", "Grape", "Plum"}
	inspectSlice(slice)

	// Take a slice from element two for length 3 - 2
	// with a capacity of 5 - 3
	takeOne := slice[2:3]
	inspectSlice(takeOne)

	// For slice[ i : j : k ] the
	// Length:   j - i
	// Capacity: k - i

	// Take a slice from element two for length 3-2
	// with a capacity of 3 - 2
	takeOneCapOne := slice[2:3:3] // Use the third index position to
	inspectSlice(takeOneCapOne)   // set the capacity.

	// Append a new element which will create a new
	// underlying array to increase capacity.
	takeOneCapOne = append(takeOneCapOne, "Kiwi")
	inspectSlice(takeOneCapOne)
}

// inspectSlice exposes the slice header for review.
func inspectSlice(slice []string) {
	// Capture the address to the slice structure.
	address := unsafe.Pointer(&slice)

	// Capture the address where the length and cap size is stored.
	lenAddr := uintptr(address) + uintptr(8)
	capAddr := uintptr(address) + uintptr(16)

	// Create pointers to the length and cap size.
	lenPtr := (*int)(unsafe.Pointer(lenAddr))
	capPtr := (*int)(unsafe.Pointer(capAddr))

	// Create a pointer to the underlying array.
	addPtr := (*[8]string)(unsafe.Pointer(*(*uintptr)(address)))

	fmt.Printf("Slice Addr[%p] Len Addr[0x%x] Cap Addr[0x%x]\n",
		address,
		lenAddr,
		capAddr)

	fmt.Printf("Slice Length[%d] Cap[%d]\n",
		*lenPtr,
		*capPtr)

	for index := 0; index < *lenPtr; index++ {
		fmt.Printf("[%d] %p %s\n",
			index,
			&(*addPtr)[index],
			(*addPtr)[index])
	}

	fmt.Printf("\n\n")
}

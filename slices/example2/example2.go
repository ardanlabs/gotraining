// Example proves a slice is a reference type that contains a three
// word header with a pointer, the length and capacity of the
// underlying array.
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	slice := make([]string, 5, 8)
	slice[0] = "Apple"
	slice[1] = "Orange"
	slice[2] = "Banana"
	slice[3] = "Grape"
	slice[4] = "Plum"

	InspectSlice(slice)
}

func InspectSlice(slice []string) {
	// Capture the address to the slice structure
	address := unsafe.Pointer(&slice)

	// Capture the address where the length and cap size is stored
	lenAddr := uintptr(address) + uintptr(8)
	capAddr := uintptr(address) + uintptr(16)

	// Create pointers to the length and cap size
	lenPtr := (*int)(unsafe.Pointer(lenAddr))
	capPtr := (*int)(unsafe.Pointer(capAddr))

	// Create a pointer to the underlying array
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

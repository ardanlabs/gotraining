// Example shows the practical use of using a third index slice.
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	slice := []string{"Apple", "Orange", "Banana", "Grape", "Plum"}
	InspectSlice(slice)

	takeOne := slice[2:3]
	InspectSlice(takeOne)

	// For slice[ i : j : k ] the
	// Length:   j - i
	// Capacity: k - i

	takeOneCapOne := slice[2:3:3] // Use the third index position to
	InspectSlice(takeOneCapOne)   // set the capacity

	takeOneCapOne = append(takeOneCapOne, "Kiwi")
	InspectSlice(takeOneCapOne)
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

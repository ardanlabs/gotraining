// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/IIubKW34GA

// Sample program to show how the backing array for a referene type can
// be placed contiguous in memory with the header value. Must be run locally
// and not in the playground.
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	bs := []byte("HELLO")
	s := string(bs)

	// Capture the address to the slice structure
	address := unsafe.Pointer(&s)

	// Capture the address where the length
	lenAddr := uintptr(address) + uintptr(8)

	// Create pointers to the length
	lenPtr := (*int)(unsafe.Pointer(lenAddr))

	// Create a pointer to the underlying array
	arrPtr := (*[5]byte)(unsafe.Pointer(*(*uintptr)(address)))

	fmt.Printf("String Addr[%p] Len Addr[0x%x] Arr Addr[%p]\n",
		address,
		lenAddr,
		arrPtr)

	fmt.Printf("String Length[%d]\n",
		*lenPtr)

	for index := 0; index < *lenPtr; index++ {
		fmt.Printf("[%d] %p %c\n",
			index,
			&(*arrPtr)[index],
			(*arrPtr)[index])
	}

	fmt.Printf("\n\n")
}

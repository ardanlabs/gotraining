// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/hEmKF_Mbsf

// Alignment is about placing fields on address alignment boundaries
// for more efficient reads and writes to memory.

// Sample program to show how struct types align on boundaries.
package main

import (
	"fmt"
	"unsafe"
)

// main is the entry point for the application.
func main() {
	// Since all the fields are based on a single byte,
	// the all align perfectly within a single word.
	var np struct {
		a bool
		b bool
		c bool
	}
	size := unsafe.Sizeof(np)
	fmt.Printf("0 bytes of Padding - SizeOf[%d][%p %p %p]\n", size, &np.a, &np.b, &np.c)

	// Since the second field is a 2 byte int, that must
	// be aligned properly start at an address ending in
	// 0, 2, 4, 6, 8, A, C, E.
	// 1 byte of padding is included to align the int16 properly.
	var sbp struct {
		a bool
		b int16
	}
	size = unsafe.Sizeof(sbp)
	fmt.Printf("1 byte of Padding - SizeOf[%d][%p %p]\n", size, &sbp.a, &sbp.b)

	// Since the second field is a 4 byte int, that must
	// be aligned properly start at an address ending in
	// 0, 4, 8, C.
	// 3 byte of padding is included to align the int32 properly.
	var tbp struct {
		a bool
		b int32
	}
	size = unsafe.Sizeof(tbp)
	fmt.Printf("3 byte of Padding - SizeOf[%d][%p %p]\n", size, &tbp.a, &tbp.b)

	// Since the second field is a 8 byte int, that must
	// be aligned properly start at an address ending in
	// 0, 8.
	// 4 byte of padding is included to align the int64 properly.
	var fbp struct {
		a bool
		b int64
	}
	size = unsafe.Sizeof(fbp)
	fmt.Printf("7 byte of Padding - SizeOf[%d][%p %p]\n", size, &fbp.a, &fbp.b)
}

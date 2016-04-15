// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go get honnef.co/go/structlayout/cmd/...

// Alignment is about placing fields on address alignment boundaries
// for more efficient reads and writes to memory.

// Sample program to show how struct types align on boundaries.
package main

import (
	"fmt"
	"unsafe"
)

// No padding.
type np struct {
	a bool
	b bool
	c bool
}

// Single byte padding.
type sbp struct {
	a bool
	b int16
}

// Two byte padding.
type tbp struct {
	a bool
	b int32
}

// Four byte padding.
type fbp struct {
	a bool
	b int64
}

func main() {

	// No padding.
	// type np struct {
	// 	a bool
	// 	b bool
	// 	c bool
	// }
	//
	// Since all the fields are based on a single byte,
	// the all align perfectly within a single word.
	//
	// structlayout -json github.com/ardanlabs/gotraining/topics/struct_types/advanced/example1 np | structlayout-pretty
	// 	np.a bool: 0-1 (size 1, align 1)
	// 	np.b bool: 1-2 (size 1, align 1)
	// 	np.c bool: 2-3 (size 1, align 1)
	//
	// 	  +--------+
	// 	0 |        | <- np.a bool
	// 	  +--------+
	// 	1 |        | <- np.b bool
	// 	  +--------+
	// 	2 |        | <- np.c bool
	// 	  +--------+
	//
	var np np
	size := unsafe.Sizeof(np)
	fmt.Printf("0 bytes of Padding - SizeOf[%d][%p %p %p]\n", size, &np.a, &np.b, &np.c)

	// =========================================================================

	// Single byte padding.
	// type sbp struct {
	// 	a bool
	// 	b int16
	// }
	//
	// Since the second field is a 2 byte int, that must
	// be aligned properly start at an address ending in
	// 0, 2, 4, 6, 8, A, C, E.
	// 1 byte of padding is included to align the int16 properly.
	//
	// structlayout -json github.com/ardanlabs/gotraining/topics/struct_types/advanced/example1 sbp | structlayout-pretty
	// 	sbp.a bool:  0-1 (size 1, align 1)
	// 	padding:     1-2 (size 1, align 0)
	// 	sbp.b int16: 2-4 (size 2, align 2)
	//
	// 	  +--------+
	// 	0 |        | <- sbp.a bool
	// 	  +--------+
	// 	1 |        | <- padding
	// 	  +--------+
	// 	2 |        | <- sbp.b int16
	// 	  +--------+
	// 	3 |        |
	// 	  +--------+
	//
	var sbp sbp
	size = unsafe.Sizeof(sbp)
	fmt.Printf("1 byte of Padding - SizeOf[%d][%p %p]\n", size, &sbp.a, &sbp.b)

	// =========================================================================

	// Two byte padding.
	// type tbp struct {
	// 	a bool
	// 	b int32
	// }
	//
	// Since the second field is a 4 byte int, that must
	// be aligned properly start at an address ending in
	// 0, 4, 8, C.
	// 3 byte of padding is included to align the int32 properly.
	//
	// structlayout -json github.com/ardanlabs/gotraining/topics/struct_types/advanced/example1 tbp | structlayout-pretty
	// 	tbp.a bool:  0-1 (size 1, align 1)
	// 	padding:     1-4 (size 3, align 0)
	// 	tbp.b int32: 4-8 (size 4, align 4)
	//
	// 	  +--------+
	// 	0 |        | <- tbp.a bool
	// 	  +--------+
	// 	1 |        | <- padding
	// 	  +--------+
	// 	  -........-
	// 	  +--------+
	// 	3 |        | <- padding
	// 	  +--------+
	// 	4 |        | <- tbp.b int32
	// 	  +--------+
	// 	  -........-
	// 	  +--------+
	// 	7 |        |
	// 	  +--------+
	//
	var tbp tbp
	size = unsafe.Sizeof(tbp)
	fmt.Printf("3 byte of Padding - SizeOf[%d][%p %p]\n", size, &tbp.a, &tbp.b)

	// =========================================================================

	// Four byte padding.
	// type fbp struct {
	// 	a bool
	// 	b int64
	// }
	//
	// Since the second field is a 8 byte int, that must
	// be aligned properly start at an address ending in
	// 0, 8.
	// 4 byte of padding is included to align the int64 properly.
	//
	// structlayout -json github.com/ardanlabs/gotraining/topics/struct_types/advanced/example1 fbp | structlayout-pretty
	// 	fbp.a bool:  0-1  (size 1, align 1)
	// 	padding:     1-8  (size 7, align 0)
	// 	fbp.b int64: 8-16 (size 8, align 8)
	//
	//    +--------+
	//  0 |        | <- fbp.a bool
	//    +--------+
	//  1 |        | <- padding
	//    +--------+
	//    -........-
	//    +--------+
	//  7 |        | <- padding
	//    +--------+
	//  8 |        | <- fbp.b int64
	//    +--------+
	//    -........-
	//    +--------+
	// 15 |        |
	//    +--------+
	//
	var fbp fbp
	size = unsafe.Sizeof(fbp)
	fmt.Printf("7 byte of Padding - SizeOf[%d][%p %p]\n", size, &fbp.a, &fbp.b)
}

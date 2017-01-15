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

// No byte padding.
type nbp struct {
	a bool // 	1 byte				sizeof 1
	b bool // 	1 byte				sizeof 2
	c bool // 	1 byte				sizeof 3 - Aligned on 1 byte
}

// Single byte padding.
type sbp struct {
	a bool //	1 byte				sizeof 1
	//			1 byte padding		sizeof 2
	b int16 // 	2 bytes				sizeof 4 - Aligned on 2 bytes
}

// Two byte padding.
type tbp struct {
	a bool //	1 byte				size 1
	//			3 bytes padding		size 4
	b int32 //	4 bytes				size 8 - Aligned on 2 bytes
}

// Four byte padding.
type fbp struct {
	a bool //	1 byte				size 1
	//			7 bytes padding		size 8
	b int64 //	8 bytes				size 16 - Aligned on 8 bytes
}

// Eight byte padding on 64bit Arch. Word size is 8 bytes.
type ebp64 struct {
	a string //	16 bytes			size 16
	b int32  //	 4 bytes			size 20
	//  		 4 bytes padding	size 24
	c string //	16 bytes			size 40
	d int32  //	 4 bytes			size 44
	//  		 4 bytes padding	size 48 - Aligned on 8 bytes
}

// No padding on 32bit Arch. Word size is 4 bytes.
// To see this build as 32 bit: GOARCH=386 go build
type ebp32 struct {
	a string //	 8 bytes			size  8
	b int32  //	 4 bytes			size 12
	c string //	 8 bytes			size 20
	d int32  //	 4 bytes			size 24 - Aligned on 4 bytes
}

// No padding.
type np struct {
	a string // 16 bytes			size 16
	b string // 16 bytes			size 32
	c int32  //  8 bytes			size 40
	d int32  //  8 bytes			size 48 - Aligned on 8 bytes
}

func main() {

	// structlayout -json github.com/ardanlabs/gotraining/topics/language/struct_types/advanced/example1 nbp | structlayout-pretty
	var nbp nbp
	size := unsafe.Sizeof(nbp)
	fmt.Printf("SizeOf[%d][%p %p %p]\n", size, &nbp.a, &nbp.b, &nbp.c)

	// =========================================================================

	// structlayout -json github.com/ardanlabs/gotraining/topics/language/struct_types/advanced/example1 sbp | structlayout-pretty
	var sbp sbp
	size = unsafe.Sizeof(sbp)
	fmt.Printf("SizeOf[%d][%p %p]\n", size, &sbp.a, &sbp.b)

	// =========================================================================

	// structlayout -json github.com/ardanlabs/gotraining/topics/language/struct_types/advanced/example1 tbp | structlayout-pretty
	var tbp tbp
	size = unsafe.Sizeof(tbp)
	fmt.Printf("SizeOf[%d][%p %p]\n", size, &tbp.a, &tbp.b)

	// =========================================================================

	// structlayout -json github.com/ardanlabs/gotraining/topics/language/struct_types/advanced/example1 fbp | structlayout-pretty
	var fbp fbp
	size = unsafe.Sizeof(fbp)
	fmt.Printf("SizeOf[%d][%p %p]\n", size, &fbp.a, &fbp.b)

	// =========================================================================

	// structlayout -json github.com/ardanlabs/gotraining/topics/language/struct_types/advanced/example1 ebp | structlayout-pretty
	var ebp64 ebp64
	size = unsafe.Sizeof(ebp64)
	fmt.Printf("SizeOf[%d][%p %p %p %p]\n", size, &ebp64.a, &ebp64.b, &ebp64.c, &ebp64.d)

	// =========================================================================

	// structlayout -json github.com/ardanlabs/gotraining/topics/language/struct_types/advanced/example1 np | structlayout-pretty
	var np np
	size = unsafe.Sizeof(np)
	fmt.Printf("SizeOf[%d][%p %p %p %p]\n", size, &np.a, &np.b, &np.c, &np.d)
}

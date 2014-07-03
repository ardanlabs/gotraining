// Alignment is about placing types on boundaries that make the
// CPU access the fastest.

package main

import (
	"fmt"
	"unsafe"
)

type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	e := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	eNext := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	alignmentBoundary := unsafe.Alignof(e)

	sizeBool := unsafe.Sizeof(e.flag)
	offsetBool := unsafe.Offsetof(e.flag)

	sizeInt := unsafe.Sizeof(e.counter)
	offsetInt := unsafe.Offsetof(e.counter)

	sizeFloat := unsafe.Sizeof(e.pi)
	offsetFloat := unsafe.Offsetof(e.pi)

	sizeBoolNext := unsafe.Sizeof(eNext.flag)
	offsetBoolNext := unsafe.Offsetof(eNext.flag)

	fmt.Printf("Alignment Boundary: %d\n", alignmentBoundary)

	fmt.Printf("flag = Size: %d Offset: %d Addr: %v\n",
		sizeBool, offsetBool, &e.flag)

	fmt.Printf("counter = Size: %d Offset: %d Addr: %v\n",
		sizeInt, offsetInt, &e.counter)

	fmt.Printf("pi = Size: %d Offset: %d Addr: %v\n",
		sizeFloat, offsetFloat, &e.pi)

	fmt.Printf("Next = Size: %d Offset: %d Addr: %v\n",
		sizeBoolNext, offsetBoolNext, &eNext.flag)
}

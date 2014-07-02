package main

import (
	"fmt"
	"unsafe"
)

type Example struct {
	IntValue int16
	spacer   int
}

func main() {
	example := Example{
		IntValue: 10,
	}

	exampleNext := Example{
		IntValue: 10,
	}

	alignmentBoundary := unsafe.Alignof(example)

	sizeBool := unsafe.Sizeof(example.IntValue)
	offsetBool := unsafe.Offsetof(example.IntValue)

	sizeS := unsafe.Sizeof(example.spacer)
	offsetS := unsafe.Offsetof(example.spacer)

	sizeBoolNext := unsafe.Sizeof(exampleNext.IntValue)
	offsetBoolNext := unsafe.Offsetof(exampleNext.IntValue)

	fmt.Printf("Alignment Boundary: %d\n", alignmentBoundary)

	fmt.Printf("IntValue = Size: %d Offset: %d Addr: %p\n",
		sizeBool, offsetBool, &example.IntValue)

	fmt.Printf("spacer = Size: %d Offset: %d Addr: %p\n",
		sizeS, offsetS, &example.spacer)

	fmt.Printf("Next = Size: %d Offset: %d Addr: %p\n",
		sizeBoolNext, offsetBoolNext, &exampleNext.IntValue)
}

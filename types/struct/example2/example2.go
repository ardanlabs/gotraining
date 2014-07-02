// Alignment is about placing types on boundaries that make the
// CPU access the fastest.

package main

import (
	"fmt"
	"unsafe"
)

type Example struct {
	BoolValue  bool
	IntValue   int16
	FloatValue float32
}

func main() {
	example := Example{
		BoolValue:  true,
		IntValue:   10,
		FloatValue: 3.141592,
	}

	exampleNext := Example{
		BoolValue:  true,
		IntValue:   10,
		FloatValue: 3.141592,
	}

	alignmentBoundary := unsafe.Alignof(example)

	sizeBool := unsafe.Sizeof(example.BoolValue)
	offsetBool := unsafe.Offsetof(example.BoolValue)

	sizeInt := unsafe.Sizeof(example.IntValue)
	offsetInt := unsafe.Offsetof(example.IntValue)

	sizeFloat := unsafe.Sizeof(example.FloatValue)
	offsetFloat := unsafe.Offsetof(example.FloatValue)

	sizeBoolNext := unsafe.Sizeof(exampleNext.BoolValue)
	offsetBoolNext := unsafe.Offsetof(exampleNext.BoolValue)

	fmt.Printf("Alignment Boundary: %d\n", alignmentBoundary)

	fmt.Printf("BoolValue = Size: %d Offset: %d Addr: %v\n",
		sizeBool, offsetBool, &example.BoolValue)

	fmt.Printf("IntValue = Size: %d Offset: %d Addr: %v\n",
		sizeInt, offsetInt, &example.IntValue)

	fmt.Printf("FloatValue = Size: %d Offset: %d Addr: %v\n",
		sizeFloat, offsetFloat, &example.FloatValue)

	fmt.Printf("Next = Size: %d Offset: %d Addr: %v\n",
		sizeBoolNext, offsetBoolNext, &exampleNext.BoolValue)
}

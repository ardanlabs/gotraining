package main

import "fmt"

const (
	Enum1 = iota
	Enum2
	Enum3
)

const (
	FieldA = 1 << iota
	FieldB
	FieldC
)

func main() {
	fmt.Println("Enum1:", Enum1)
	fmt.Println("Enum2:", Enum2)
	fmt.Println("Enum3:", Enum3)

	fmt.Println()

	fmt.Println("FieldA:", FieldA)
	fmt.Println("FieldB:", FieldB)
	fmt.Println("FieldC:", FieldC)
}

package main

import (
	"fmt"
	"reflect"
)

// =============================================================================

// These functions are concrete implementations of print functions that can
// only work with slices of the specified type.

func printNumbers(numbers []int) {
	fmt.Print("Numbers: ")
	for _, num := range numbers {
		fmt.Print(num, " ")
	}
	fmt.Print("\n")
}

func printStrings(strings []string) {
	fmt.Print("Strings: ")
	for _, str := range strings {
		fmt.Print(str, " ")
	}
	fmt.Print("\n")
}

// =============================================================================

// This function provides an empty interface solution which uses type assertions
// for the different concrete slices to be supported. We've basically moved the
// functions from above into case statements.

func printAssert(v interface{}) {
	fmt.Print("Assert: ")
	switch list := v.(type) {
	case []int:
		for _, num := range list {
			fmt.Print(num, " ")
		}
	case []string:
		for _, str := range list {
			fmt.Print(str, " ")
		}
	}
	fmt.Print("\n")
}

// =============================================================================

// This function provides a reflection solution which allows a slice of any
// type to be provided and printed. This is a generic function thanks to the
// reflect package.

func printReflect(v interface{}) {
	fmt.Print("Reflect: ")
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Slice {
		return
	}
	for i := 0; i < val.Len(); i++ {
		fmt.Print(val.Index(i).Interface(), " ")
	}
	fmt.Print("\n")
}

// =============================================================================

// This function provides a generics solution which allows a slice of any type
// T (to be determined later) to be passed and printed.
//
// To avoid the ambiguity with array declarations, type parameters require a
// constraint to be applied. The `any` constraint states there is no constraint
// on what type T can become. The predeclared identifier `any` is an alias for
// `interface{}`.
//
// This code more closely resembles the concrete implementations that we started
// with and is easier to read than the reflect implementation.

func print[T any](slice []T) {
	fmt.Print("Generic: ")
	for _, v := range slice {
		fmt.Print(v, " ")
	}
	fmt.Print("\n")
}

// =============================================================================

func main() {
	numbers := []int{1, 2, 3}
	printNumbers(numbers)
	printAssert(numbers)
	printReflect(numbers)
	print(numbers)

	strings := []string{"A", "B", "C"}
	printStrings(strings)
	printAssert(strings)
	printReflect(strings)
	print(strings)
}

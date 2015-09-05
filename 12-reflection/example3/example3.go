// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/bWQ6hiVECQ

// https://github.com/goinggo/mapstructure
// Sample code provided by Mitchell Hashimoto

// Sample program to show how to use reflection to decode an integer.
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// main is the entry point for the application.
func main() {
	var number int
	decodeInt("10", &number)
	fmt.Println("number:", number)

	var age int
	decodeInt(45, &age)
	fmt.Println("age:", age)
}

// decodeInt accepts a value of any type and will decode
// that value to an integer.
func decodeInt(data interface{}, number *int) error {
	// Retrieve the value that the interface contains or points to.
	val := reflect.ValueOf(number).Elem()

	// Retrieve meta-data for the value to decode.
	dataVal := reflect.ValueOf(data)
	dataKind := getKind(dataVal)

	// Based on the kind of data to decode, perform the
	// specific logic.
	switch {
	case dataKind == reflect.Int:
		val.SetInt(dataVal.Int())
	case dataKind == reflect.Uint:
		val.SetInt(int64(dataVal.Uint()))
	case dataKind == reflect.Float32:
		val.SetInt(int64(dataVal.Float()))
	case dataKind == reflect.Bool:
		if dataVal.Bool() {
			val.SetInt(1)
		} else {
			val.SetInt(0)
		}
	case dataKind == reflect.String:
		i, err := strconv.ParseInt(dataVal.String(), 0, val.Type().Bits())
		if err == nil {
			val.SetInt(i)
		} else {
			return fmt.Errorf("cannot parse as int: %s", err)
		}
	default:
		return fmt.Errorf("expected type '%s', got unconvertible type '%s'", val.Type(), dataVal.Type())
	}

	return nil
}

// getKind provides support for identifying predeclared numeric
// types with implementation-specific sizes.
func getKind(val reflect.Value) reflect.Kind {
	kind := val.Kind()

	switch {
	case kind >= reflect.Int && kind <= reflect.Int64:
		return reflect.Int
	case kind >= reflect.Uint && kind <= reflect.Uint64:
		return reflect.Uint
	case kind >= reflect.Float32 && kind <= reflect.Float64:
		return reflect.Float32
	default:
		return kind
	}
}

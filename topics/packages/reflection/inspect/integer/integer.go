// All material is licensed under the Apache License Version 2.0, January 2016
// http://www.apache.org/licenses/LICENSE-2.0

// Example shows how to use reflection to decode an integer.
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	// Decode the string into the integer variable.
	var number int
	decodeInt("10", &number)
	fmt.Println("number:", number)

	// Decode the integer into the integer variable.
	var age int
	decodeInt(45, &age)
	fmt.Println("age:", age)
}

// decodeInt accepts a value of any type and will decode
// that value to an integer.
func decodeInt(v interface{}, number *int) error {

	// Inspect the concrete type value that is passed in.
	rv := reflect.ValueOf(v)

	// Retrieve the value that the integer pointer points to.
	val := reflect.ValueOf(number).Elem()

	// Based on the kind of data to decode, perform the specific logic.
	switch getKind(rv) {

	case reflect.Int:
		val.SetInt(rv.Int())
		return nil

	case reflect.Uint:
		val.SetInt(int64(rv.Uint()))
		return nil

	case reflect.Float32:
		val.SetInt(int64(rv.Float()))
		return nil

	case reflect.Bool:
		if rv.Bool() {
			val.SetInt(1)
			return nil
		}

		val.SetInt(0)
		return nil

	case reflect.String:
		i, err := strconv.ParseInt(rv.String(), 0, val.Type().Bits())
		if err != nil {
			return fmt.Errorf("cannot parse as int: %s", err)
		}

		val.SetInt(i)
		return nil

	default:
		return fmt.Errorf("expected type '%s', got unconvertible type '%s'", val.Type(), rv.Type())
	}
}

// getKind provides support for identifying predeclared numeric
// types with implementation-specific sizes.
func getKind(val reflect.Value) reflect.Kind {

	// Capture the value's Kind.
	kind := val.Kind()

	// Check each condition until a case is true.
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

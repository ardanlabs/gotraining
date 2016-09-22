// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type that represents a request for a customer invoice. Include a CustomerID and InvoiceID field. Define
// tags that can be used to validate the request. Define tags that specify both the length and range for the ID to be valid.
// Declare a function named validate that accepts values of any type and processes the tags. Display the resutls of the validation.
package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Customer represents a customer
type Customer struct {
	CustomerID int `length:"3" range:"100-300"`
	InvoiceID  int `length:"5" range:"60000-99999"`
}

func main() {

	// Declare a variable of type Customer.
	customer := Customer{
		CustomerID: 202,
		InvoiceID:  76545,
	}

	// Validate the value and display the results.
	validate(&customer)
}

// validate performs data validation on any struct type value.
func validate(value interface{}) {

	// Retrieve the value that the interface contains or points to.
	val := reflect.ValueOf(value).Elem()

	// Iterate over the fields of the struct value.
	for i := 0; i < val.NumField(); i++ {

		// Retrieve the field information.
		typeField := val.Type().Field(i)

		// Get the value as an int, string and the length.
		field := typeField.Name
		value := int(val.Field(i).Int())
		stringValue := strconv.Itoa(value)
		valueLength := utf8.RuneCountInString(stringValue)

		// Test the length first
		length, _ := strconv.Atoi(typeField.Tag.Get("length"))
		if valueLength != length {
			fmt.Printf("Invalid Length: Field[%s] Value[%d] - Len[%d] - ExpLen[%d]\n", field, value, valueLength, length)
			continue
		}

		// Test the range.
		r := strings.Split(typeField.Tag.Get("range"), "-")
		front, _ := strconv.Atoi(r[0])
		end, _ := strconv.Atoi(r[1])
		if value < front || value > end {
			fmt.Printf("Invalid Range: Field[%s] Value[%d] - Front[%d] - End[%d]\n", field, value, front, end)
			continue
		}

		fmt.Printf("VALID: Field[%s] Value[%d]\n", field, value)
	}
}

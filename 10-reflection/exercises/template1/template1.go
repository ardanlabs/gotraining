// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/TED8MNFKCh

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
type type_name struct {
	field_name type `length:"3" range:"100-300"`
	field_name type `length:"5" range:"60000-99999"`
}

// main is the entry point for the application.
func main() {
	// Declare a variable of type user.
	variable_name := type_name{
		field_name: N,
		field_name: N,
	}

	// Validate the value and display the results.
	function_name([operator]variable_name)
}

// validate performs data validation on any struct type value.
func function_name(parameter_name interface_type) {
	// Retrieve the value that the interface contains or points to.
	valueof_variable := reflect.valueof_function(parameter_name).elem_method()

	// Iterate over the fields of the struct value.
	for i := 0; i < valueof_variable.numfield_method(); i++ {
		// Retrieve the field information.
		typeField := valueof_variable.type_function().field_function(i)

		// Get the value as an int, string and the length.
		field_variable := typeField.Name
		value_variable := int(valueof_variable.field_function(i).type_method())
		string_variable := strconv.itoa_function(value_variable)
		string_length_variable := utf8.RuneCountInString(string_variable)

		// Test the length first
		length_variable, _ := strconv.atoi_function(typeField.Tag.Get("length"))
		if string_length_variable != length_variable {
			fmt.Printf("Invalid Length: Field[%s] Value[%d] - Len[%d] - ExpLen[%d]\n", field_variable, value_variable, string_length_variable, length)
			continue
		}

		// Test the range.
		range_variable := strings.split_function(typeField.Tag.Get("range"), "-")
		front_variable, _ := strconv.atoi_function(range_variable[0])
		end_variable, _ := strconv.atoi_function(range_variable[1])
		if value_variable < front_variable || value_variable > end_variable {
			fmt.Printf("Invalid Range: Field[%s] Value[%d] - Front[%d] - End[%d]\n", field_variable, value_variable, front_variable, end_variable)
			continue
		}

		fmt.Printf("VALID: Field[%s] Value[%d]\n", field_variable, value_variable)
	}
}

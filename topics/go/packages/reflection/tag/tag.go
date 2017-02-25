// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Example shows how to reflect on a struct type with tags.
package main

import (
	"fmt"
	"reflect"
	"regexp"
)

// User is a sample struct.
type User struct {
	Name  string `valid:"exists"`
	Email string `valid:"regexp" exp:"[\w.%+-]+@(?:[[:alnum:]-]+\\.)+[[:alpha:]]{2,6}"`
}

// Result provides a detail view of the validation results.
type Result struct {
	Field  string
	Type   string
	Value  string
	Test   string
	Result bool
}

func main() {

	// Declare a variable of type user.
	user := User{
		Name:  "Henry Ford",
		Email: "henry@ford.com",
	}

	// Validate the value and display the results.
	results := validate(&user)
	for _, result := range results {
		fmt.Printf("%+v\n", result)
	}
}

// validate performs data validation on any struct type value.
func validate(value interface{}) []Result {

	// Declare a nil slice of Result values.
	var results []Result

	// Retrieve the value that the interface contains or points to.
	val := reflect.ValueOf(value).Elem()

	// Iterate over the fields of the struct value.
	for i := 0; i < val.NumField(); i++ {

		// Retrieve the field information.
		typeField := val.Type().Field(i)

		// Declare a variable of type Result and initialize
		// it with all the meta-data.
		result := Result{
			Field: typeField.Name,
			Type:  typeField.Type.String(),
			Value: val.Field(i).String(),
			Test:  typeField.Tag.Get("valid"),
		}

		// Perform the requested tests.
		switch result.Test {
		case "exists":
			if result.Value != "" {
				result.Result = true
			}

		case "regexp":
			m, err := regexp.MatchString(typeField.Tag.Get("exp"), result.Value)
			if err == nil && m == true {
				result.Result = true
			}
		}

		// Append the results to the slice.
		results = append(results, result)
	}

	return results
}

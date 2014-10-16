// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/y0WyYezH05

/*
ValueOf returns a new Value initialized to the concrete value stored in the interface i.
ValueOf(nil) returns the zero Value.
func ValueOf(i interface{}) Value {
*/

// Sample program to show how to reflect on a struct type with tags.
package main

import (
	"fmt"
	"reflect"
	"regexp"
)

// User is a sample struct.
type User struct {
	Name  string `valid:"exists"`
	Email string `valid:"regexp" exp:"(?i)[A-Z0-9._%+-]+@(?:[A-Z0-9-]+\\.)+[A-Z]{2,6}"`
}

// Result provides a detail view of the validation results.
type Result struct {
	Field  string
	Type   string
	Value  string
	Test   string
	Result bool
}

// main is the entry point for the application.
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
	// Declare a nil slce of Result values.
	var results []Result

	// Retrieve the value that the interface contains or points to.
	val := reflect.ValueOf(value).Elem()

	// Iterate over the fields of the struct value.
	for i := 0; i < val.NumField(); i++ {
		// Retrieve the field information.
		typeField := val.Type().Field(i)

		// Declare a variable of type Result and initalize
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

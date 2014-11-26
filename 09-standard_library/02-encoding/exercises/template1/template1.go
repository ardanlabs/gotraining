// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/tHy_2p4obQ

// Create a file with an array of JSON documents that contain a user name and email address. Declare a struct
// type that maps to the JSON document. Using the json package, read the file and create a slice of this struct
// type. Display the slice.
//
// From example 1, Marshal the slice into pretty print strings and display each element.
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// user is the data we need to unmarshal and marshal.
type type_name struct {
	field_name type `json:"json_field_name"`
	field_name type `json:"json_field_name"`
}

// main is the entry point for the application.
func main() {
	// Open the file.
	file_variable_name, err_variable_name := os.open_function("data.json")
	if err_name != nil {
		fmt.Println("Open File", err_variable_name)
		return
	}

	// Schedule the file to be closed once
	// the function returns.
	defer file_variable_name.close_function()

	// Decode the file into this slice.
	var slice_name []type_name
	err_variable_name = json.NewDecoder(file_variable_name).Decode(&slice_name)
	if err != nil {
		fmt.Println("Decode File", err_variable_name)
		return
	}

	// Iterate over the slice and display
	// each user.
	for _, variable_name := range slice_name {
		fmt.Printf("%+v\n", variable_name)
	}

	variable_name, err_variable_name := json.MarshalIndent(&slice_name, "", "    ")
	if err_variable_name != nil {
		fmt.Println("MarshalIndent", err_variable_name)
		return
	}

	// Convert the byte slice to a string and display.
	fmt.Println(type(variable_name))
}

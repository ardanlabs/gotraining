// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Test program for the benchmark example.
package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var vars = map[string]string{
	"variable_name": "bill",
}

func main() {
	variable := "#string:variable_name"
	v := getValue(variable, vars)

	v, ok := v.(string)
	if !ok {
		log.Fatalln("Invalid type was returned")
	}

	if v != `"bill"` {
		log.Fatalf("Invalid value was returned : %v", v)
	}

	log.Printf("Value was returned : %v", v)
}

// getValue takes a string with a specialized format, processes it and return
// an native value.
func getValue(variable string, vars map[string]string) interface{} {

	// variable: "#cmd:variable_name"

	// Trim the # symbol from the string.
	value := strings.TrimLeft(variable, "#")

	// Find the first instance of the separator.
	idx := strings.Index(value, ":")
	if idx == -1 {
		return nil
	}

	// Split the key and variable apart.
	cmd := value[0:idx]
	vari := value[idx+1:]

	// Does the variable_name exist in the vars map.
	param, exists := vars[vari]
	if !exists {
		return nil
	}

	// What native format do we need.
	switch cmd {
	case "number":
		i, err := strconv.Atoi(param)
		if err != nil {
			return nil
		}
		return i

	case "string":
		return fmt.Sprintf("%q", param)

	default:
		return nil
	}
}

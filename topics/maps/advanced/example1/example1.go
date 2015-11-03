// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/mycosI0zpN

// Sample program to show how to compose maps of maps.
package main

import "fmt"

// Data defines a key/value store.
type Data map[string]string

// main is the entry point for the application.
func main() {
	// Declare and make a map of Data type values.
	users := make(map[string]Data)

	// Initialize some data into our map of maps.
	users["clients"] = Data{"123": "Henry", "654": "Joan"}
	users["admins"] = Data{"398": "Bill", "076": "Steve"}

	// Iterate over the map of maps.
	for key, data := range users {
		fmt.Println(key)
		for key, value := range data {
			fmt.Println("\t", key, value)
		}
	}
}

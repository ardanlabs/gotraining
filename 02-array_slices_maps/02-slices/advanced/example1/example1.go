// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/-qQgO7NbLm

// Sample program to show the practical use of slices.
package main

import (
	"encoding/json"
	"fmt"
)

//  bson is a named type that declares a map of key/value pairs.
type bson map[string]interface{}

// user is a struct type that declares user information.
type user struct {
	ID   int
	Name string
}

// main is the entry point for the application.
func main() {
	// Retrieve profiles for all users in the system.
	users, err := retrieveUsers()
	if err != nil {
		fmt.Println(err)
	}

	// Display each user from the slice.
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}
}

// retrieveUsers retrieves user documents for the specified
// user and returns a pointer to a user type value.
func retrieveUsers() ([]user, error) {
	// Make a call to get the json response.
	r, err := getUsers()
	if err != nil {
		return nil, err
	}

	// Unmarshal the json document into a value of
	// the user struct type.
	var users []user
	err = json.Unmarshal([]byte(r), &users)
	return users, err
}

// getUsers simulates a web call that returns an array
// of json documents with users.
func getUsers() (string, error) {
	response := `[{"id":1432, "name":"sally"},{"id":4312, "name":"bill"},{"id":6721, "name":"jane"}]`
	return response, nil
}

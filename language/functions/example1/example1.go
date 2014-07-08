// Example shows how functions can return multiple values while using
// named and struct types.
package main

import (
	"encoding/json"
	"fmt"
)

//  bson is a named type that declares a map of key/value pairs.
type bson map[string]interface{}

// user is a struct type that declares user information.
type user struct {
	Id   int
	Name string
}

func main() {
	// Retrieve the user profile.
	u, err := RetrieveUser("sally")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display the user profile.
	fmt.Printf("%+v\n", *u)
}

// RetrieveUser retrieves the user document for the specified
// user and returns a pointer to a user type value.
func RetrieveUser(name string) (*user, error) {
	// Make the web call to get the json response.
	r, err := GetUser(name)
	if err != nil {
		return nil, err
	}

	// Unmarshal the json document into a value of
	// the user struct type.
	var u user
	err = json.Unmarshal([]byte(r), &u)
	return &u, err
}

// GetUser simulates a web call that returns a json
// document for the specified user.
func GetUser(name string) (string, error) {
	response := `{"id":1432, "name":"sally"}`
	return response, nil
}

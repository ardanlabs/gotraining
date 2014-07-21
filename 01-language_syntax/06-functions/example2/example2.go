// Sample program to show how we can use the blank identifier to ignore return values.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type (
	// user is a struct type that declares user information.
	user struct {
		ID   int
		Name string
	}

	// updateStats provides update stats.
	updateStats struct {
		Modified int
		Duration float64
		Success  bool
		Message  string
	}
)

// main is the entry point for the application.
func main() {
	// Declare and initalize a value of type user.
	u := user{
		ID:   1432,
		Name: "Betty",
	}

	// Update the user name. Don't care about the update stats.
	if _, err := updateUser(&u); err != nil {
		fmt.Println(err)
		return
	}

	// Display the update was successful.
	fmt.Println("Updated user record for ID", u.ID)
}

// updateUser updates the specified user document.
func updateUser(u *user) (*updateStats, error) {
	// Make a call to post the update.
	r, err := postUpdate(u)
	if err != nil {
		return nil, err
	}

	// Unmarshal the json document into a value of
	// the userStats struct type.
	var us updateStats
	if err = json.Unmarshal([]byte(r), &us); err != nil {
		return nil, err
	}

	// Check the update status to verify the update
	// was successful.
	if us.Success != true {
		return nil, errors.New(us.Message)
	}

	return &us, nil
}

// postUpdate simulates a web call that returns a json
// document for the specified user update.
func postUpdate(u *user) (string, error) {
	response := `{"modified":1, "duration":0.005, "success" : true, "message": "updated"}`
	return response, nil
}

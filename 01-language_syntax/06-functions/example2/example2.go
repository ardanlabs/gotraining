// NEED PLAYGROUND

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
		id   int
		name string
	}

	// updateStats provides update stats.
	updateStats struct {
		modified int
		duration float64
		success  bool
		message  string
	}
)

// main is the entry point for the application.
func main() {
	// Declare and initalize a value of type user.
	u := user{
		id:   1432,
		name: "Betty",
	}

	// Update the user name. Don't care about the update stats.
	if _, err := updateUser(&u); err != nil {
		fmt.Println(err)
		return
	}

	// Display the update was successful.
	fmt.Println("Updated user record for id", u.id)
}

// updateUser updates the specified user document.
func updateUser(u *user) (*updateStats, error) {
	// response simulates a JSON response.
	response := `{"modified":1, "duration":0.005, "success" : true, "message": "updated"}`

	// Unmarshal the json document into a value of
	// the userStats struct type.
	var us updateStats
	if err := json.Unmarshal([]byte(response), &us); err != nil {
		return nil, err
	}

	// Check the update status to verify the update
	// was successful.
	if us.success != true {
		return nil, errors.New(us.message)
	}

	return &us, nil
}

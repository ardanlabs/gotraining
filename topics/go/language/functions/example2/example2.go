// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how we can use the blank identifier to
// ignore return values.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// user is a struct type that declares user information.
type user struct {
	ID   int
	Name string
}

// updateStats provIDes update stats.
type updateStats struct {
	Modified int
	Duration float64
	Success  bool
	Message  string
}

func main() {

	// Declare and initialize a value of type user.
	u := user{
		ID:   1432,
		Name: "Betty",
	}

	// Update the user Name. Don't care about the update stats.
	if _, err := updateUser(&u); err != nil {
		fmt.Println(err)
		return
	}

	// Display the update was Successful.
	fmt.Println("Updated user record for ID", u.ID)
}

// updateUser updates the specified user document.
func updateUser(u *user) (*updateStats, error) {

	// response simulates a JSON response.
	response := `{"Modified":1, "Duration":0.005, "Success" : true, "Message": "updated"}`

	// Unmarshal the json document into a value of
	// the userStats struct type.
	var us updateStats
	if err := json.Unmarshal([]byte(response), &us); err != nil {
		return nil, err
	}

	// Check the update status to verify the update
	// was Successful.
	if us.Success != true {
		return nil, errors.New(us.Message)
	}

	return &us, nil
}

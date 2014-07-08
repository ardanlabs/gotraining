// Example shows how we can use the blank identifier to ignore return values.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// user is a struct type that declares user information.
type user struct {
	Id   int
	Name string
}

// updateStats provides update stats.
type updateStats struct {
	Modified int
	Duration float64
	Success  bool
	Message  string
}

func main() {
	// Declare and initalize a value of type user.
	u := user{
		Id:   1432,
		Name: "Betty",
	}

	// Update the user name. Don't care about the update stats.
	if _, err := UpdateUser(&u); err != nil {
		fmt.Println(err)
		return
	}

	// Display the update was successful.
	fmt.Println("Updated user record for Id", u.Id)
}

// RetrieveUser retrieves the user document for the specified
// user and returns a pointer to a user type value.
func UpdateUser(u *user) (*updateStats, error) {
	// Make the web call to post the update.
	r, err := PostUpdate(u)
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

// PostUpdate simulates a web call that returns a json
// document for the specified user update.
func PostUpdate(u *user) (string, error) {
	response := `{"modified":1, "duration":0.005, "success" : true, "message": "updated"}`
	return response, nil
}

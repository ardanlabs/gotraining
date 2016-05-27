// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package api provides an example on how to use go-fuzz.
package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Routes initializes the routes.
func Routes() {
	http.HandleFunc("/process", Process)
}

// SendError responds with an error.
func SendError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(struct{ Error string }{err.Error()})
}

// Process handles the processing of data.
func Process(w http.ResponseWriter, r *http.Request) {

	// Capture the data that was posted over.
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		SendError(w, err)
		return
	}

	// If we received no data return an error.
	if len(data) == 0 {
		SendError(w, errors.New("Empty data value"))
		return
	}

	// Split the data by comma.
	parts := strings.Split(string(data), ",")

	// Need a named type for our user.
	type user struct {
		Type string
		Name string
		Age  int
	}

	// Create a slice of users.
	var users []user

	// Iterate over the set of users we received.
	for _, part := range parts {

		// Capture the type of user.
		typ := part[:3]

		// Capture the age and convert to integer.
		age, err := strconv.Atoi(part[3:5])
		if err != nil {
			SendError(w, err)
			return
		}

		// Capture the users name.
		name := part[5:]

		// Add a user to the slice.
		users = append(users, user{typ, name, age})
	}

	// Respond with the processed data.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

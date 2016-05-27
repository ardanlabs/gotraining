// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package fuzzprot provides the ability to unpack user values from
// our binary protocol.
package fuzzprot

import (
	"errors"
	"strconv"
)

// User represents the data we are receiving.
type User struct {
	Type string
	Name string
	Age  int
}

// UnpackUsers knows how to extract user values from the binary data.
func UnpackUsers(data []byte) ([]User, error) {

	// How many users did we receive.
	count := data[0]

	// Create a slice of users for the number we received.
	users := make([]User, count)

	// Slice the header byte away.
	data = data[1:]

	// Index for each user we are processing.
	var uidx int

	// Don't stop until we extract all the users.
	for len(data) > 0 {

		switch data[0] {
		case 0:
			uidx++
			data = data[1:]

		case 1:
			users[uidx].Type, data = grabString(data[1:])

		case 2:
			users[uidx].Name, data = grabString(data[1:])

		case 3:
			var err error
			data = data[1:]

			users[uidx].Age, err = strconv.Atoi(string(data[:2]))
			if err != nil {
				return nil, err
			}

			data = data[2:]

		default:
			return nil, errors.New("unknown field type")
		}
	}

	return users, nil
}

// grabString knows how to take the specified bytes and convert
// those bytes into a string.
func grabString(data []byte) (string, []byte) {
	l, data := data[0], data[1:]
	return string(data[:l]), data[l:]
}

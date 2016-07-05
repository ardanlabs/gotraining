// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0
package main

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"message"`
}

// GetUsers is a sample database request to get all the users.
func GetUsers(conn *nats.EncodedConn, m *nats.Msg) error {
	log.Println("UserHandler:", m)

	users := []User{
		{17896678, "Bill"},
		{89778799, "Lisa"},
	}

	j, err := json.Marshal(users)

	if err != nil {
		msg := Message{Status: 1, Text: "Failed to marshal user data."}
		return conn.Publish(m.Reply, &msg)
	}

	msg := Message{
		Status: 0,
		JSON:   string(j), // would normally use []byte, but this is readable :)
	}

	return conn.Publish(m.Reply, &msg)
}

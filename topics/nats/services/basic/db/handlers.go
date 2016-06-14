// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0
package main

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/nats-io/nats"
)

// GetUsers is a sample database requrest to get all the users.
func GetUsers(conn *nats.EncodedConn, reply string, req Request) {
	log.Println("UserHandler:", req)

	data := []struct {
		ID   int    `json:"id"`
		Name string `json:"message"`
	}{
		{17896678, "Bill"},
		{89778799, "Lisa"},
	}

	SendResponse(conn, reply, data)
	return
}

//==============================================================================

// ResponseHandler handles all the response routing.
func ResponseHandler(conn *nats.EncodedConn, m *nats.Msg) {
	log.Printf("Message received: [%s] %s\n", m.Subject, string(m.Data))

	// Unmarshal the reuqest.
	var req Request
	if err := json.Unmarshal(m.Data, &req); err != nil {
		log.Println("ERROR:", err)
		SendResponseError(conn, m.Reply, err)
		return
	}

	// Find the route for this request.
	f, exists := Route(req.URI)
	if !exists {
		err := errors.New("Routes does not exist")
		log.Println("ERROR:", err)
		SendResponseError(conn, m.Reply, err)
		return
	}

	// Execute the request.
	f(conn, m.Reply, req)
}

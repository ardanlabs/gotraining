// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats"
)

// Request represents the message structure we receive for request.
type Request struct {
	URI     string                 `json:"uri"`
	Post    []byte                 `json:"post"`
	Headers map[string]interface{} `json:"headers"`
}

// Response represents the message structure we respond with.
type Response struct {
	Status int                      `json:"status"`
	JSON   []map[string]interface{} `json:"json"`
}

//==============================================================================

// SendResponse returns a response for a processed request.
func SendResponse(conn *nats.EncodedConn, reply string, v interface{}) error {
	log.Println("Reponding:", reply)

	data, err := json.Marshal(v)
	if err != nil {
		return SendResponseError(conn, reply, err)
	}

	var m []map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return SendResponseError(conn, reply, err)
	}

	resp := Response{
		Status: 0,
		JSON:   m,
	}

	return conn.Publish(reply, resp)
}

// SendResponseError returns a response for a processed request that failed.
func SendResponseError(conn *nats.EncodedConn, reply string, err error) error {
	log.Println("Reponding Error:", reply, err)

	var m []map[string]interface{}
	json.Unmarshal([]byte(fmt.Sprintf(`[{"error":"%s"}]`, err.Error())), &m)

	resp := Response{
		Status: 1,
		JSON:   m,
	}

	return conn.Publish(reply, resp)
}

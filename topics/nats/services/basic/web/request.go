package main

import (
	"log"
	"time"
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

// SendRequest sends a request and waits for the response.
func SendRequest(subject string, req Request) (Response, error) {
	log.Printf("SendRequest: Subject[%s] Req[%+v]\n", subject, req)

	// Send the request and wait for the response.
	var resp Response
	err := conn.Request(subject, req, &resp, time.Second)

	return resp, err
}

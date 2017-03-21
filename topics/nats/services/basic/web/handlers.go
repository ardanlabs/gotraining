// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// requestSub is the general subject for DB requests.
const requestSubject = "dbReq"

// GetUsers is a sample database requrest handler.
func GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("GetUsers: Started")
	defer log.Println("GetUsers: Completed")

	// Build the request.
	req := Request{
		URI: "/db/users",
	}

	// Send the request for processing.
	resp, err := SendRequest(requestSubject, req)
	if err != nil {
		log.Println("GetUsers: ERROR:", err)
		json.NewEncoder(w).Encode(struct{ Error error }{err})
		return
	}

	log.Println("GetUsers: RESP:", resp.JSON)

	// Return the response back to the caller.
	json.NewEncoder(w).Encode(resp.JSON)
	return
}

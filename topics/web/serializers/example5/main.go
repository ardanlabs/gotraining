// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program that sends and receives JSON.
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Status is the status of our logging system.
type Status struct {
	Load     float64 `json:"load"`
	Messages int     `json:"messages"`
}

// Event is a message we need to log.
type Event struct {
	Host    string `json:"host"`
	Message string `json:"message"`
	Level   int    `json:"level"`
}

// App gives the handler which is the main entry point for our application.
func App() http.Handler {

	h := func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			status(w, r)
		case http.MethodPost:
			logEvent(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}

	return http.HandlerFunc(h)
}

// status responds to GET requests by guaging the status of the system and
// writing it to the ResponseWriter as JSON.
func status(res http.ResponseWriter, req *http.Request) {

	s := Status{
		Load:     1.0, // TODO
		Messages: 42,  // TODO
	}

	res.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err := json.NewEncoder(res).Encode(s); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

// logEvent reads a JSON string from the request body and turns it into an
// Event value. It is then logged and we respond 204 No Content.
func logEvent(res http.ResponseWriter, req *http.Request) {

	var e Event

	// Encode the json value received into the event value.
	if err := json.NewDecoder(req.Body).Decode(&e); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf(
		"Level %d event from %s: %s",
		e.Level,
		e.Host,
		e.Message,
	)

	res.WriteHeader(http.StatusNoContent)
}

func main() {
	log.Print("Listening on localhost:3000")
	log.Fatal(http.ListenAndServe("localhost:3000", App()))
}

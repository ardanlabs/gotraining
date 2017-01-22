// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program that implements a simple web service that will allow us to
// explore how to use the http/pprof tooling.
package main

import (
	"encoding/json"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/sendjson", sendJSON)

	log.Println("listener : Started : Listening on: http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}

// sendJSON returns a simple JSON document.
func sendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "bill",
		Email: "bill@ardanlabs.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}

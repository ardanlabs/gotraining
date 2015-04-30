// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Sample program that implements a simple web service.
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// main is the entry point for the application.
func main() {
	Routes()

	log.Println("listener : Started : Listening on: http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}

// Routes sets the routes for the web service.
func Routes() {
	http.HandleFunc("/sendjson", SendJSON)
}

// SendJSON returns a simple JSON document.
func SendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "Bill",
		Email: "bill@ardanstudios.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)

	if err := json.NewEncoder(rw).Encode(u); err != nil {
		log.Panic(err)
	}

	LogResponse(&u)
}

// LogResponse is used to write the response to the log.
func LogResponse(v interface{}) {
	d, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		log.Println("Unable to Marshal Response", err)
		return
	}

	log.Println(string(d))
}

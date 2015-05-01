// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Sample program that implements a simple web service.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

	data, err := json.Marshal(&u)
	if err != nil {
		// We want this error condition to panic so we get a stack trace. This should
		// never happen. The http package will catch the panic and provide logging
		// and return a 500 back to the caller.
		log.Panic("Unable to unmatshal response", err)
	}

	datalen := len(data) + 1 // account for trailing LF
	h := rw.Header()
	h.Set("Content-Type", "application/json")
	h.Set("Content-Length", strconv.Itoa(datalen))
	rw.WriteHeader(200)
	fmt.Fprintf(rw, "%s\n", data)

	LogResponse(&u)
}

// LogResponse is used to write the response to the log.
func LogResponse(v interface{}) {
	d, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		log.Println("Unable to marshal response", err)
		return
	}

	log.Println(string(d))
}

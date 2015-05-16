// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Package handlers provides the endpoints for the web service.
package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

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
		log.Panic("Unable to unmarshal response", err)
	}

	datalen := len(data) + 1 // account for trailing LF
	h := rw.Header()
	h.Set("Content-Type", "application/json")
	h.Set("Content-Length", strconv.Itoa(datalen))
	rw.WriteHeader(200)
	fmt.Fprintf(rw, "%s\n", data)
}

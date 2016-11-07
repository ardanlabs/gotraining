// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to handle forms with JSON.
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/schema"
)

// App provides a handler to handle GET and POST calls
// for every request.
func App() http.Handler {

	// Delcare the handler function to handle the GET and POST call.
	h := func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			GetHandler(res, req)

		case "POST":
			PostHandler(res, req)

		default:
			res.WriteHeader(http.StatusMethodNotAllowed)
		}
	}

	// Store and return the handler function within
	// a http.Handler interface value.
	return http.HandlerFunc(h)
}

// GetHandler provides support for the GET reponse.
func GetHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	res.Write([]byte(`
<form action="/" method="POST">
<p>
	<input type="text" name="FirstName" placeholder="First Name" />
</p>
<p>
	<input type="text" name="LastName" placeholder="Last Name" />
</p>
<p>
	<input type="submit" value="CLICK ME!!" />
</p>
</form>`))
}

// User represents a user in the system.
type User struct {
	FirstName string
	LastName  string
}

// PostHandler provides support for the POST reponse.
func PostHandler(res http.ResponseWriter, req *http.Request) {

	// Parse the raw query from the URL and update r.Form.
	if err := req.ParseForm(); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a user value.
	var u User

	// Decode the JSON from the Post data to the User value.
	if err := schema.NewDecoder().Decode(&u, req.PostForm); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write this formatted string into the response.
	fmt.Fprintf(res, "First Name: %s\nLast Name: %s", u.FirstName, u.LastName)
}

func main() {

	// Start the http server to handle the request for
	// both versions of the API.
	log.Fatal(http.ListenAndServe(":3000", App()))
}

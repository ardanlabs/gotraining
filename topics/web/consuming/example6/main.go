// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to how to use JWT for authentication.
package main

import (
	"log"
	"net/http"

	jose "github.com/dvsekhvalnov/jose2go"
)

// sharedSecret contains our key for decoding the JWT token.
var sharedSecret = []byte("some shared secret")

// App returns a handler for handling requests with JWT.
func App() http.Handler {

	// Handler function to process the request.
	h := func(res http.ResponseWriter, req *http.Request) {

		// Extract the JWT from the request header.
		s := req.Header.Get("x-signature")
		if s == "" {
			res.WriteHeader(http.StatusPreconditionRequired)
			return
		}

		// Verify, decrypt and decompress the JWT received in the request.
		payload, _, err := jose.Decode(s, sharedSecret)
		if err != nil || payload == "" {
			res.WriteHeader(http.StatusPreconditionFailed)
			return
		}

		// Respond with a 200 and return the payload we extracted.
		token, err := jose.Sign("", jose.HS256, sharedSecret)
		if err != nil {
			res.WriteHeader(500)
			res.Write([]byte(err.Error()))
			return
		}
		res.Header().Set("x-signature", token)
		res.WriteHeader(200)
		res.Write([]byte(payload))

		// Note that we are using the payload from the signature and are
		// effectively ignoring the actual request body. If you intend to use
		// the request body for anything you should confirm that it matches the
		// payload or they could reuse a valid signature for some other body
		// but change the body. Why even include the body then? Why not just
		// send the JWT in the body instead of a header? The benefit here is
		// improved readability in dev tools. You'd have to weigh that cost
		// against the cost of larger requests.
	}

	return http.HandlerFunc(h)
}

func main() {

	// Start the http server to handle requests
	log.Print("Listening on localhost:3000")
	log.Fatal(http.ListenAndServe("localhost:3000", App()))
}

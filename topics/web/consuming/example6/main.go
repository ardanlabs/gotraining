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

// App returns a handler for handling requets with JWT.
func App() http.Handler {

	// Handler function to process the reuqest.
	h := func(res http.ResponseWriter, req *http.Request) {

		// Extract the JWT from the request header.
		s := req.Header.Get("x-signature")
		if s == "" {
			res.WriteHeader(http.StatusPreconditionRequired)
			return
		}

		// Verify, decrypt and decompresses the JWT received in the request.
		payload, _, err := jose.Decode(s, sharedSecret)
		if err != nil || payload == "" {
			res.WriteHeader(http.StatusPreconditionFailed)
			return
		}

		// Response with a 200 and return the payload we extracted.
		token, err := jose.Sign("", jose.HS256, sharedSecret)
		if err != nil {
			res.WriteHeader(500)
			res.Write([]byte(err.Error()))
			return
		}
		res.Header().Set("x-signature", token)
		res.WriteHeader(200)
		res.Write([]byte(payload))
	}

	return http.HandlerFunc(h)
}

func main() {

	// Start the http server to handle the request for
	// both versions of the API.
	log.Fatal(http.ListenAndServe(":3000", App()))
}

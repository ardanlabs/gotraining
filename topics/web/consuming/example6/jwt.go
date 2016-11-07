// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	jose "github.com/dvsekhvalnov/jose2go"
)

// JWTTransporter provides a custom transporter for making requests
// with JWT authentication.
type JWTTransporter struct {
	transporter  http.RoundTripper
	sharedSecret []byte
}

// RoundTrip implements the RoundTripper interface.
func (c *JWTTransporter) RoundTrip(req *http.Request) (*http.Response, error) {

	// Read in the response from the api call.
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	// Replace the request Body with this new NopCloser.
	req.Body = ioutil.NopCloser(bytes.NewBuffer(b))

	// Produce a signed JWT token from our payload, using HS256 and the
	// shared secret.
	token, err := jose.Sign(string(b), jose.HS256, c.sharedSecret)
	if err != nil {
		return nil, err
	}

	// Set the JWT inside the request.
	req.Header.Set("x-signature", token)

	// Perform the request with the token and payload.
	return c.transporter.RoundTrip(req)
}

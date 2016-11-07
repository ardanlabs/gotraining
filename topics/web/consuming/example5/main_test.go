// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests to show how to consume a web API using a custom transporter
// by implementing the RoundTripper interface.
package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// App returns a handler that can be used to mock the call.
func App() http.Handler {

	// Handler function will be used for mocking.
	h := func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello World!"))
	}

	// Return the handler function.
	return http.HandlerFunc(h)
}

// customTransporter provides our custom transporter
// for sending requests.
type customTransporter struct {
	transporter http.RoundTripper
	logs        []string
}

// RoundTrip implements the RoundTripper interface.
func (c *customTransporter) RoundTrip(req *http.Request) (*http.Response, error) {

	// Log the beginning of the request.
	c.logs = append(c.logs, "started request")

	// Make the request call using the default implementation.
	res, err := c.transporter.RoundTrip(req)

	// Log the end of the request.
	c.logs = append(c.logs, "completed request")

	return res, err
}

func TestApp(t *testing.T) {

	// Startup a server to handle processing these routes.
	ts := httptest.NewServer(App())
	defer ts.Close()

	// Create a new request for the GET call.
	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a value of the custom transporter.
	trans := customTransporter{
		transporter: http.DefaultTransport,
		logs:        []string{},
	}

	// Create a client value to make the request. Bind our
	// custom transporter.
	client := http.Client{
		Transport: &trans,
	}

	// Perform the GET call with our client and transport.
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// Read in the response from the api call.
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Validate we received the expected response.
	got := strings.TrimSpace(string(b))
	want := `Hello World!`
	if got != want {
		t.Log("Wanted:", want)
		t.Log("Got   :", got)
		t.Fatal("Mismatch")
	}

	// Validate we got the logs.
	elogs := []string{"started request", "completed request"}
	if strings.Join(elogs, " ") != strings.Join(trans.logs, " ") {
		t.Fatalf("expected %s to equal %s", trans.logs, elogs)
	}
}

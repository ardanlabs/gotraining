// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests for the sample program to show how to use JWTs
package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
)

func TestLoginHandler(t *testing.T) {
	tests := []struct {
		Body       string
		StatusCode int
	}{
		{`{"email": "jacob@example.com", "password": "rory"}`, http.StatusOK},
		{`{"email": "anna@example.com", "password": "rory"}`, http.StatusUnauthorized},
		{`{"email": "jacob@example.com", "password": "BAD"}`, http.StatusUnauthorized},
		{`{"email": "", "password": ""}`, http.StatusUnauthorized},
		{``, http.StatusBadRequest},
	}

	// Start a server to handle these requests.
	ts := httptest.NewServer(App())
	defer ts.Close()

	for _, tt := range tests {

		// Send the request body to /login
		body := strings.NewReader(tt.Body)
		res, err := http.Post(ts.URL+"/login", "application/json", body)
		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != tt.StatusCode {
			t.Log("Wanted:", tt.StatusCode)
			t.Log("Got   :", res.StatusCode)
			t.Fatal("Mismatch status code")
		}

		// If we didn't expect success then move on to the next test
		if tt.StatusCode != http.StatusOK {
			continue
		}

		// Read in the response from the api call.
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		// Validate we received a token in the response. We're just checking
		// that it looks kind of like a token here.
		re := regexp.MustCompile(`^([-\w=+/]+\.){2}[-\w=+/]+$`)
		got := strings.TrimSpace(string(b))
		if !re.MatchString(got) {
			t.Log("Wanted a token")
			t.Log("Got:", got)
			t.Fatal("Mismatch")
		}
	}
}

// expired is a valid but old token for testing
const expired = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODg0OTM0ODMsImlhdCI6MTQ4ODQ5MzQ4MiwibmJmIjoxNDg4NDkzNDgyLCJzdWIiOiJqYWNvYkBleGFtcGxlLmNvbSJ9.e0cWExi4hHxODo44x5P1wHHPKlZukFk93ib3f2UOwEY`

func TestSecureHandler(t *testing.T) {
	// Start a server to handle these requests.
	ts := httptest.NewServer(App())
	defer ts.Close()

	// Before we test the secure handler let's get a valid token
	// Send a request to /login to get a token
	body := strings.NewReader(`{"email": "jacob@example.com", "password": "rory"}`)
	res, err := http.Post(ts.URL+"/login", "application/json", body)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Log("Wanted:", http.StatusOK)
		t.Log("Got   :", res.StatusCode)
		t.Fatal("Could not log in")
	}

	// Read in the token
	tkn, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		Token      string
		Want       string
		StatusCode int
	}{
		{string(tkn), "Be sure to drink your Ovaltine!", http.StatusOK},
		{"bad token", "Not authorized", http.StatusUnauthorized},
		{expired, "Not authorized", http.StatusUnauthorized},
		{"", "Not authorized", http.StatusUnauthorized},
	}
	for _, tt := range tests {

		// Create a new request for the GET call.
		req, err := http.NewRequest("GET", ts.URL+"/secure", nil)
		if err != nil {
			t.Fatal(err)
		}

		// If we have a token then add it to the Authentication header
		if tt.Token != "" {
			req.Header.Set("Authentication", "Bearer "+tt.Token)
		}

		// Create a Client and perform the GET call.
		var c http.Client
		res, err := c.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != tt.StatusCode {
			t.Log("Wanted:", tt.StatusCode)
			t.Log("Got   :", res.StatusCode)
			t.Fatal("Mismatch")
		}

		// Read in the response from the api call.
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		// Validate we received the expected response.
		got := strings.TrimSpace(string(b))
		want := tt.Want
		if got != want {
			t.Log("Wanted:", want)
			t.Log("Got   :", got)
			t.Fatal("Mismatch")
		}
	}
}

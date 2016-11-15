// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests for the sample program to how to use JWT for authentication.
package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestApp(t *testing.T) {

	// Startup a server to handle processing these routes.
	ts := httptest.NewServer(App())
	defer ts.Close()

	// Create the payload which is a simple user document.
	payload, err := json.Marshal(map[string]string{"name": "Jane Smith"})
	if err != nil {
		t.Fatal(err)
	}

	// Test the request without the JWT.
	t.Run("without signature", func(st *testing.T) {

		// Create a new request for the POST call with our payload.
		req, err := http.NewRequest("POST", ts.URL, bytes.NewReader(payload))
		if err != nil {
			st.Fatal(err)
		}

		// Create a Client and perform the POST call.
		var client http.Client
		res, err := client.Do(req)
		if err != nil {
			st.Fatal(err)
		}

		// Did the call fail authentication?
		if res.StatusCode != http.StatusPreconditionRequired {
			st.Fatalf("expected response status to be %d, got %d", http.StatusPreconditionRequired, res.StatusCode)
		}
	})

	// Test the request with a JWT.
	t.Run("with signature", func(st *testing.T) {
		table := []struct {
			secret []byte
			status int
			want   string
		}{
			{sharedSecret, http.StatusOK, `{"name":"Jane Smith"}`},
		}

		for _, tt := range table {

			// Create a new request for the POST call with our payload.
			req, err := http.NewRequest("POST", ts.URL, bytes.NewReader(payload))
			if err != nil {
				st.Fatal(err)
			}

			// Create a Client using our custom transporter. Provide
			// the shared secret.
			client := http.Client{
				Transport: &JWTTransporter{
					transporter:  http.DefaultTransport,
					sharedSecret: tt.secret,
				},
			}

			// Perform the request passing our credentials for auth.
			res, err := client.Do(req)
			if err != nil {
				st.Fatal(err)
			}

			// Did the call fail authentication?
			if res.StatusCode != tt.status {
				st.Fatalf("expected response status to be %d, got %d", tt.status, res.StatusCode)
			}

			if tt.status == http.StatusOK {

				// Read in the response from the api call.
				b, err := ioutil.ReadAll(res.Body)
				if err != nil {
					t.Fatal(err)
				}

				// Validate we received the expected response.
				got := strings.TrimSpace(string(b))
				want := tt.want
				if got != want {
					t.Log("Wanted:", want)
					t.Log("Got   :", got)
					t.Fatal("Mismatch")
				}
			}
		}
	})
}

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_App(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	payload, err := json.Marshal(map[string]string{"name": "Jane Smith"})
	if err != nil {
		t.Fatal(err)
	}

	t.Run("without signature", func(st *testing.T) {
		client := http.Client{}

		req, err := http.NewRequest("POST", ts.URL, bytes.NewReader(payload))
		if err != nil {
			st.Fatal(err)
		}

		res, err := client.Do(req)
		if err != nil {
			st.Fatal(err)
		}

		if res.StatusCode != http.StatusPreconditionRequired {
			st.Fatalf("expected response status to be %d, got %d", http.StatusPreconditionRequired, res.StatusCode)
		}
	})

	t.Run("with a valid signature", func(st *testing.T) {
		client := http.Client{
			Transport: &JWTTransporter{
				transporter:  http.DefaultTransport,
				sharedSecret: SharedSecret,
			},
		}

		req, err := http.NewRequest("POST", ts.URL, bytes.NewReader(payload))
		if err != nil {
			st.Fatal(err)
		}

		res, err := client.Do(req)
		if err != nil {
			st.Fatal(err)
		}

		if res.StatusCode != http.StatusOK {
			st.Fatalf("expected response status to be %d, got %d", http.StatusOK, res.StatusCode)
		}

		b, _ := ioutil.ReadAll(res.Body)
		exp := `{"name":"Jane Smith"}`
		act := string(b)
		if exp != act {
			st.Fatalf("expected matching payloads: got %s expected %s", act, exp)
		}
	})

	t.Run("with an invalid signature", func(st *testing.T) {
		client := http.Client{
			Transport: &JWTTransporter{
				transporter:  http.DefaultTransport,
				sharedSecret: []byte("invalid shared secret"),
			},
		}

		req, err := http.NewRequest("POST", ts.URL, bytes.NewReader(payload))
		if err != nil {
			st.Fatal(err)
		}

		res, err := client.Do(req)
		if err != nil {
			st.Fatal(err)
		}

		if res.StatusCode != http.StatusPreconditionFailed {
			st.Fatalf("expected response status to be %d, got %d", http.StatusPreconditionFailed, res.StatusCode)
		}
	})
}

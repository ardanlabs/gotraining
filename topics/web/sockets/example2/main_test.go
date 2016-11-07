// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests for the sample program on how to use the Gorilla web socket
// package to bind HTTP requests.
package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/net/websocket"
)

func TestApp(t *testing.T) {

	// Startup a server to handle processing these routes.
	ts := httptest.NewServer(App())
	defer ts.Close()

	// Replace the `http to `ws` in the url.
	wsURL := strings.Replace(ts.URL, "http", "ws", 1) + "/socket"

	// Dial a websocket into our mock server.
	ws, err := websocket.Dial(wsURL, "", "http://127.0.0.1")
	if err != nil {
		t.Fatal(err)
	}

	// Send this string across the websocket connection.
	if _, err = ws.Write([]byte("hello, world!")); err != nil {
		t.Fatal(err)
	}

	// Wait to read back the response from our handler.
	msg := make([]byte, 512)
	read, err := ws.Read(msg)
	if err != nil {
		t.Fatal(err)
	}

	// Encode the received byte into a message.
	var message Message
	err = json.NewDecoder(bytes.NewReader(msg[:read])).Decode(&message)
	if err != nil {
		t.Fatal(err)
	}

	// Create a table of what we expect.
	tests := []struct {
		Got  string
		Want string
	}{
		{message.Formatted, "HELLO, WORLD!"},
		{message.Original, "hello, world!"},
	}

	// Check the different fields.
	for _, tt := range tests {
		if tt.Got != tt.Want {
			t.Log("Wanted:", tt.Want)
			t.Log("Got   :", tt.Got)
			t.Fatal("Mismatch")
		}
	}
}

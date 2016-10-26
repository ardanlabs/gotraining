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
	ts := httptest.NewServer(App())
	defer ts.Close()

	wsURL := strings.Replace(ts.URL, "http", "ws", 1) + "/socket"

	ws, err := websocket.Dial(wsURL, "", "http://127.0.0.1")
	if err != nil {
		t.Fatal(err)
	}

	_, err = ws.Write([]byte("hello, world!"))
	if err != nil {
		t.Fatal(err)
	}

	msg := make([]byte, 512)
	read, err := ws.Read(msg)
	if err != nil {
		t.Fatal(err)
	}

	message := &Message{}
	err = json.NewDecoder(bytes.NewReader(msg[:read])).Decode(message)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		Got  string
		Want string
	}{
		{message.Formatted, "HELLO, WORLD!"},
		{message.Original, "hello, world!"},
	}

	for _, tt := range tests {
		if tt.Got != tt.Want {
			t.Log("Wanted:", tt.Want)
			t.Log("Got   :", tt.Got)
			t.Fatal("Mismatch")
		}
	}

}

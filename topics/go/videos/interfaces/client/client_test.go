package client

import (
	"fmt"
	"net/http"
	"testing"
)

type errTransport struct{}

// RoundTrip implements http.RoundTripper interface
func (e errTransport) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, fmt.Errorf("can't connect")
}

func TestConnectionError(t *testing.T) {
	c := New("http://example.com")
	c.c.Transport = errTransport{}
	err := c.Health()
	if err == nil {
		t.Fatal("no error")
	}
}

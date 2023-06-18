package client

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

type ErrTransport struct{}

// implement http.RoundTripper
func (e ErrTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	return nil, fmt.Errorf("%s", "connection error")
}

func TestHealthConnectionError(t *testing.T) {
	c := NewAPIClient("https://example.com")
	c.client.Transport = ErrTransport{}
	err := c.Health(context.Background())
	require.Error(t, err)
}

package client

import (
	"context"
	"fmt"
	"net/http"
)

type Client struct {
	baseURL string
	client  http.Client
}

func NewAPIClient(baseURL string) Client {
	return Client{baseURL: baseURL}
}

func (c *Client) Health(ctx context.Context) error {
	url := fmt.Sprintf("%s/health", c.baseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", resp.Status)
	}

	return nil
}

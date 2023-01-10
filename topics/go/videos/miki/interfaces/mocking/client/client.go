package client

import (
	"fmt"
	"net/http"
)

type Client struct {
	baseURL string
	c       http.Client
}

func New(url string) *Client {
	c := Client{
		baseURL: url,
	}

	return &c
}

func (c *Client) Health() error {
	url := fmt.Sprintf("%s/health", c.baseURL)
	resp, err := c.c.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", resp.Status)
	}

	return nil
}

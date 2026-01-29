package vms

import (
	"shorts/vms/after/model"
)

type StartRequest = model.StartRequest
type StartResponse = model.StartResponse

type Client struct{}

func (c *Client) Start(req StartRequest) StartResponse {
	// FIXME: Implement
	return StartResponse{}
}

package tests

import (
	"io"
	"net/http"
	"net/url"
)

// Succeed is the UTF-8 byte squence for a check mark.
const Succeed = "\xE2\x9C\x93"

// Succeed is the UTF-8 byte squence for an X mark.
const Failed = "\xE2\x9C\x97"

// NewRequest used to setup a request for mocking API calls with httptreemux.
func NewRequest(method, path string, body io.Reader) *http.Request {
	r, _ := http.NewRequest(method, path, body)
	u, _ := url.Parse(path)
	r.URL = u
	r.RequestURI = path
	return r
}

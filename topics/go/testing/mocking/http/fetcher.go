// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package http is a sample package to show how to mock a http call response
package http

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
)

// Doer executes *http.Request
// this interface helps us to mock a http.Client.Do function
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

// DoerFunc implements Doer
type DoerFunc func(*http.Request) (*http.Response, error)

func (d DoerFunc) Do(req *http.Request) (*http.Response, error) {
	return d(req)
}

// Fetcher fetches data from url
type Fetcher interface {
	Fetch(ctx context.Context, url string) ([]byte, error)
}

// NewGetter constructs get struct
func NewGetter(doer Doer) Fetcher {
	return &get{
		doer: doer,
	}
}

// get implements Fetcher
// get fetch data from url using http.MethodGet
type get struct {
	doer        Doer
}

// Fetch fetches url result using http.MethodGet
func (g get) Fetch(ctx context.Context, url string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := g.doer.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(http.StatusText(resp.StatusCode))
	}

	return io.ReadAll(resp.Body)
}

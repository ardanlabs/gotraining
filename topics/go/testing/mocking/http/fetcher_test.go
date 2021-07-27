package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

var (
	getter Fetcher
)

// mockedDoer implements Doer
type mockedDoer struct {
	StatusCode int
	Body       string
}

func (m mockedDoer) Do(*http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: m.StatusCode,
		Body:       io.NopCloser(strings.NewReader(m.Body)),
	}, nil
}

func init() {
	// real implementation
	//getter = NewGetter(http.DefaultClient)

	// mocked implementation using Doer example
	//getter = NewGetter(mockedDoer{http.StatusNotFound, "mocked result"})

	// mocked implementation using DoerFunc example
	getter = NewGetter(DoerFunc(func(request *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader("mocked result")),
		}, nil
	}))
}

func TestGet_Fetch(t *testing.T) {
	resp, err := getter.Fetch(context.Background(), "https://google.com")
	fmt.Println(string(resp), err)
}

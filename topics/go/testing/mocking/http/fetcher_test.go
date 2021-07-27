// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package http is a sample package to show how to mock a http call response
package http

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
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

// TestGet_FetchSingleTest test fetcher using DoerFunc implementation
func TestGet_FetchSingleTest(t *testing.T) {
	getter := NewGetter(DoerFunc(func(request *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader("mocked result")),
		}, nil
	}))

	want := "mocked result"
	got, err := getter.Fetch(context.Background(), "https://google.com")

	if err != nil {
		t.Fatalf("err is not expected: %v", err)
	}

	if want != string(got) {
		t.Fatalf(`result is not expected. want: %v, got: %v`, want, string(got))
	}
}

// TestGet_FetchTestTable test fetcher using mockedDoer implementation
func TestGet_FetchTestTable(t *testing.T) {
	tt := []struct {
		MockedResponse     string
		ExpectedResponse   string
		ExpectedStatusCode int
	}{
		{
			MockedResponse:     "mocked result",
			ExpectedResponse:   "mocked result",
			ExpectedStatusCode: http.StatusOK,
		},
		{
			MockedResponse:     "mocked result",
			ExpectedResponse:   "Not Found",
			ExpectedStatusCode: http.StatusNotFound,
		},
	}

	for _, test := range tt {
		doer := &mockedDoer{
			StatusCode: test.ExpectedStatusCode,
			Body:       test.MockedResponse,
		}
		resp, err := NewGetter(doer).Fetch(context.Background(), "mocked")

		if test.ExpectedStatusCode == http.StatusOK {
			if err != nil {
				t.Fatalf("err is not expected: %v", err)
			}

			if test.ExpectedResponse != string(resp) {
				t.Fatalf(`result is not expected. want: %v, got: %v`, test.ExpectedResponse, string(resp))
			}
		} else {
			if resp != nil {
				t.Fatalf("resp is not expected: %v", resp)
			}

			if err == nil {
				t.Fatal("err is expected")
			}

			if err.Error() != test.ExpectedResponse {
				t.Fatalf(`result is not expected. want: %v, got: %v`, test.ExpectedResponse, err.Error())
			}
		}
	}
}

// Package endpointTests implements tests for the buoy endpoints.
package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/goinggo/concurrentwebservice/service"
)

// TestRSSSearch tests the search screen loads.
func TestRSSSearch(t *testing.T) {
	r, _ := http.NewRequest("GET", "/search", nil)
	w := httptest.NewRecorder()

	http.DefaultServeMux.ServeHTTP(w, r)

	response := w.Body.String()

	t.Log("Calling the search endpoint should return status 200.")
	if w.Code != http.StatusOK {
		t.Fatalf("Expecting 200 => Received %d", w.Code)
	}

	t.Log("The response body should not be empty.")
	if w.Body.Len() == 0 {
		t.Fatalf("Length of response Body is %d", w.Body.Len())
	}

	t.Log("The search term \"Sample App\" should be in the response.")
	if strings.Index(response, "Sample App") < 0 {
		t.Error("Search term \"Sample App\" not found in response.")
	}
}

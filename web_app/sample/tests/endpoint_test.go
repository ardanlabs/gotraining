// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Package endpointTests implements tests for the buoy endpoints.
package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/ArdanStudios/gotraining/web_app/sample/service"
)

// TestSearchLoad tests the search screen loads.
func TestSearchLoad(t *testing.T) {
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

// TestSearch tests an actual search is working.
func TestSearch(t *testing.T) {
	r, _ := http.NewRequest("POST", "/search", strings.NewReader("searchterm=golang&google=on"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
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

	t.Log("The search term \"golang\" should be in the response.")
	if strings.Index(response, "golang") < 0 {
		t.Error("Search term \"golang\" not found in response.")
	}
}

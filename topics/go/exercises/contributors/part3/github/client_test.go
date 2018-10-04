package github

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// token is a valid token we can reuse in testing.
const token = "781b0bdab2315c62134544eff45811333c663797"

// TestNewClient builds clients with different token values and asserts on what
// kind of values should error.
func TestNewClient(t *testing.T) {

	tests := []struct {
		name      string
		token     string
		shouldErr bool
	}{
		{"success", token, false},
		{"missing token", "", true},
		{"short token", "abc", true},
		{"long token", "abcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabc", true},
		{"invalid token", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", true},
	}

	for _, test := range tests {
		// Kick off a subtest for this scenario.
		fn := func(t *testing.T) {
			_, err := NewClient(API, test.token)

			if test.shouldErr && err == nil {
				t.Errorf("NewClient(%q) should error but did not", test.token)
			} else if !test.shouldErr && err != nil {
				t.Errorf("NewClient(%q) should not error but gave %v", test.token, err)
			}
		}
		t.Run(test.name, fn)
	}
}

func TestContributorsSuccess(t *testing.T) {
	// Create a server to mock out the response from GitHub.
	f := func(w http.ResponseWriter, r *http.Request) {

		// Ensure they made a GET request.
		if got, want := r.Method, http.MethodGet; got != want {
			t.Errorf("Method did not match: Got %q want %q", got, want)
			return
		}

		// Ensure they're asking for the correct path based on our repo.
		if got, want := r.URL.Path, "/repos/golang/go/contributors"; got != want {
			t.Errorf("Path did not match: Got %q want %q", got, want)
			return
		}

		// Ensure they sent the token in the auth header.
		if got, want := r.Header.Get("Authorization"), `Bearer `+token; got != want {
			t.Errorf("Auth token did not match: Got %q want %q", got, want)
			return
		}

		body := `[{"login": "anna", "contributions": 27}, {"login": "jacob", "contributions": 18}, {"login": "kell", "contributions": 9}, {"login": "carter", "contributions": 6}, {"login": "rory", "contributions": 1}]`
		if _, err := w.Write([]byte(body)); err != nil {
			t.Fatal(err)
		}
	}
	srv := httptest.NewServer(http.HandlerFunc(f))
	defer srv.Close()

	// Create the client using the test server's URL for the api.
	c, err := NewClient(srv.URL, token)
	if err != nil {
		t.Fatal(err)
	}

	// Call the method under test.
	got, err := c.ContributorList("golang/go")

	if err != nil {
		t.Fatalf("Client should not error. Got %v", err)
	}

	want := []Contributor{
		{Login: "anna", Contributions: 27},
		{Login: "jacob", Contributions: 18},
		{Login: "kell", Contributions: 9},
		{Login: "carter", Contributions: 6},
		{Login: "rory", Contributions: 1},
	}

	// Use the nice github.com/google/go-cmp/cmp library to compare what we got
	// against what we expected. This could be done manually or with
	// reflect.DeepEqual but I prefer this method for its diffing output and for
	// its awareness of Equal methods.
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("contributors returned from client did not match expected:\n%s", diff)
	}
}

func TestContributorsAPIFailure(t *testing.T) {

	// Make a mock server where the API fails
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	}
	srv := httptest.NewServer(http.HandlerFunc(fn))
	defer srv.Close()

	c, err := NewClient(srv.URL, token)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := c.ContributorList("golang/go"); err == nil {
		t.Fatal("Client should error but did not")
	}
}

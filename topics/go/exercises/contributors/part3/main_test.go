package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/exercises/contributors/part3/github"
)

// mock implements the "contributors" interface without requiring us to
// actually call the GitHub API.
type mock struct{}

// Contributors satisfies the main package's "contributors" interface. It
// returns predefined result sets for different repo values.
func (mock) Contributors(repo string) ([]github.Contributor, error) {
	switch repo {
	case "golang/go":
		return []github.Contributor{
			{Login: "anna", Contributions: 27},
			{Login: "jacob", Contributions: 18},
			{Login: "kell", Contributions: 9},
			{Login: "carter", Contributions: 6},
			{Login: "rory", Contributions: 1},
		}, nil
	}

	return nil, errors.New("could not reach API")
}

func TestPrintContributors(t *testing.T) {

	tests := []struct {
		name   string
		repo   string
		status int
		want   string
	}{
		{
			name:   "failure",
			repo:   "should/fail",
			status: 1,
			want:   `Error fetching contributors: could not reach API`,
		},
		{
			name:   "success",
			repo:   "golang/go",
			status: 0,
			want: `0 anna 27
1 jacob 18
2 kell 9
3 carter 6
4 rory 1`,
		},
	}

	for _, test := range tests {
		test := test // Copy to this scope for our closure.

		t.Run(test.name, func(t *testing.T) {

			// Create a mock of the API.
			var c mock
			// Use a bytes.Buffer to hold the output from our function.
			var buf bytes.Buffer

			// Call the function under test.
			status := printContributors(&buf, test.repo, c)

			// Assert on our results.
			got := strings.TrimSpace(buf.String())

			if got != test.want {
				t.Errorf("Printed report did not match expectation")
				t.Logf("Got:\n%s", got)
				t.Logf("Want:\n%s", test.want)
			}

			if status != test.status {
				t.Errorf("printContributors should return status code %d, got %d", test.status, status)
			}
		})
	}
}

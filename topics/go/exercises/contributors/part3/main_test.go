package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/exercises/contributors/part3/github"
)

// mock implements the "contributorLister" interface without requiring us to
// actually call the GitHub API.
type mock struct{}

// ContributorList satisfies the main package's "contributorLister" interface.
// It returns predefined result sets for different repo values.
func (mock) ContributorList(repo string) ([]github.Contributor, error) {
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

func TestProcess(t *testing.T) {

	tests := []struct {
		name      string
		repo      string
		want      string
		shouldErr bool
	}{
		{
			name:      "failure",
			repo:      "should/fail",
			shouldErr: true,
		},
		{
			name:      "success",
			repo:      "golang/go",
			want:      "0 anna 27\n1 jacob 18\n2 kell 9\n3 carter 6\n4 rory 1",
			shouldErr: false,
		},
	}

	for _, test := range tests {
		// Create a mock of the API.
		var c mock

		fn := func(t *testing.T) {

			// Use a bytes.Buffer to hold the output from our function.
			var buf bytes.Buffer

			// Call the function under test.
			err := process(&buf, test.repo, c)

			// Assert on our results.
			got := strings.TrimSpace(buf.String())

			if test.shouldErr && err == nil {
				t.Fatal("process should error but did not")
			}

			if !test.shouldErr && err != nil {
				t.Fatalf("process should not error but did: %v", err)
			}

			if got != test.want {
				t.Errorf("Printed report did not match expectation")
				t.Logf("Got:\n%s", got)
				t.Logf("Want:\n%s", test.want)
			}
		}
		t.Run(test.name, fn)
	}
}

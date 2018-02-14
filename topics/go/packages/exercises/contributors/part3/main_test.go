package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/packages/exercises/contributors/part3/github"
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

func TestPrintContributorsSuccess(t *testing.T) {

	// Create a mock of the API.
	var c mock

	// Use a bytes.Buffer to hold the output from our function.
	var buf bytes.Buffer

	// Call the function under test.
	status := printContributors(&buf, "golang/go", c)

	// Assert on our results.
	want := `0 anna 27
1 jacob 18
2 kell 9
3 carter 6
4 rory 1`
	got := strings.TrimSpace(buf.String())

	if got != want {
		t.Errorf("Printed report did not match expectation")
		t.Logf("Got:\n%s", got)
		t.Logf("Want:\n%s", want)
	}

	if status != 0 {
		t.Errorf("Successful run should return status code 0, got %d", status)
	}
}

func TestPrintContributorsFailure(t *testing.T) {

	var c mock
	var buf bytes.Buffer

	status := printContributors(&buf, "failure", c)

	want := `Error fetching contributors: could not reach API`
	got := strings.TrimSpace(buf.String())

	if got != want {
		t.Errorf("Printed report did not match expectation")
		t.Logf("Got:\n%s", got)
		t.Logf("Want:\n%s", want)
	}

	if status != 1 {
		t.Errorf("Failed run should return status code 1, got %d", status)
	}
}

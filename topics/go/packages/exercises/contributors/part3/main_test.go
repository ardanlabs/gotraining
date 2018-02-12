package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/jcbwlkr/contributors/github"
)

// mock implements the "contributors" interface in this package.
type mock func() ([]github.Contributor, error)

// Contributors satisfies the main package's "contributors" interface.
func (fn mock) Contributors() ([]github.Contributor, error) {
	return fn()
}

func TestPrintContributorsSuccess(t *testing.T) {

	// create a mock where the API works
	c := mock(func() ([]github.Contributor, error) {
		return []github.Contributor{
			{Login: "anna", Contributions: 27},
			{Login: "jacob", Contributions: 18},
			{Login: "kell", Contributions: 9},
			{Login: "carter", Contributions: 6},
			{Login: "rory", Contributions: 1},
		}, nil
	})
	// use a bytes.Buffer to hold the output from our function
	var buf bytes.Buffer

	// Call the function under test
	status := printContributors(&buf, c)

	// Assert on our results
	want := `anna	27
jacob	18
kell	9
carter	6
rory	1`
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

	c := mock(func() ([]github.Contributor, error) {
		return nil, errors.New("could not reach API")
	})
	var buf bytes.Buffer

	status := printContributors(&buf, c)

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

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Go + testing = ❤!
//
// This is our first iteration of the test. It only tests a single
// input/output combination. Run this with:
//
// go test
package strings

import "testing"

func TestReverse(t *testing.T) {
	in := "jacob"
	want := "bocaj"
	got := Reverse(in)

	if got != want {
		t.Errorf("Reverse(%q) = %q, want %q", in, got, want)
	}
}

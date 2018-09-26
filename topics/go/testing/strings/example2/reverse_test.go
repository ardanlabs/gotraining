// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Go + testing = ❤!
//
// This is our second test iteration. We define a type for a test scenario as
// well as a "table" of test inputs. We loop over that and run each test. This
// is commonly called a "Table Driven Test".
package strings

import "testing"

// reverseTest defines a single scenario for testing the Reverse function.
type reverseTest struct {
	in   string
	want string
}

func TestReverse(t *testing.T) {

	// tests is a "table" of test scenarios.
	tests := []reverseTest{
		{"jacob", "bocaj"},
		{"john", "nhoj"},
	}

	// loop through the table and run each test.
	for _, test := range tests {
		got := Reverse(test.in)
		if got != test.want {
			t.Errorf("Reverse(%q) = %q, want %q", test.in, got, test.want)
		}
	}
}

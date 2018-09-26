// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Go + testing = ❤!
//
// This is our third test iteration. We added more scenarios but one is
// failing. We name each of our scenarios so we can identify the failing
// scenario. We also launch them in subtests so we can run just the failing
// scenario on demand. Run with -v to see each subtest in action. To limit
// execution use -run:
//
// go test -v
// go test -v -run TestReverse/multibyte
package strings

import "testing"

// reverseTest defines a single scenario for testing the Reverse function.
type reverseTest struct {
	name string
	in   string
	want string
}

func TestReverse(t *testing.T) {

	// tests is a "table" of test scenarios.
	tests := []reverseTest{
		{"odd", "jacob", "bocaj"},
		{"even", "john", "nhoj"},
		{"multibyte", "I ❤ NY", "YN ❤ I"},
		{"empty", "", ""},
	}

	// loop through the table and run each test.
	for _, test := range tests {

		// Encapsulate the test behavior in a closure.
		fn := func(t *testing.T) {
			got := Reverse(test.in)
			if got != test.want {
				t.Errorf("Reverse(%q) = %q, want %q", test.in, got, test.want)
			}
		}

		// Run the test closure as a subtest
		t.Run(test.name, fn)
	}
}

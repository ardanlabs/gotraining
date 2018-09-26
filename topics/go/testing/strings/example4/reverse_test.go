// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Go + testing = ❤!
//
// In this iteration we haven't really changed the test but we have added a
// benchmark. This is how we can measure the performance of our function. To
// run benchmarks run these commands:
//
// go test -bench .
// go test -bench . -benchmem
// go test -bench . -benchmem -benchtime 5s
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
		{"alphabet", "abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba"},
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

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("abcdefghijklmnopqrstuvwxyz")
	}
}

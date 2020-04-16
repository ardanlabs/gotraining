package strings_test

import (
	strings "github.com/ardanlabs/gotraining/topics/go/algorithms/strings/permuation"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestIsPermutation(t *testing.T) {

	permutationTests := []struct {
		name     string
		input    string
		input2   string
		expected bool
	}{
		{"empty string", "", "", true},
		{"old number of string test", "god", "dog", true},
		{"different size inputs", "god", "do", false},
		{"binary string (even number of strings)", "1001", "0110", true},
	}

	for _, tt := range permutationTests {
		got := strings.IsPermutation(tt.input, tt.input2)
		if got != tt.expected {
			t.Logf("\t%s\tString is a palindrome: %s\n.", failed, tt.input)
			t.Fatalf("\t\tGot %v, Expected %v.", got, tt.expected)
		}
		t.Logf("\t%s\tString %s is a palindrome.", succeed, tt.input)

	}
}

package strings_test

import (
	strings "github.com/ardanlabs/gotraining/topics/go/algorithms/strings/palindrome"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestIsPalindrome(t *testing.T) {

	revTests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", true},
		{"string with length of 1", "G", true},
		{"string with odd length", "bob", true},
		{"string with even length", "otto", true},
		{"chinese", "汉字汉", true},
		{"failed test", "test", true},
	}

	for _, tt := range revTests {
		got := strings.IsPalindrome(tt.input)
		if got != tt.expected {
			t.Logf("\t%s\tString is a palindrome: %s\n.", failed, tt.input)
			t.Fatalf("\t\tGot %v, Expected %v.", got, tt.expected)
		}
		t.Logf("\t%s\tString %s is a palindrome.", succeed, tt.input)

	}
}

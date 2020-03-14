package strings_test

import (
	strings "github.com/ardanlabs/gotraining/topics/go/algorithms/strings/reverse"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestReverseString(t *testing.T) {

	revTests := []struct {
		name     string
		input    string
		expected string
	}{
		{"odd-string-length", "Hello World", "dlroW olleH"},
		{"even-string-length", "go", "og"},
		{"chinese", "汉字", "字汉"},
		{"two-runes", "é́́", "é́́"},
	}

	for _, tt := range revTests {
		got := strings.ReverseString(tt.input)
		if got != tt.expected {
			t.Logf("\t%s\tShould be able to reverse string: %s\n.", failed, tt.input)
			t.Fatalf("\t\tGot %s, Expected %s.", got, tt.expected)
		}
		t.Logf("\t%s\tShould be able to reverse string.", succeed)
	}
}

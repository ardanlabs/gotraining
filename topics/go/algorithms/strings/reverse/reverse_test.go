package strings_test

import (
	"testing"

	strings "github.com/ardanlabs/gotraining/topics/go/algorithms/strings/reverse"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestReverseString(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{"odd", "Hello World", "dlroW olleH"},
		{"even", "go", "og"},
		{"chinese", "汉字", "字汉"},

		{"tworunes", "é́́", "é́́"},
	}

	for _, test := range tt {
		tf := func(t *testing.T) {
			t.Log("Given the need to test reverse string functionality.")
			{
				t.Logf("\tWhen checking the word %q.", test.input)
				{
					got := strings.ReverseString(test.input)
					if got != test.expected {
						t.Logf("\t%s\tShould have gotten back the string reversed.", failed)
						t.Fatalf("\t\tGot %q, Expected %q", got, test.expected)
					}
					t.Logf("\t%s\tShould have gotten back the string reversed.", succeed)
				}
			}
		}
		t.Run(test.name, tf)
	}
}

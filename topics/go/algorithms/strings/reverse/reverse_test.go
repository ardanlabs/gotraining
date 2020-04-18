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

	t.Log("Given the need to test reverse string functionality.")
	{
		for testID, test := range tt {
			tf := func(t *testing.T) {
				t.Logf("\tTest: %d\tWhen checking the word %q.", testID, test.input)
				{
					got := strings.ReverseString(test.input)
					if got != test.expected {
						t.Logf("\t%s\tTest: %d\tShould have gotten back the string reversed.", failed, testID)
						t.Fatalf("\t\tTest: %d\tGot %q, Expected %q", testID, got, test.expected)
					}
					t.Logf("\t%s\tTest: %d\tShould have gotten back the string reversed.", succeed, testID)
				}
			}
			t.Run(test.name, tf)
		}
	}
}

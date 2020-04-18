package strings_test

import (
	"testing"

	strings "github.com/ardanlabs/gotraining/topics/go/algorithms/strings/palindrome"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestIsPalindrome(t *testing.T) {
	tt := []struct {
		name    string
		input   string
		success bool
	}{
		{"empty", "", true},
		{"one", "G", true},
		{"odd", "bob", true},
		{"even", "otto", true},
		{"chinese", "汉字汉", true},
		{"not", "test", false},
	}

	for _, test := range tt {
		tf := func(t *testing.T) {
			t.Log("Given the need to test palindrome functionality.")
			{
				t.Logf("\tWhen checking the word %q.", test.input)
				{
					got := strings.IsPalindrome(test.input)
					switch test.success {
					case true:
						if !got {
							t.Fatalf("\t%s\tShould have seen the string was a palindrome.", failed)
						}
						t.Logf("\t%s\tShould have seen the string was a palindrome.", succeed)
					case false:
						if got {
							t.Fatalf("\t%s\tShould have seen the string was not a palindrome.", failed)
						}
						t.Logf("\t%s\tShould have seen the string was not a palindrome.", succeed)
					}
				}
			}
		}
		t.Run(test.name, tf)
	}
}

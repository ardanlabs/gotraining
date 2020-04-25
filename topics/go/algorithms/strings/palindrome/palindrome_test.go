/*
	// This is the API you need to build for these tests. You will need to
	// change the import path in this test to point to your code.

	package palindrome

	// Is checks if a string is a Palindrome.
	func Is(input string) bool
*/

package palindrome_test

import (
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/algorithms/strings/palindrome"
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

	t.Log("Given the need to test palindrome functionality.")
	{
		for testID, test := range tt {
			tf := func(t *testing.T) {
				t.Logf("\tTest %d:\tWhen checking the word %q.", testID, test.input)
				{
					got := palindrome.Is(test.input)
					switch test.success {
					case true:
						if !got {
							t.Fatalf("\t%s\tTest %d:\tShould have seen the string was a palindrome.", failed, testID)
						}
						t.Logf("\t%s\tTest %d:\tShould have seen the string was a palindrome.", succeed, testID)
					case false:
						if got {
							t.Fatalf("\t%s\tTest %d:\tShould have seen the string was not a palindrome.", failed, testID)
						}
						t.Logf("\t%s\tTest %d:\tShould have seen the string was not a palindrome.", succeed, testID)
					}
				}
			}
			t.Run(test.name, tf)
		}
	}
}

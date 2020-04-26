package palindrome_test

import (
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/algorithms/numbers/palindrome"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestIsPalindrome(t *testing.T) {
	tt := []struct {
		name    string
		input   int
		success bool
	}{
		{"negative", -1, false},
		{"one", 1, true},
		{"nine", 9, true},
		{"ten", 10, false},
		{"even", 1001, true},
		{"odd", 151, true},
	}

	t.Log("Given the need to test palindrome functionality.")
	{
		for testID, test := range tt {
			tf := func(t *testing.T) {
				t.Logf("\tTest %d:\tWhen checking the integer %d.", testID, test.input)
				{
					got := palindrome.Is(test.input)
					switch test.success {
					case true:
						if !got {
							t.Fatalf("\t%s\tTest %d:\tShould have seen the integer was a palindrome.", failed, testID)
						}
						t.Logf("\t%s\tTest %d:\tShould have seen the integer was a palindrome.", succeed, testID)
					case false:
						if got {
							t.Fatalf("\t%s\tTest %d:\tShould have seen the integer was not a palindrome.", failed, testID)
						}
						t.Logf("\t%s\tTest %d:\tShould have seen the integer was not a palindrome.", succeed, testID)
					}
				}
			}
			t.Run(test.name, tf)
		}
	}
}

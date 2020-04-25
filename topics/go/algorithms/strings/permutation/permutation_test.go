/*
	// This is the API you need to build for these tests. You will need to
	// change the import path in this test to point to your code.

	package permutation

	// Is check if two strings are permutations.
	func Is(str1, str2 string) bool
*/

package permutation_test

import (
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/algorithms/strings/permutation"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestIsPermutation(t *testing.T) {
	tt := []struct {
		name    string
		input   string
		input2  string
		success bool
	}{
		{"empty", "", "", true},
		{"reverse", "god", "dog", true},
		{"diffsize", "god", "do", false},
		{"binary", "1001", "0110", true},
	}

	t.Log("Given the need to test permutation functionality.")
	{
		for testID, test := range tt {
			tf := func(t *testing.T) {
				t.Logf("\tTest %d:\tWhen checking the words %q and %q.", testID, test.input, test.input2)
				{
					got := permutation.Is(test.input, test.input2)
					switch test.success {
					case true:
						if !got {
							t.Fatalf("\t%s\tTest %d:\tShould have seen the string was a permutation.", failed, testID)
						}
						t.Logf("\t%s\tTest %d:\tShould have seen the string was a permutation.", succeed, testID)
					case false:
						if got {
							t.Fatalf("\t%s\tTest %d:\tShould have seen the string was not a permutation.", failed, testID)
						}
						t.Logf("\t%s\tTest %d:\tShould have seen the string was not a permutation.", succeed, testID)
					}
				}
			}
			t.Run(test.name, tf)
		}
	}
}

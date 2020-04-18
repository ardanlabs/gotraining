package strings_test

import (
	"testing"

	strings "github.com/ardanlabs/gotraining/topics/go/algorithms/strings/permuation"
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

	for _, test := range tt {
		tf := func(t *testing.T) {
			t.Log("Given the need to test permutation functionality.")
			{
				t.Logf("\tWhen checking the words %q and %q.", test.input, test.input2)
				{
					got := strings.IsPermutation(test.input, test.input2)
					switch test.success {
					case true:
						if !got {
							t.Fatalf("\t%s\tShould have seen the string was a permutation.", failed)
						}
						t.Logf("\t%s\tShould have seen the string was a permutation.", succeed)
					case false:
						if got {
							t.Fatalf("\t%s\tShould have seen the string was not a permutation.", failed)
						}
						t.Logf("\t%s\tShould have seen the string was not a permutation.", succeed)
					}
				}
			}
		}
		t.Run(test.name, tf)
	}
}

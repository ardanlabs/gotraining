package iseven_test

import (
	"github.com/ardanlabs/gotraining/topics/go/algorithms/fun/bits/iseven"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestIsEven(t *testing.T) {

	tt := []struct {
		name     string
		input    int
		expected bool
	}{
		{"one digit even number", 4, true},
		{"double digits even number", 14, true},
		{"one digit odd number", 5, false},
		{"double digit odd number", 15, false},
	}

	t.Log("Given the need to test IsEven functionality.")
	{
		for testID, test := range tt {
			t.Logf("\tTest %d:\tWhen checking the value %d.", testID, test.input)

			got := iseven.IsEven(test.input)

			if got != test.expected {
				t.Logf("\t%s\tTest %d:\tShould have gotten back if int was even", failed, testID)
				t.Fatalf("\t\tTest %d:\tGot %v, Excpeted %v", testID, got, test.expected)
			}
			t.Logf("\t%s\tTest %d:\tShould have gotten back if int was even", succeed, testID)
		}
	}

}

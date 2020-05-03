package iseven_test

import (
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/algorithms/bits/iseven"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestIsEven(t *testing.T) {

	tt := []struct {
		name     string
		input    int
		expected bool
	}{
		{"one-even", 4, true},
		{"double-even", 14, true},
		{"one-odd", 5, false},
		{"double-odd", 15, false},
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

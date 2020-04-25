package numbers_test

import (
	"testing"

	numbers "github.com/ardanlabs/gotraining/topics/go/algorithms/numbers/reverse"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestReverse(t *testing.T) {
	tt := []struct {
		name     string
		input    int
		expected int
	}{
		{"one", 1, 1},
		{"even", 5025, 5205},
		{"odd", 125, 521},
		{"negative", -502, -205},
	}

	t.Log("Given the need to test reverse functionality.")
	{
		for testID, test := range tt {
			t.Logf("\tTest %d:\tWhen checking the value %d.", testID, test.input)
			{
				got := numbers.Reverse(test.input)

				if got != test.expected {
					t.Logf("\t%s\tTest %d:\tShould have gotten back reverse integer.", failed, testID)
					t.Fatalf("\t\tTest %d:\tGot %v, Expected %v", testID, got, test.expected)
				}
				t.Logf("\t%s\tTest %d:\tShould have gotten back reverse integer.", succeed, testID)
			}
		}
	}
}

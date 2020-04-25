package numbers_test

import (
	numbers "github.com/ardanlabs/gotraining/topics/go/algorithms/numbers/reverse"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestReverse(t *testing.T) {
	tt := []struct {
		name     string
		input    int
		expected int
	}{
		{"one digit", 1, 1},
		{"even number of digits", 5025, 5205},
		{"even number of digits", 125, 521},
		{"negative digits", -502, -205},
	}

	for testID, test := range tt {

		got := numbers.Reverse(test.input)

		if got != test.expected {
			t.Logf("\t%s\tTest %d:\tShould have gotten back reverse integer.", failed, testID)
			t.Fatalf("\t\tTest %d:\tGot %v, Expected %v", testID, got, test.expected)
		}
		t.Logf("\t%s\tTest %d:\tShould have gotten back reverse integer.", succeed, testID)
	}
}

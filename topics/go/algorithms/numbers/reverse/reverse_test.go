/*
	// This is the API you need to build for these tests. You will need to
	// change the import path in this test to point to your code.

	package reverse

	// Reverse takes the specified integer and reverses it.
	func Reverse(num int) int
*/

package reverse_test

import (
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/algorithms/numbers/reverse"
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
				got := reverse.Reverse(test.input)

				if got != test.expected {
					t.Logf("\t%s\tTest %d:\tShould have gotten back reverse integer.", failed, testID)
					t.Fatalf("\t\tTest %d:\tGot %v, Expected %v", testID, got, test.expected)
				}
				t.Logf("\t%s\tTest %d:\tShould have gotten back reverse integer.", succeed, testID)
			}
		}
	}
}

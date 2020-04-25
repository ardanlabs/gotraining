/*
	// This is the API you need to build for these tests. You will need to
	// change the import path in this test to point to your code.

	package min

	// Min returns the minimum integer in the slice.
	func Min(n []int) (int, error)
*/

package min_test

import (
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/algorithms/slices/min"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestMax(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		expected int
		success  bool
	}{
		{"empty", []int{}, 0, false},
		{"nil", nil, 0, false},
		{"one", []int{10}, 10, true},
		{"even", []int{20, 30, 10, 50}, 10, true},
		{"odd", []int{30, 50, 10}, 10, true},
	}

	t.Log("Given the need to test Min functionality.")
	{
		for testID, test := range tt {
			tf := func(t *testing.T) {
				t.Logf("\tTest %d:\tWhen checking the %q state.", testID, test.name)
				{
					got, err := min.Min(test.input)
					switch test.success {
					case true:
						if err != nil {
							t.Fatalf("\t%s\tTest %d:\tShould be able to run Min without an error : %v", failed, testID, err)
						}
						t.Logf("\t%s\tTest %d:\tShould be able to run Min without an error.", succeed, testID)

					case false:
						if err == nil {
							t.Fatalf("\t%s\tTest %d:\tShould have seen an error for Min.", failed, testID)
						}
						t.Logf("\t%s\tTest %d:\tShould have seen an error for Min.", succeed, testID)
					}

					if got != test.expected {
						t.Logf("\t%s\tTest %d:\tShould have gotten back the right min value.", failed, testID)
						t.Fatalf("\t\tTest %d:\tGot %v, Expected %v", testID, got, test.expected)
					}
					t.Logf("\t%s\tTest %d:\tShould have gotten back the right min value.", succeed, testID)
				}
			}
			t.Run(test.name, tf)
		}
	}
}

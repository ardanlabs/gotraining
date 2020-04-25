/*
	// This is the API you need to build for these tests. You will need to
	// change the import path in this test to point to your code.

	package max

	// Max returns the maximum integer in the slice.
	func Max(n []int) (int, error)
*/

package max_test

import (
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/algorithms/slices/max"
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
		{"even", []int{10, 30}, 30, true},
		{"odd", []int{10, 50, 30}, 50, true},
	}

	t.Log("Given the need to test Max functionality.")
	{
		for testID, test := range tt {
			tf := func(t *testing.T) {
				t.Logf("\tTest %d:\tWhen checking the %q state.", testID, test.name)
				{
					got, err := max.Max(test.input)
					switch test.success {
					case true:
						if err != nil {
							t.Fatalf("\t%s\tTest %d:\tShould be able to run Max without an error : %v", failed, testID, err)
						}
						t.Logf("\t%s\tTest %d:\tShould be able to run Max without an error.", succeed, testID)

					case false:
						if err == nil {
							t.Fatalf("\t%s\tTest %d:\tShould have seen an error for Max.", failed, testID)
						}
						t.Logf("\t%s\tTest %d:\tShould have seen an error for Max.", succeed, testID)
					}

					if got != test.expected {
						t.Logf("\t%s\tTest %d:\tShould have gotten back the right max value.", failed, testID)
						t.Fatalf("\t\tTest %d:\tGot %v, Expected %v", testID, got, test.expected)
					}
					t.Logf("\t%s\tTest %d:\tShould have gotten back the right max value.", succeed, testID)
				}
			}
			t.Run(test.name, tf)
		}
	}
}

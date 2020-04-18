package slices_test

import (
	"testing"

	slices "github.com/ardanlabs/gotraining/topics/go/algorithms/slices/max"
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

	for _, test := range tt {
		tf := func(t *testing.T) {
			t.Log("Given the need to test Max functionality.")
			{
				t.Logf("\tWhen checking the %q state.", test.name)
				{
					got, err := slices.Max(test.input)
					switch test.success {
					case true:
						if err != nil {
							t.Fatalf("\t%s\tShould be able to run Max without an error : %v", failed, err)
						}
						t.Logf("\t%s\tShould be able to run Max without an error.", succeed)

					case false:
						if err == nil {
							t.Fatalf("\t%s\tShould have seen an error for Max.", failed)
						}
						t.Logf("\t%s\tShould have seen an error for Max.", succeed)
					}

					if got != test.expected {
						t.Logf("\t%s\tShould have gotten back the right max value.", failed)
						t.Fatalf("\t\tGot %v, Expected %v", got, test.expected)
					}
					t.Logf("\t%s\tShould have gotten back the right max value.", succeed)
				}
			}
		}
		t.Run(test.name, tf)
	}
}

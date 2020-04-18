package slices_test

import (
	"testing"

	slices "github.com/ardanlabs/gotraining/topics/go/algorithms/slices/min"
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

	for _, test := range tt {
		tf := func(t *testing.T) {
			t.Log("Given the need to test Min functionality.")
			{
				t.Logf("\tWhen checking the %q state.", test.name)
				{
					got, err := slices.Min(test.input)
					switch test.success {
					case true:
						if err != nil {
							t.Fatalf("\t%s\tShould be able to run Min without an error : %v", failed, err)
						}
						t.Logf("\t%s\tShould be able to run Min without an error.", succeed)

					case false:
						if err == nil {
							t.Fatalf("\t%s\tShould have seen an error for Min.", failed)
						}
						t.Logf("\t%s\tShould have seen an error for Min.", succeed)
					}

					if got != test.expected {
						t.Logf("\t%s\tShould have gotten back the right min value.", failed)
						t.Fatalf("\t\tGot %v, Expected %v", got, test.expected)
					}
					t.Logf("\t%s\tShould have gotten back the right min value.", succeed)
				}
			}
		}
		t.Run(test.name, tf)
	}
}

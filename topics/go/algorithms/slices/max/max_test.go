package slices_test

import (
	slices "github.com/ardanlabs/gotraining/topics/go/algorithms/slices/max"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestMax(t *testing.T) {

	maxTests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Slice with one element", []int{10}, 10},
		{"Slice with even number of elements", []int{10, 30}, 30},
		{"Slice with odd number of elements", []int{10, 50, 30}, 50},
	}

	for _, tt := range maxTests {
		got := slices.Max(tt.input)
		if got != tt.expected {
			t.Logf("\t%s\t Incorrect value returned: %d\n.", failed, tt.input)
			t.Fatalf("\t\tGot %d, Expected %d.", got, tt.expected)
		}
		t.Logf("\t%s\tCorrect value returend.", succeed)
	}
}

package slices_test

import (
	slices "github.com/ardanlabs/gotraining/topics/go/algorithms/slices/min"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestMax(t *testing.T) {

	minTests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Slice with one element", []int{10}, 10},
		{"Slice with even number of elements", []int{10, 30}, 10},
		{"Slice with odd number of elements", []int{10, 50, 30, 1, 5}, 1},
	}

	for _, tt := range minTests {
		got := slices.Min(tt.input)
		if got != tt.expected {
			t.Logf("\t%s\t Incorrect value returned: %d\n.", failed, tt.input)
			t.Fatalf("\t\tGot %d, Expected %d.", got, tt.expected)
		}
		t.Logf("\t%s\tCorrect value returend.", succeed)
	}
}

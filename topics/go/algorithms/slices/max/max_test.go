package slices_test

import (
	"testing"

	slices "github.com/ardanlabs/gotraining/topics/go/algorithms/slices/max"
)

func TestMax(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		shouldError bool
	}{
		"Empty Slice":                        {[]int{}, 0, true},
		"nil Slice":                          {[]int(nil), 0, true},
		"Slice with one element":             {[]int{10}, 10, false},
		"Slice with even number of elements": {[]int{10, 30}, 30, false},
		"Slice with odd number of elements":  {[]int{10, 50, 30}, 50, false},
	}

	t.Parallel()
	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			got, err := slices.Max(test.input)
			if err == nil && test.shouldError {
				t.Fatalf("slices.Max(%#v) returned a nil error; expected a non-nil error value", test.input)
			}
			if err != nil {
				if test.shouldError == false {
					t.Fatalf("slices.Max(%#v) returned error %v", test.input, err)
				}
				return
			}
			if got != test.expected {
				t.Fatalf("slices.Max(%#v) returned %d; expected %d.", test.input, got, test.expected)
			}
		})
	}
}

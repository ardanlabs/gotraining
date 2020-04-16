package slices_test

import (
	"testing"

	slices "github.com/ardanlabs/gotraining/topics/go/algorithms/slices/max"
)

func TestMax(t *testing.T) {
	maxTests := map[string]struct {
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
	for name, tt := range maxTests {
		name, tt = name, tt
		t.Run(name, func(t *testing.T) {
			got, err := slices.Max(tt.input)
			if err != nil {
				if tt.shouldError == false {
					t.Fatalf("slices.Max(%#v) returned error %v", tt.input, err)
				}
				return
			}
			if got != tt.expected {
				t.Fatalf("slices.Max(%#v) returned %d; expected %d.", tt.input, got, tt.expected)
			}
		})
	}
}

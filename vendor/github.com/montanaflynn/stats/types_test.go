package stats

import (
	"reflect"
	"testing"
)

var fd = Float64Data{-10, -10.001, 5, 1.1, 2, 3, 4.20, 5}

func TestInterfaceMethods(t *testing.T) {
	// Test Get
	a := fd.Get(1)
	if a != -10.001 {
		t.Errorf("Get(2) => %.1f != %.1f", a, -10.001)
	}

	// Test Len
	l := fd.Len()
	if l != 8 {
		t.Errorf("Len() => %v != %v", l, 8)
	}

	// Test Less
	b := fd.Less(0, 5)
	if b != true {
		t.Errorf("Less() => %v != %v", b, true)
	}

	// Test Swap
	fd.Swap(0, 2)
	if fd.Get(0) != 5 {
		t.Errorf("Len() => %v != %v", l, 8)
	}
}

func TestHelperMethods(t *testing.T) {

	// Test Min
	m, _ := fd.Min()
	if m != -10.001 {
		t.Errorf("Min() => %v != %v", m, -10.001)
	}

	// Test Max
	m, _ = fd.Max()
	if m != 5 {
		t.Errorf("Max() => %v != %v", m, 5)
	}

	// Test Sum
	m, _ = fd.Sum()
	if m != 0.2990000000000004 {
		t.Errorf("Sum() => %v != %v", m, 0.2990000000000004)
	}

	// Test Mean
	m, _ = fd.Mean()
	if m != 0.03737500000000005 {
		t.Errorf("Mean() => %v != %v", m, 0.03737500000000005)
	}

	// Test Median
	m, _ = fd.Median()
	if m != 2.5 {
		t.Errorf("Median() => %v != %v", m, 2.5)
	}

	// Test Mode
	mo, _ := fd.Mode()
	if !reflect.DeepEqual(mo, []float64{5.0}) {
		t.Errorf("Mode() => %.1f != %.1f", mo, []float64{5.0})
	}

}

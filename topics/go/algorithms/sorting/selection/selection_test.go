package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	got := generateList(1e2)
	want := make([]int, 1e2)
	copy(want, got)

	sort.Ints(want)
	selectionSort(got)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	// b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		b.StopTimer()
		numbers := generateList(1e2)
		b.StartTimer()
		selectionSort(numbers)
	}
}

func BenchmarkStdSort(b *testing.B) {
	// b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		b.StopTimer()
		numbers := generateList(1e2)
		b.StartTimer()
		sort.Ints(numbers)
	}
}

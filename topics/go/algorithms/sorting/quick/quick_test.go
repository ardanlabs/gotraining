package quicksort

import (
	"math/rand"
	"sort"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

var snum []int

// generateList is for generate a random list of numbers.
func generateList(totalNumbers int) []int {
	numbers := make([]int, totalNumbers)
	for i := 0; i < totalNumbers; i++ {
		numbers[i] = rand.Intn(totalNumbers)
	}
	return numbers
}

// TestQuickSort to test our quick sort algorithm with different data.
func TestQuickSort(t *testing.T) {
	dataNumber := []struct {
		randomList []int
	}{
		{randomList: generateList(100)},
		{randomList: generateList(985)},
		{randomList: generateList(852)},
		{randomList: generateList(1000)},
		{randomList: generateList(1)},
		{randomList: generateList(9999)},
		{randomList: []int{-1}},
		{randomList: []int{}},
		{randomList: []int{-1, 2, 2, 2, 2, 22, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10000000}},
		{randomList: []int{99999999999, -1, 2, 2, 2, 2, 22, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10000000}},
		{randomList: []int{-99999999999, -1, 2, 2, 2, 2, 22, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10000000}},
	}

	t.Run("Quick Sort Random Numbers", func(t *testing.T) {
		t.Log("Start the testing quick sort for random numbers.")
		{
			for _, x := range dataNumber {
				result := quickSort(x.randomList, 0, len(x.randomList)-1)

				if !sort.IntsAreSorted(result) {
					t.Fatalf("\t%s\t\n Got: \n\t %v \n", failed, result)
				}
				t.Logf("\t%s\tEverything is looks fine", succeed)
			}
		}
	})
}

// BenchmarkQuickSort a simple benchmark for the quick sort algorithm.
func BenchmarkQuickSort(b *testing.B) {
	var sn []int
	list := generateList(1000)

	for i := 0; i < b.N; i++ {
		sn = quickSort(list, 0, len(list)-1)
	}

	snum = sn
}

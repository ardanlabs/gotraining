package heap_test

import (
	"math/rand"
	"sort"
	"testing"

	heapsort "github.com/ardanlabs/gotraining/topics/go/algorithms/sorting/heap"
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

// TestHeapSort to test our heap sort algorithm with different data.
func TestHeapSort(t *testing.T) {
	dataNumber := []struct {
		randomList []int
	}{
		{randomList: generateList(8)},
		{randomList: generateList(985)},
		{randomList: generateList(852)},
		{randomList: generateList(1000)},
		{randomList: generateList(1)},
		{randomList: generateList(9999)},
		{randomList: []int{-1}},
		{randomList: []int{}},
		{randomList: []int{-1, 2, 2, 2, 2, 22, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10000000}},
	}

	t.Run("Heap Sort Random Numbers", func(t *testing.T) {
		t.Log("Start the testing heap sort for random numbers.")
		{
			for x := range dataNumber {
				result := heapsort.HeapSort(dataNumber[x].randomList)

				if !sort.IntsAreSorted(result) {
					t.Fatalf("\t%s\t \n Got: \n\t %v \n", failed, result)
				}
				t.Logf("\t%s\tEverything is looks fine", succeed)
			}
		}
	})
}

// BenchmarkHeapSort a simple benchmark for the heap sort algorithm.
func BenchmarkHeapSort(b *testing.B) {
	var sn []int
	list := generateList(1000)

	for i := 0; i < b.N; i++ {
		sn = heapsort.HeapSort(list)
	}

	snum = sn
}

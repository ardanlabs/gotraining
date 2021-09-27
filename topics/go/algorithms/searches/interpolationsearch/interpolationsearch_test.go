package interpolationsearch

import (
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

// generateSortList create a sorted list of number.
func generateSortList(value int) (list []int) {
	if value <= 0 {
		return list
	}

	for i := 1; i < value; i++ {
		list = append(list, i)
	}

	return list
}

// TestInterpolationSearch a function to test interpolationSearchIterative and interpolationSearchRecursive algorithms.
func TestInterpolationSearch(t *testing.T) {
	// generate a list of numbers between 1 and 10000.
	listData := generateSortList(10000)

	// list of data for testing.
	data := []struct {
		list   []int
		find   int
		expect int
	}{
		{listData, 50, 49},
		{listData, 159, 158},
		{listData, 1, 0},
		{listData, 500, 499},
		{listData, 9999, 9998},
		{[]int{2}, 2, 0},
		{[]int{-1, 2, 0}, 3, -1},
		{generateSortList(-1), -999, -1},
	}

	t.Run("Interpolation Search Iterative", func(t *testing.T) {
		t.Log("Start the testing interpolation search in iterative way.")
		{
			for i := range data {

				index := interpolationSearchIterative(data[i].list, data[i].find)

				if index != data[i].expect {
					t.Fatalf("\t%s\tExpected %d, but got %d", failed, data[i].expect, index)
				}
				t.Logf("\t%s\tEverything is looks fine, test %d", succeed, i)
			}
		}
	})

	t.Run("Interpolation Search Recursive", func(t *testing.T) {
		t.Log("Start the testing interpolation search in recursive way.")
		{
			for i := range data {
				index := interpolationSearchRecursive(data[i].list, data[i].find, 0, len(data[i].list)-1)

				if index != data[i].expect {
					t.Fatalf("\t%s\tExpected %d, but got %d", failed, data[i].expect, index)
				}
				t.Logf("\t%s\tEverything is looks fine, test %d", succeed, i)
			}
		}
	})

}

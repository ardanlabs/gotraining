package interpolationsearch

import (
"testing"
)

//	generateSortList create a sort list numbers
func generateSortList(value int) (list []int) {
	if value <= 0 {
		return list
	}

	for i := 1; i < value; i++ {
		list = append(list, i)
	}

	return list
}

// 	TestInterpolationSearch
func TestInterpolationSearch(t *testing.T) {
	// generate list of 10000 numbers
	listData := generateSortList(10000)

	// 	list of data to check
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
		{[]int{-1,2,0}, 3, -1},
		{generateSortList(-1), -999, -1},
	}

	t.Run("Interpolation Search Iterative", func(t *testing.T) {
		for i := range data {
			index := interpolationSearchIterative(data[i].list, data[i].find)

			if index != data[i].expect {
				t.Errorf("expected %d, but got %d", data[i].expect, index)
			}
		}
	})

	t.Run("Interpolation Search Recursive", func(t *testing.T) {
		for i := range data {
			index := interpolationSearchRecursive(data[i].list, data[i].find, 0, len(data[i].list)-1)

			if index != data[i].expect {
				t.Errorf("expected %d, but got %d", data[i].expect, index)
			}
		}
	})

}

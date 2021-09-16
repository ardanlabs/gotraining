package linearsearch

// linearSearchIterative is the algorithm that search the value inside the array index by index
// the worst case of this algorithm is O(n)
// the best case of this algorithm is O(1)
func linearSearchIterative(list []int, target int) int {

	if len(list) <= 0{
		return -1
	}

	for i := range list {
		if list[i] == target {
			return i
		}
	}

	return -1
}

// linearSearchRecursive is the algorithm that search the value inside the array index by index with recall the function
func linearSearchRecursive(list []int, target int, index int) int {

	if len(list) <= index {
		return -1
	}

	// check if the value of array is equal with target
	// return the index
	if list[index] == target {
		return index
	}

	if len(list) >= index {
		index++ // increase by one
		return linearSearchRecursive(list, target, index)
	}

	// if value is not found I returned -1
	return -1
}

// doubleLinearSearchIterative is the algorithm that search the value inside the array in both side
// the worst case of this algorithm is O(n)
// the best case of this algorithm is O(1)
func doubleLinearSearchIterative(list []int, target int) int {

	var low int           // first index
	high := len(list) - 1 // last index

	if len(list) <= 0{
		return -1
	}

	// continue loop until low is smaller or equal to high value
	for low <= high {

		if list[low] == target {
			return low
		}

		if list[high] == target {
			return high
		}

		low++
		high--
	}

	// if value is not found I returned -1
	return -1
}

// doubleLinearSearchRecursive is the algorithm that search the value inside the array in both side
func doubleLinearSearchRecursive(list []int, target int, low int, high int) int {

	if len(list) <= 0{
		return -1
	}

	if list[low] == target {
		return low
	}

	if list[high] == target {
		return high
	}

	if low <= high {
		low++
		high--
		return doubleLinearSearchRecursive(list, target, low, high)
	}

	// if value is not found I returned -1
	return -1
}

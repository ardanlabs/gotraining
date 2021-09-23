package interpolationsearch

// interpolationSearchIterative this algorithm is like binary search but it improved
// the algorithm is not start from the middle of array it has own position
// to calculate the position it use this => low + int(float64(high-low) / float64(list[high]-list[low])) * (target - list[low])
// low is the first index of array
// high is the last index of array
func interpolationSearchIterative(list []int, target int) int {
	var low int
	high := len(list) - 1

	if len(list) <= 0 {
		return -1
	}

	for low <= high && target >= list[low] && target <= list[high] {

		// calculate the position
		position := low + int(float64(high-low)/float64(list[high]-list[low]))*(target-list[low])

		if list[position] == target {
			return position
		}

		if list[position] > target {
			high = position - 1
		}

		if list[position] < target {
			low = position + 1
		}
	}

	return -1
}

// interpolationSearchRecursive the same as interpolationSearchIterative with the recursive method
func interpolationSearchRecursive(list []int, target int, low int, high int) int {

	if len(list) <= 0 {
		return -1
	}

	if low <= high && target >= list[low] && target <= list[high] {

		position := low + int(float64(high-low)/float64(list[high]-list[low]))*(target-list[low])

		if list[position] == target {
			return position
		}

		if list[position] > target {
			return interpolationSearchRecursive(list, target, low, position-1)
		}

		if list[position] < target {
			return interpolationSearchRecursive(list, target, position+1, high)
		}
	}

	return -1
}

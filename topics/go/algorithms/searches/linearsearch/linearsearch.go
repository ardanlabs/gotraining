package linearsearch

// linearSearchIterative is the algorithm that search the value inside the array index by index
// the worst case of this algorithm is O(n)
// the best case of this algorithm is O(1)
func linearSearchIterative(list []int, target int) int {

	for i := range list{
		if list[i] == target {
			return i
		}
	}

	return -1
}

// linearSearchRecursive is the algorithm that search the value inside the array index by index with recall the function
func linearSearchRecursive(list []int, target int, index int) int {

	if len(list) <= index{
		return -1
	}

	// check if the value of array is equal with target
	// return the index
	if list[index] == target {
		return index
	}

	if len(list) >= index{
		index++ // increase by one
		return linearSearchRecursive(list, target,index)
	}

	// if value is not found I returned -1
	return -1
}
package binarySearch

func binarySearch(slice []int, element int) int {
	low := 0
	high := len(slice) - 1

	for low <= high {
		mid := (low + high) / 2
		if element == slice[mid] {
			return mid + 1
		} else if element > slice[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

package interpolationSearch

func InterpolationSearch(slice []int, element int) int {
	low := 0
	high := len(slice) - 1
	for low <= high && element >= slice[low] && element <= slice[high] {
		pos := low + ((high-low)*(element-slice[low]))/(slice[high]-slice[low])
		if slice[pos] == element {
			return pos + 1
		} else if element > slice[pos] {
			low = pos + 1
		} else {
			high = pos - 1
		}
	}

	return -1
}

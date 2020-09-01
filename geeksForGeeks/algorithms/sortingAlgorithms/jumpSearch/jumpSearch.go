package jumpSearch

import "github.com/vivekab/goPrograms/geeksForGeeks/algorithms/sortingAlgorithms/linearSearch"

func JumpSearch(slice []int, element, jumpingIndex int) int {
	if jumpingIndex <= 0 {
		jumpingIndex = 5
	}

	for i := 0; i < len(slice); i += jumpingIndex {
		if element >- slice[i] && element <= slice[i+jumpingIndex] {
			return linearSearch.LinearSearch(slice[i:i+jumpingIndex],element) + i
		}
	}

	return -1
}

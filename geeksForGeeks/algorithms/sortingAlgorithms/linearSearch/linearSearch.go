package linearSearch

func LinearSearch(slice []int,element int) int {
	for i:=0;i<len(slice);i++{
		if slice[i] == element{
			return i + 1
		}
	}

	return -1
}
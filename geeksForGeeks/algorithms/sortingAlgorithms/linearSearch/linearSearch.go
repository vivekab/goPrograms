package main

import "fmt"

func main()  {
	fmt.Println("Linear Search:",linearSearch([]int{1,2,3,4,5,6,7,22,8,9,10,11,12,13,15,16,16},10))
	fmt.Println("Linear Search:",linearSearch([]int{1,2,3,4,5,6,7,22,8,9,10,11,12,13,15,16,16},0))
}

func linearSearch(slice []int,element int) int {
	for i:=0;i<len(slice);i++{
		if slice[i] == element{
			return i + 1
		}
	}

	return -1
}
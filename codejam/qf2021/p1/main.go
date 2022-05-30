package main

import (
	"fmt"
)

func main() {
	tc := 0
	fmt.Scanf("%d", &tc)
	out := make([]int, tc)
	for i := 0; i < tc; i++ {
		n := 0
		fmt.Scanln(&n)
		arr := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &arr[j])
		}
		out[i] = compute(arr, n)
	}

	for i := 0; i < len(out)-1; i++ {
		fmt.Printf("Case #%d: %d\n", i+1, out[i])
	}
	fmt.Printf("Case #%d: %d", len(out), out[len(out)-1])

}
func compute(arr []int, n int) int {
	count := 0
	for i := range arr {
		_, j := findMin(arr[i:])
		j += i
		if i != j {
			fmt.Println(arr,i,j)
			reverse(&arr, i, j)
			count += j - i + 1
		}
	}
	return count
}
func reverse(arr *[]int, i, j int) {
	if j == i+1 {
		temp := (*arr)[i]
		(*arr)[i] = (*arr)[j]
		(*arr)[j] = temp
		return
	}
	mid := (i + j) / 2
	for k := 0; k <= mid; k++ {
		temp := (*arr)[k+i]
		(*arr)[k+i] = (*arr)[j-k]
		(*arr)[j-k] = temp
	}
}

func findMin(arr []int) (min, index int) {
	min = 9999
	index = -1
	for i := range arr {
		if arr[i] < min {
			min = arr[i]
			index = i
		}
	}
	return min, index
}

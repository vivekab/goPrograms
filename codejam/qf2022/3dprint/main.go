package main

import "fmt"

func main() {
	t := 0
	fmt.Scanf("%d", &t)
	in := make([][]int, t*3)
	for i := range in {
		in[i] = make([]int, 4)
	}
	for i := 0; i < t*3; i++ {
		fmt.Scanf("%d", &in[i][0])
		fmt.Scanf("%d", &in[i][1])
		fmt.Scanf("%d", &in[i][2])
		fmt.Scanf("%d", &in[i][3])
	}
	for i := 0; i < t; i++ {
		fmt.Printf("Case #%d: ", i+1)
		res := threedPrint(in, i*3)
		if res != nil {
			fmt.Printf("%d %d %d %d", res[0], res[1], res[2], res[3])
		} else {
			fmt.Printf("IMPOSSIBLE")
		}
		if i != t-1 {
			fmt.Println("")
		}
	}
}
func threedPrint(in [][]int, ind int) []int {
	min := make([]int, 4)
	sol := make([]int, 4)
	for i := range min {
		min[i] = 1000001
	}
	for j := 0; j < 4; j++ {
		for i := 0; i < 3; i++ {
			if in[ind+i][j] < min[j] {
				min[j] = in[ind+i][j]
			}
		}
	}
	total := 1000000
	for i := range min {
		if total <= min[i] {
			sol[i] = total
			total -= min[i]
			break
		}
		sol[i] = min[i]
		total -= min[i]
	}
	if total > 0 {
		return nil
	}
	return sol
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	t := 0
	fmt.Scanf("%d", &t)
	in := make([][]int, t)
	for i := 0; i < t; i++ {
		n := 0
		fmt.Scanf("%d", &n)
		in[i] = make([]int, n+1)
		in[i][0] = n
		for j := 1; j <= n; j++ {
			fmt.Scanf("%d", &in[i][j])
		}
	}
	for i := 0; i < t; i++ {
		fmt.Printf("Case #%d: ", i+1)
		res := longestConsecutiveSubset(in[i])
		fmt.Printf("%d", res)
		if i != t-1 {
			fmt.Println("")
		}
	}
}

func longestConsecutiveSubset(in []int) int {
	m := map[int]int{}
	max := 0
	n := in[0]

	for i := 1; i <= n; i++ {
		if in[i] > max {
			max = in[i]
		}
		m[in[i]]++
	}

	keys := make([]int, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	numberIndex := 0

	for i := range keys {
		req := keys[i] - numberIndex
		consumable := min(req, m[keys[i]])
		numberIndex += consumable
		m[keys[i]] = 0
	}
	return numberIndex
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

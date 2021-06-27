package main

func candy(ratings []int) int {
	l2r := make([]int, len(ratings))
	r2l := make([]int, len(ratings))
	for i := 0; i < len(l2r); i++ {
		l2r[i] = 1
		r2l[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			l2r[i] = l2r[i-1] + 1
		}
		if ratings[len(ratings)-1-i] > ratings[len(ratings)-i] {
			r2l[len(ratings)-1-i] = r2l[len(ratings)-i] + 1
		}
	}
	total := 0
	for i := 0; i < len(ratings); i++ {
		total += max(l2r[i], r2l[i])
	}
	return total
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

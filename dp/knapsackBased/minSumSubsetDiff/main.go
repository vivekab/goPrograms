package main

func solve(nums []int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	halfSum := 0
	if sum%2 == 0 {
		halfSum = sum / 2
	} else {
		halfSum = sum/2 + 1
	}

	t := make([][]bool, len(nums)+1)
	for i := range t {
		t[i] = make([]bool, halfSum+1)
	}

	for i := range t {
		for j := range t[i] {
			if i == 0 {
				t[i][j] = false
			}
			if j == 0 {
				t[i][j] = true
			}
		}
	}

	for i := 1; i < len(t); i++ {
		for j := 1; j < len(t[i]); j++ {
			if nums[i-1] <= j {
				t[i][j] = t[i-1][j-nums[i-1]] || t[i-1][j]
			} else {
				t[i][j] = t[i-1][j]
			}
		}
	}
	maxS1 := -1

	i := len(nums)
	for j := range t[i] {
		if t[i][j] && j > maxS1 {
			maxS1 = j
		}
	}

	s2 := sum - maxS1

	return abs(s2 - maxS1)
}
func abs(num int) int {
	if num < 0 {
		return -1 * num
	}
	return num
}

package main

import "fmt"

func main() {
	t := 0
	fmt.Scanf("%d", &t)
	in := make([][]int, t)
	for i := range in {
		in[i] = make([]int, 2)
	}
	for i := 0; i < t; i++ {
		fmt.Scanf("%d", &in[i][0])
		fmt.Scanf("%d", &in[i][1])
	}
	for i := range in {
		fmt.Printf("Case #%d:\n", i+1)
		printPunchCard(in[i][0], in[i][1])
		if i != len(in)-1{
			fmt.Println("")
		}
	}
}

func printPunchCard(r, c int) {
	for i := 1; i <= (2*r + 1); i++ {
		for j := 1; j <= (2*c + 1); j++ {
			if i == 1 && j == 1 {
				fmt.Printf(".")
				continue
			}
			if i == 1 && j == 2 {
				fmt.Printf(".")
				continue
			}
			if j == 1 && i == 2 {
				fmt.Printf(".")
				continue
			}
			if i&1 == 1 {
				if j&1 == 1 {
					fmt.Printf("+")
				} else {
					fmt.Printf("-")
				}
			} else {
				if j&1 == 1 {
					fmt.Printf("|")
				} else {
					fmt.Printf(".")
				}
			}
		}
		if i != (2*r+1){
			fmt.Println("")
		}
	}
}

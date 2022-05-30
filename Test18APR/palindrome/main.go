package main

import "fmt"

func main() {
	a := "aabcaa"

	revStr := reverse(a)

	if revStr != a{
		fmt.Println("Not a palindrome")
	}else{
		fmt.Println("Palindrome")
	}
}

func reverse(a string) string {
	rev := ""
	for i := len(a) - 1; i >= 0; i-- {
		rev += string(a[i])
	}
	return rev
}

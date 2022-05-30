package main

import "fmt"

var x = map[string]bool{}

func mock(a string) bool {
	if a == "google.com" {
		return true
	}
	return false
}

func main() {
	stringSlice := make([]string,0,3)
	stringSlice = []string{
		"google.com",
		"facebook.com",
		"twitter.com",
	}
	val := evaluateSlice(stringSlice)

	fmt.Println(val)
}

func evaluateSlice(stringSlice []string) map[string]bool {
	sliceMap := map[string]bool{}

	for i := range stringSlice {
		sliceMap[stringSlice[i]] = mock(stringSlice[i])
	}

	return sliceMap
}

package main

import "log"

func main() {
	val := numSubarrayProductLessThanK([]int{10,5,2,1,6},100)
	if val != 12{
		log.Fatal("Failed")
	}
	log.Println("Success")
}
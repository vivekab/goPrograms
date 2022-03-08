package main

import "testing"

func Test(t *testing.T) {
	arr := []int{5,6,2,1}
	count := make([]int,len(arr))
	m := NewNode(5,0)
	for i:=1;i<len(arr);i++{
		m.insert(arr[i],i,count)
	}

}

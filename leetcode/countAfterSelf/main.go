package main

type node struct {
	left      *node
	right     *node
	val       int
	leftNodes int
	index     int
}

func NewNode(rootElement int, index int) *node {
	return &node{
		left:  nil,
		right: nil,
		val:   rootElement,
		index: index,
	}
}

func (n *node) insert(e int, index int, count []int) {
	prev := n
	cur := n
	for cur != nil {
		if e > cur.val {
			prev, cur = cur, cur.right
		} else {
			if e < cur.val {
				cur.leftNodes++
				count[cur.index]++
			}
			prev, cur = cur, cur.left
		}
	}
	eNode := NewNode(e, index)
	if e > prev.val {
		prev.right = eNode
	} else {
		prev.left = eNode
	}
}

func (n *node) balance() int {
	if n.left == nil && n.right == nil{
		return 0
	}
	if n.left == nil{
		return -1-n.right.height(0)
	}
	if n.right == nil{
		return n.left.height(0)+1
	}
	return n.left.height(0)-n.right.height(0)
}

func (n *node) height(h int) int {
	if n.left == nil && n.right == nil {
		return h
	}
	if n.left == nil {
		return n.right.height(h + 1)
	}
	if n.right == nil {
		return n.left.height(h + 1)
	}
	return max(n.left.height(h+1), n.right.height(h+1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func countSmaller(nums []int) []int {
	if len(nums) == 1 {
		return []int{0}
	}
	m := NewNode(nums[0],0)
	count := make([]int, len(nums))
	for i:=1;i<len(nums);i++ {
		m.insert(nums[i],i,count)
	}
	return count
}

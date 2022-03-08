package main

func numSubarrayProductLessThanK(nums []int, k int) int {
	result := 0
	if k == 0 {
		return result
	}
	que := NewQueue()
	for _, v := range nums {
		que.insert(v)
		if que.product() >= k {
			que.remove()
		}
		if que.isProductLessThanK(k) {
			result += que.len()
		}
	}
	return result
}

type queue struct {
	front int
	rear  int
	arr   []int
	prod  int
}

func NewQueue() *queue {
	return &queue{
		front: -1,
		rear:  0,
		arr:   []int{},
		prod:  1,
	}
}

func (q *queue) len() int {
	return q.rear - q.front - 1
}

func (q *queue) insert(a int) {
	if q.front > 10 {
		for i, o := q.front+1, 0; i < q.rear; i, o = i+1, o+1 {
			q.arr[o] = q.arr[i]
		}
		q.front -= 10
		q.rear -= 10
	}
	if len(q.arr) < q.rear+1 {
		q.arr = append(q.arr, a)
	} else {
		q.arr[q.rear] = a
	}
	q.rear++
	q.prod *= a
}

func (q *queue) remove() int {
	if q.front+1 == q.rear {
		return -1
	}
	q.front++
	val := q.arr[q.front]
	q.prod = q.prod / val
	return val
}
func (q *queue) product() int {
	return q.prod
}
func (q *queue) isProductLessThanK(k int) bool {
	return q.prod < k
}

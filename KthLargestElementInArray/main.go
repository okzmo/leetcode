package main

import (
	"container/heap"
	"fmt"
)

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(item interface{}) {
	*h = append(*h, item.(int))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func findKthLargest(nums []int, k int) int {
	h := &maxHeap{}
	heap.Init(h)
	for _, v := range nums {
		heap.Push(h, v)
	}

	var item interface{}
	for i := 0; i < k; i++ {
		item = heap.Pop(h)
	}

	return item.(int)
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(nums, 2))
}

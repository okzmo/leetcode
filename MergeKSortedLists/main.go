package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// USING MERGE//

func merge(l, r *ListNode) *ListNode {
	head := &ListNode{}
	tail := head

	for l != nil && r != nil {
		if l.Val < r.Val {
			tail.Next = l
			l = l.Next
		} else {
			tail.Next = r
			r = r.Next
		}
		tail = tail.Next
	}

	for l != nil {
		tail.Next = l
		l = l.Next
		tail = tail.Next
	}

	for r != nil {
		tail.Next = r
		r = r.Next
		tail = tail.Next
	}

	return head.Next
}

func mergeKLists2(lists []*ListNode) *ListNode {
	n := len(lists)

	if n == 0 {
		return nil
	}

	curr := lists[0]
	if n == 1 {
		return curr
	}

	for i := 1; i < n; i++ {
		curr = merge(curr, lists[i])
	}

	return curr
}

//
//

type MinHeap []*ListNode

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	var minHeap MinHeap
	heap.Init(&minHeap)
	for _, list := range lists {
		if list != nil {
			heap.Push(&minHeap, list)
		}
	}

	head := &ListNode{}
	tail := head

	for minHeap.Len() > 0 {
		node := heap.Pop(&minHeap).(*ListNode)
		tail.Next = node
		if node.Next != nil {
			heap.Push(&minHeap, node.Next)
		}
		tail = node
	}

	return head.Next
}

func main() {
	arr1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	arr2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	arr3 := &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 7}}}
	arr4 := &ListNode{Val: 4, Next: &ListNode{Val: 6, Next: &ListNode{Val: 9}}}
	fullArr := []*ListNode{arr1, arr2, arr3, arr4}

	final := mergeKLists(fullArr)

	for final != nil {
		fmt.Println(final.Val)
		final = final.Next
	}
}

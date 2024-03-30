package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Heap struct {
	items []*ListNode
}

func left(i int) int {
	return i*2 + 1
}

func right(i int) int {
	return i*2 + 2
}

func parent(i int) int {
	return (i - 1) / 2
}

func (h *Heap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *Heap) heapifyDown(idx int) {
	lastIdx := len(h.items) - 1
	l, r := left(idx), right(idx)
	childToCmp := 0

	for l <= lastIdx {
		if l == lastIdx {
			childToCmp = l
		} else if h.items[l].Val < h.items[r].Val {
			childToCmp = l
		} else {
			childToCmp = r
		}

		if h.items[childToCmp].Val < h.items[idx].Val {
			h.swap(childToCmp, idx)
			idx = childToCmp
			l, r = left(idx), right(idx)
		} else {
			break
		}
	}
}

func (h *Heap) heapifyUp(idx int) {
	for h.items[idx].Val < h.items[parent(idx)].Val {
		h.swap(idx, parent(idx))
		idx = parent(idx)
	}
}

func (h *Heap) Pop() *ListNode {
	n := len(h.items)
	item := h.items[0]
	h.items[0] = h.items[n-1]
	h.items = h.items[:n-1]
	h.heapifyDown(0)

	return item
}

func (h *Heap) Push(item *ListNode) {
	h.items = append(h.items, item)
	h.heapifyUp(len(h.items) - 1)
}

func (h Heap) Len() int {
	return len(h.items)
}

func mergeKLists(lists []*ListNode) *ListNode {
	var minHeap Heap
	for _, list := range lists {
		if list != nil {
			minHeap.Push(list)
		}
	}

	head := &ListNode{}
	tail := head

	for minHeap.Len() > 0 {
		node := minHeap.Pop()
		tail.Next = node
		if node.Next != nil {
			minHeap.Push(node.Next)
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

package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	curr := head
	var prev *ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func main() {
	head := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}

	fmt.Println(reverseList(&head))
}

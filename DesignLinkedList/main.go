package main

type Node struct {
	Prev *Node
	Next *Node
	Val  int
}

type MyLinkedList struct {
	Head *Node
	Tail *Node
	size int
}

func Constructor() MyLinkedList {
	head := &Node{Val: -1}
	tail := &Node{Val: -1}
	head.Next = tail
	tail.Prev = head
	return MyLinkedList{
		Head: head,
		Tail: tail,
		size: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}

	curr := this.Head
	for i := 0; i <= index; i++ {
		curr = curr.Next
	}

	return curr.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{
		Val:  val,
		Next: this.Head.Next,
		Prev: this.Head,
	}

	this.Head.Next.Prev = node
	this.Head.Next = node

	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{
		Val:  val,
		Next: this.Tail,
		Prev: this.Tail.Prev,
	}

	this.Tail.Prev.Next = node
	this.Tail.Prev = node

	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}

	if index <= 0 {
		this.AddAtHead(val)
		return
	}

	if index == this.size {
		this.AddAtTail(val)
		return
	}

	curr := this.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}

	node := &Node{
		Val:  val,
		Next: curr.Next,
		Prev: curr,
	}

	curr.Next.Prev = node
	curr.Next = node

	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size {
		return
	}

	curr := this.Head
	for i := 0; i <= index; i++ {
		curr = curr.Next
	}

	curr.Prev.Next = curr.Next
	curr.Next.Prev = curr.Prev

	this.size--
}

func main() {
	myLinkedList := Constructor()
	myLinkedList.AddAtHead(1)
	myLinkedList.AddAtTail(3)
	myLinkedList.AddAtIndex(1, 2) // linked list becomes 1->2->3
	myLinkedList.Get(1)           // return 2
	myLinkedList.DeleteAtIndex(1) // linked list becomes 1->3
	// myLinkedList.Get(1) // return 3
}

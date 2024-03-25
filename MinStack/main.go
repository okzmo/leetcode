package main

import "fmt"

type Node struct {
	val     int
	currMin int
}

type MinStack struct {
	stack []Node
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	node := Node{val: val, currMin: val}
	if len(this.stack) > 0 {
		prevNode := this.stack[len(this.stack)-1]
		if prevNode.currMin > val {
			node.currMin = val
		} else {
			node.currMin = prevNode.currMin
		}
	}

	this.stack = append(this.stack, node)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].val
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].currMin
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}

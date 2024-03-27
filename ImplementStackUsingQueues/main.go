package main

import (
	"container/list"
	"fmt"
)

type MyStack struct {
	q *list.List
}

func Constructor() MyStack {
	return MyStack{
		q: list.New(),
	}
}

func (this *MyStack) Push(x int) {
	this.q.PushBack(x)
	for i := 0; i < this.q.Len(); i++ {
		front := this.q.Front()
		this.q.Remove(front)
		this.q.PushBack(front.Value)
	}
}

func (this *MyStack) Pop() int {
	back := this.q.Back()
	this.q.Remove(back)

	return back.Value.(int)
}

func (this *MyStack) Top() int {
	return this.q.Back().Value.(int)
}

func (this *MyStack) Empty() bool {
	return this.q.Len() == 0
}

func main() {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	fmt.Println(stack.q.Back())
	item := stack.Pop()
	fmt.Println(item)
	fmt.Println(stack.q.Back())
	// item := stack.Top()
	// fmt.Println(item, stack)
	// item = stack.Pop()
	// fmt.Println(item, stack)
	// fmt.Println(stack.Empty())
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

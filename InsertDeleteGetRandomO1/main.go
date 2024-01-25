package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	set map[int]bool
	el  []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		set: make(map[int]bool),
		el:  []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if ok := this.set[val]; ok {
		return false
	}

	this.set[val] = true
	this.el = append(this.el, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if ok := this.set[val]; ok {
		delete(this.set, val)
		for i := 0; i < len(this.el); i++ {
			if this.el[i] == val {
				this.el = append(this.el[:i], this.el[i+1:]...)
			}
		}
		return true
	}

	return false
}

func (this *RandomizedSet) GetRandom() int {
	randomIdx := rand.Intn(len(this.el))
	return this.el[randomIdx]
}

func main() {
	set := Constructor()
	set.Insert(1)
	set.Insert(2)
	set.Insert(4)
	fmt.Println(set.GetRandom())
}

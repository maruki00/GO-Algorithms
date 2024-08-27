package main

import (
	"fmt"
	"slices"
)

type MinStack struct {
	minValue int
	items    []int
}

func (this *MinStack) isEmpty() bool {
	return len(this.items) == 0
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.minValue = min(this.minValue, val)
	this.items = append(this.items, val)
}

func (this *MinStack) Pop() {
	if this.isEmpty() {
		return
	}
	remv := this.items[len(this.items)-1]
	this.items = this.items[:len(this.items)-1]
	// for _, val := range this.items {
	// 	this.minValue = min(this.minValue, val)
	// }
	if remv != this.minValue {
		this.minValue = slices.Min(this.items)
	}

}

func (this *MinStack) Top() int {
	if this.isEmpty() {
		return 0
	}
	return this.items[len(this.items)-1]
}

func (this *MinStack) GetMin() int {
	if this.isEmpty() {
		return 0
	}
	minVal := this.items[0]
	for _, val := range this.items {
		minVal = min(minVal, val)
	}
	return minVal
}

func main() {

	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	x := obj.Top()
	xx := obj.GetMin()
	obj.Pop()
	xxx := obj.Top()

	fmt.Println(x, xx, xxx)

}

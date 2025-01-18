package main

import "fmt"

type MyStack struct {
	items []int
}

func Constructor() MyStack {
	return MyStack{
		items: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.items = append([]int{x}, this.items...)
}

func (this *MyStack) Pop() int {
	lst := this.items[0]
	this.items = this.items[1:]
	return lst
}

func (this *MyStack) Top() int {
	if this.Empty() {
		return 0
	}
	return this.items[0]
}

func (this *MyStack) Empty() bool {
	return len(this.items) == 0
}

func main() {

	obj := Constructor()
	obj.Push(1234)
	param_2 := obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.Empty()
	fmt.Println(param_4, param_2, param_3)

}

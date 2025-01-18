package main

type MyQueue struct {
	items []int
}

func Constructor() MyQueue {
	return MyQueue{
		items: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.items = append(this.items, x)
}

func (this *MyQueue) Pop() int {
	lst := this.items[0]
	this.items = this.items[1:]
	return lst
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return 0
	}
	return this.items[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.items) == 0
}

func main() {

}

package main

import "fmt"

type Stack struct {
	items  []int
	lenght int
}

func (obj *Stack) push(item int) {
	obj.items = append(obj.items, item)
	obj.lenght++
}

func (obj *Stack) pop() int {
	item := obj.items[obj.lenght]
	obj.items = obj.items[:obj.lenght-2]
	obj.lenght--
	return item
}

func main() {
	stack := &Stack{}
	stack.push(1)
	stack.push(12)
	stack.push(123)
	stack.push(1234)
	p := stack.pop()
	fmt.Println("items: ", stack.items, p)
}

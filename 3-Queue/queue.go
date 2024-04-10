package main

import "fmt"

type Stack struct {
	items  []int
	lenght int
}

func (obj *Stack) append(item int) {
	obj.items = append(obj.items, item)
	obj.lenght++
}

func (obj *Stack) get() int {
	item := obj.items[obj.lenght-1]
	obj.items = obj.items[:obj.lenght-2]
	obj.lenght--
	return item
}

func main() {
	stack := &Stack{}
	stack.append(1)
	stack.append(12)
	stack.append(123)
	stack.append(1234)
	p := stack.get()
	fmt.Println("items: ", stack.items, p)
}

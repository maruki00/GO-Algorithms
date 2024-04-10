package main

import "fmt"

type Node struct {
	next *Node
	item any
}
type Stack struct {
	root *Node
}

func (obj *Stack) add(item any) {
	newNode := &Node{next: nil, item: item}
	if obj.root == nil {
		obj.root = newNode
	} else {
		root := obj.root
		for root.next != nil {
			root = root.next
		}
		root.next = newNode
	}
}

func (obj*Stack)print(){
	root := obj.root
	for root.next !=nil{
		println('Item : ',root.item )
	}
}
func main() {
	stack := &Stack{}
	stack.add(1234)
	stack.add(12344536)
	//stack.root = &Node{next: nil, item: 123}
	stack.print()
	// fmt.Println("Stack: ", stack.root.item)
}

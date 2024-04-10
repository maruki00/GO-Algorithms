package main

import (
	"fmt"
	"sync"
)

type Node struct {
	next *Node
	item any
}
type Stack struct {
	first  *Node
	root   *Node
	last   *Node
	lenght int
	mutex  sync.Mutex
}

func (obj *Stack) add(item any) {
	obj.mutex.Lock()
	defer obj.mutex.Unlock()

	newNode := &Node{next: nil, item: item}
	if obj.root == nil {
		obj.root = newNode
		return
	}
	root := obj.root
	for root.next != nil {
		root = root.next
	}
	root.next = newNode
}

func (obj *Stack) print() {
	obj.mutex.Lock()
	defer obj.mutex.Unlock()
	root := obj.root
	for root != nil {
		fmt.Println("Item : ", root.item)
		root = root.next
	}
}

func (obj *Stack) RemoveAt(index int) {
	root := obj.root
	if index == 0 {
		root = root.next.next
		return
	}
	for index >= 0 {
		root = root.next
		fmt.Println(" ---- 1 Item : ", root.item)
		index--
	}
	for root.next.next != nil {
		root = root.next.next
		fmt.Println(" ---- 2 Item : ", root.item)
	}
}
func main() {
	stack := &Stack{}
	stack.add(1234)
	stack.add(11111)
	stack.add(22222222)
	stack.add(333333333333)
	stack.add(444444444)
	stack.add(6666666666)
	stack.add(77777777)
	stack.add(8888888888)
	stack.add(999999999)
	//stack.root = &Node{next: nil, item: 123}
	stack.RemoveAt(0)
	stack.print()
	// fmt.Println("Stack: ", stack.root.item)
}

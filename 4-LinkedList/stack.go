package main

import (
	"fmt"
	"sync"
)

type NodeSimple struct {
	next *NodeSimple
	item any
}

type SimpleLinkedList struct {
	root   *NodeSimple
	lenght int
	mutex  sync.Mutex
}

func (obj *SimpleLinkedList) add(item any) {
	obj.mutex.Lock()
	defer obj.mutex.Unlock()

	newNode := &NodeSimple{next: nil, item: item}
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

func (obj *SimpleLinkedList) print() {
	obj.mutex.Lock()
	defer obj.mutex.Unlock()
	root := obj.root
	for root != nil {
		fmt.Println("Item : ", root.item)
		root = root.next
	}
}

func (obj *SimpleLinkedList) RemoveAt(index int) {
	root := obj.root
	prev := root
	counter := 0
	for root != nil {
		if root.item != index {
			prev = root
			root = root.next
		} else {
			prev = root.next
			root = prev.next
		}
		counter++
	}
}
func main() {
	stack := &SimpleLinkedList{}
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
	stack.RemoveAt(22222222)
	stack.print()
	// fmt.Println("Stack: ", stack.root.item)
}

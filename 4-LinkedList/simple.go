package main

import (
	"fmt"
	"sync"
)

type NodeSimple struct {
	next  *NodeSimple
	value any
}

type SimpleLinkedList struct {
	root   *NodeSimple
	lenght int
	mutex  sync.Mutex
}

func (obj *SimpleLinkedList) append(value any) {
	obj.mutex.Lock()
	defer obj.mutex.Unlock()

	newNode := &NodeSimple{next: nil, value: value}
	if obj.lenght == 0 {
		obj.root = newNode
		obj.lenght++
		return
	}
	root := obj.root
	for root.next != nil {
		root = root.next
	}
	root.next = newNode
	obj.lenght++
}

func (obj *SimpleLinkedList) preappend(value any) {
	obj.mutex.Lock()
	defer obj.mutex.Unlock()
	if obj.lenght == 0 {
		obj.root = &NodeSimple{value: value, next: nil}
		obj.lenght++
		return
	}

	node := &NodeSimple{value: value, next: nil}
	node.next = obj.root
	obj.root = node
	obj.lenght++
}

func (obj *SimpleLinkedList) print() {
	obj.mutex.Lock()
	defer obj.mutex.Unlock()
	root := obj.root
	for root != nil {
		fmt.Println("value : ", root.value)
		root = root.next
	}
}

func (obj *SimpleLinkedList) remove(value any) {
	root := obj.root
	prev := root
	counter := 0
	for root != nil {
		if root.value != value {
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
	stack.append(1234)
	stack.append(11111)
	stack.append(22222222)
	stack.append(333333333333)
	stack.append(444444444)
	stack.append(6666666666)
	stack.append(77777777)
	stack.append(8888888888)
	stack.preappend(999999999)
	stack.preappend(1000000001)
	stack.preappend("hello world")
	stack.print()
	// fmt.Println("Stack: ", stack.root.value)
}

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
	obj.mutex.Lock()
	defer obj.mutex.Unlock()
	if obj.lenght == 0 {
		return
	}
	if obj.root.value == value {
		obj.root = obj.root.next
		obj.lenght--
		return
	}
	previos := obj.root
	for previos.next.next != nil {
		if previos.next.value != value {
			previos = previos.next
		} else {
			previos.next = previos.next.next
			obj.lenght--
			return
		}
	}
}

func (obj *SimpleLinkedList) clear() {
	obj.mutex.Lock()
	defer obj.mutex.Unlock()
	obj.lenght = 0
	obj.root = nil
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
	fmt.Println("Lenght: ", stack.lenght)
	stack.print()
	stack.remove("hello world")
	stack.remove(999999999)
	stack.remove(999999999897456)
	fmt.Println("Lenght: ", stack.lenght)
	println("---------------------------------------------")
	stack.print()
	stack.clear()
	stack.print()

	// fmt.Println("Stack: ", stack.root.value)
}

package main

import (
	"fmt"
	"sync"
)

type DoubleNode struct {
	previous *DoubleNode
	next     *DoubleNode
	value    any
}

type DoubleLinkedList struct {
	head   *DoubleNode
	lenght int
	sync.Mutex
}

func (obj *DoubleLinkedList) append(value any) {

	if obj.lenght == 0 {
		obj.head = &DoubleNode{value: value, next: nil, previous: nil}
		obj.lenght++
		return
	}

	head := obj.head
	fmt.Println("head: ", head)
	for head.next != nil {
		head.previous = head
		head = head.next
		fmt.Println(value)
	}
	head.next = &DoubleNode{value: value, next: nil, previous: head.previous}
	obj.lenght++
}

func (obj *DoubleLinkedList) print() {
	head := obj.head
	for head.next != nil {
		fmt.Println("Item: ", head.value)
		head.previous = head
		head = head.next
	}
}

func main() {
	doube := &DoubleLinkedList{}
	doube.append(10)

	doube.append(10)
	doube.append(10)
	doube.append(10)
	doube.print()
}

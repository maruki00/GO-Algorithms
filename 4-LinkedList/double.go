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
	for head.next != nil {
		head.previous = head
		head = head.next
	}
	head = &DoubleNode{value: value, next: nil, previous: head.previous}
	obj.lenght++
}

func (obj DoubleLinkedList) print() {
	head := obj.head
	for head.next != nil {
		fmt.Println("Item: ", head.value)
		head.previous = head
		head = head.next
	}
}

func main() {
	Doube := &DoubleLinkedList{}
	Doube.append(10)
	Doube.append(10)
	Doube.append(10)
	Doube.append(10)
	Doube.print()
}

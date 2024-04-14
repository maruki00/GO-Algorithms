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
		fmt.Println(value)
	}
	head.next = &DoubleNode{value: value, next: nil, previous: head.previous}
	obj.lenght++
}

func (obj *DoubleLinkedList) remove(value any) {
	if obj.lenght == 0 {
		return
	}
	if obj.head.value == value {
		obj.head = obj.head.next
		obj.lenght--
		return
	}
	head := obj.head
	for head.next != nil {
		if head.value != value {
			head.previous = head
			head = head.next
		} else {
			head.previous.next = head.next.previous
			obj.lenght--
			return
		}

	}
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
	double := &DoubleLinkedList{}
	double.append(1)

	double.append(2)
	double.append(3)
	double.append(4)
	double.append(5)
	double.append(6)
	double.append(7)

	double.print()

	double.remove(1)
	double.remove(5)
	double.remove(7)

	double.print()
}

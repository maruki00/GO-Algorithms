package main

import "fmt"

/*
1. Reverse a linked list.
**/

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	root   *Node
	tail   *Node
	lenght int
}

func NewLinkedList(val int) *LinkedList {
	root := &Node{Val: val}
	return &LinkedList{
		root:   root,
		tail:   root,
		lenght: 1,
	}
}
func Insert(val int, l *LinkedList) *LinkedList {
	node := &Node{Val: val}
	l.tail.Next = node
	l.tail = node
	l.lenght++
}

func Print(l *LinkedList) {
	tmp := l.root
	for tmp != nil {
		fmt.Println("item : ", tmp.Val)
		tmp = tmp.Next
	}
}

func mai() {
	list := NewLinkedList(1)
	Insert(2, list)
	Insert(3, list)
	Insert(4, list)
	Insert(5, list)
	Print(list)
}

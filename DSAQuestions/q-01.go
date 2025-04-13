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

func Insert(val int, l *LinkedList) {
	node := &Node{Val: val}
	l.tail.Next = node
	l.tail = node
	l.lenght++
}

func Reverse(l *Node) *Node {
	if l.Next == nil {
		return l
	}
	root := l
	curr := l
	for root.Next.Next != nil {
		tmp := root.Next.Next
		curr := tmp
		next := tmp.Next
		next.Next = curr
		curr = next
	}
	return curr
}

func Print(l *Node) {
	tmp := l
	for tmp != nil {
		fmt.Println("item : ", tmp.Val)
		tmp = tmp.Next
	}
}

func main() {
	list := NewLinkedList(1)
	for i := range 1000000 {
		Insert(i, list)
	}

	// Insert(3, list)
	// Insert(4, list)
	// Insert(5, list)
	Print(list.root)
	list.root = Reverse(list.root)
	Print(list.root)
}

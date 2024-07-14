package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func add(list *LinkedList, Val int) *LinkedList {

	tmp := &Node{Val, nil}

	if list == nil {
		return &LinkedList{tmp, tmp}
	}
	list.Tail.Next = tmp
	list.Tail = list.Tail.Next
	return list
}

func reverse(list *LinkedList) *LinkedList {
	var l *Node
	ret := l
	tmp := list.Head
	for {
		l = &Node{tmp.Val, l}
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}
	return &LinkedList{ret, tmp}
}

func print(list *LinkedList) {
	l := list.Head
	for l != nil {
		fmt.Println("item : ", l.Val)
		l = l.Next
	}
}

func main() {

	var l *LinkedList //&LinkedList{1, nil}
	l = add(l, 1)
	l = add(l, 2)
	l = add(l, 4)
	l = add(l, 78)
	l = reverse(l)
	print(l)
}

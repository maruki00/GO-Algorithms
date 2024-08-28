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
	tmp := list.Head
	tail := list.Head
	for tmp != nil {
		l = &Node{tmp.Val, l}
		tmp = tmp.Next
	}
	return &LinkedList{l, tail}
}

func print(list *LinkedList) {
	l := list.Head
	for l != nil {
		fmt.Println("item : ", l.Val, "tail : ", list.Tail.Val)
		l = l.Next
	}
}

func main() {

	var l *LinkedList //&LinkedList{1, nil}
	l = add(l, 1)
	l = add(l, 2)
	l = add(l, 4)
	l = add(l, 78)
	l = add(l, 7676)
	print(l)
	l = add(l, 987)
	fmt.Println("-------------------------", l.Tail.Val)
	l = reverse(l)
	l = add(l, 321)
	print(l)
	fmt.Println("-------------------------", l.Tail.Val)
}

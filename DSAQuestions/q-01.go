package main

import (
	"dsa/ds"
)

/*
1. Reverse a linked list.
**/

func Reverse(l *ds.List[int]) *ds.List[int] {
	var prev *ds.List[int]
	curr := l
	next := curr.Next
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func main() {

	list := ds.NewList(1)

	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	list.Print()
	list = Reverse(list)
	list.Print()
}

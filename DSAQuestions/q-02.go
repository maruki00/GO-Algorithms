package main

import (
	"dsa/ds"
	"fmt"
)

// 2. Find the middle element of a linked list.

func MiddleFind(l *ds.List[int]) int {
	start, end := l, l.Next
	for end != nil {
		end = end.Next
		start = start.Next
	}
	return start.Val
}

func main() {
	list := ds.NewList[int](1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	list.Insert(6)
	fmt.Println("middle is : ", MiddleFind(list))
}

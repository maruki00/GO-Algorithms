package main

import (
	"fmt"
)

/*
	4. Implement a queue using arrays/linked list.
*/

type QueueArray struct {
	items []int
}

func (s *QueueArray) InQueue(val int) {
	s.items = append([]int{val}, s.items...)
}

func (s *QueueArray) DeQueue() int {
	lst := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return lst
}

type QueueLinkedList struct {
	Val  int
	Next *QueueLinkedList
	Tail *QueueLinkedList
}

func NewQueueLinkedList(val int) *QueueLinkedList {
	obj := &QueueLinkedList{
		Val:  val,
		Next: nil,
		Tail: nil,
	}
	obj.Tail = obj
	return obj
}

func (s *QueueLinkedList) InQueue(val int) {
	s.Tail.Next = &QueueLinkedList{Val: val}
	s.Tail = s.Tail.Next
}

func (s *QueueLinkedList) DeQueue() int {
	lst := s.Val
	s = s.Next
	return lst
}

func main() {
	s1 := &QueueArray{}
	s1.InQueue(1)
	s1.InQueue(2)
	fmt.Println(s1.DeQueue())

	s2 := NewQueueLinkedList(2)
	s2.InQueue(2)
	fmt.Println(s2.DeQueue())
}

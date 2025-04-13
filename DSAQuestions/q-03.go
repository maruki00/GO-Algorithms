package main

import "fmt"

/*
	3. Implement a stack using arrays/linked list.
*/

type StackArray struct {
	items []int
}

func (s *StackArray) Push(val int) {
	s.items = append(s.items, val)
}

func (s *StackArray) Pop() int {
	lst := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return lst
}

type StackLinkedList struct {
	Val  int
	Next *StackLinkedList
}

func NewStackLinkedList(val int) *StackLinkedList {
	return &StackLinkedList{
		Val:  val,
		Next: nil,
	}
}

func (s *StackLinkedList) Push(val int) {
	if s.Val == 0 {
		s = &StackLinkedList{Val: val, Next: nil}
		return
	}
	s = &StackLinkedList{Val: val, Next: s}
}

func (s *StackLinkedList) Pop() int {
	lst := s.Val
	s = s.Next
	return lst
}

func main() {
	s1 := &StackArray{}
	s1.Push(1)
	fmt.Println(s1.Pop())

	s2 := NewStackLinkedList(1)
	fmt.Println(s2.Pop())
}

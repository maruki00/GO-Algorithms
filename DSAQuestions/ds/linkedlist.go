package ds

import "fmt"

type List[T comparable] struct {
	Val  T
	Next *List[T]
	Tail *List[T]
}

func NewList[T comparable](val T) *List[T] {
	l := &List[T]{Val: val}
	l.Tail = l
	return l
}

func (l *List[T]) Insert(val T) {
	ll := &List[T]{Val: val}
	l.Tail.Next = ll
	l.Tail = l.Tail.Next
}

func (l *List[T]) Print() {
	println("----------[Begin]---------")
	tmp := l
	for tmp != nil {
		fmt.Println("item : ", tmp.Val)
		tmp = tmp.Next
	}
	println("----------[Done]---------")
}

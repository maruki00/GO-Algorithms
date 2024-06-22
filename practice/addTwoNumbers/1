package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	remain := 0
	_l1 := l1
	for _l1.Next != nil && l2.Next != nil {
		_l1.Val = (_l1.Val + l2.Val + remain) % 10
		remain = int((_l1.Val + l2.Val + remain) / 10)
		_l1 = _l1.Next
		l2 = l2.Next
	}

	return l1

}

func print(l1 *ListNode) {
	for l1 != nil {
		fmt.Println("Item : ", l1.Val)
	}
}

func main() {

	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	l1 = addTwoNumbers(l1, l2)

	print(l1)

}

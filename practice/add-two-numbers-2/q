package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	isNil := func(item *ListNode) int {
		if item != nil {
			return item.Val
		}
		return 0
	}
	remain := 0
	var list *ListNode = new(ListNode)
	l := list
	for l1 != nil || l2 != nil || remain != 0 {
		val := (isNil(l1) + isNil(l2) + remain)
		tmp := &ListNode{(val % 10), nil}
		remain = int(val / 10)
		if list == nil {
			list = tmp
		} else {
			list.Next = tmp
			list = list.Next
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return l.Next
}

func print(l1 *ListNode) {
	for l1 != nil {
		fmt.Println("Item : ", l1.Val)
		l1 = l1.Next
	}
}

func main() {

	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{8, nil}}}
	l := addTwoNumbers(l1, l2)

	print(l)

}

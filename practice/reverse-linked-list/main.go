package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	h := head
	var list *ListNode
	for h != nil {
		next := h.Next
		prev := h
		prev.Next = list
		list = prev
		h = next

	}
	return list
}

func printList(head *ListNode) {
	tmp := head
	for tmp != nil {
		fmt.Println("val : ", tmp.Val)
		tmp = tmp.Next
	}

}

func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	printList(l)
	l = reverseList(l)
	printList(l)
}

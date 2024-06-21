package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteSuplicates(head *ListNode) *ListNode {

	lastItem := head.Val

	for head.Next != nil {

		if head.Next.Val == lastItem {
			head = head.Next.Next
		} else {
			lastItem = head.Val
			head = head.Next
		}
	}

	return head
}

func print(head *ListNode) {
	tmp := head
	for tmp.Next != nil {
		fmt.Println("Item : ", tmp.Val)
		tmp = tmp.Next
	}
}
func main() {
	head := &ListNode{1, nil}
	head.Next = &ListNode{1, nil}
	head.Next.Next = &ListNode{12, nil}
	head.Next.Next.Next = &ListNode{13, nil}
	head.Next.Next.Next.Next = &ListNode{14, nil}
	head.Next.Next.Next.Next.Next = &ListNode{4, nil}
	print(head)
}

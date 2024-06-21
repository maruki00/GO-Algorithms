package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	lastItem := head.Val
	previos := head
	for previos.Next.Next != nil {
		if previos.Next.Val != lastItem {
			previos = previos.Next
			lastItem = previos.Val

		} else {
			previos.Next = previos.Next.Next
			lastItem = previos.Val
		}
	}
	previos.Next = nil
	return head
}

func print(head *ListNode) {
	tmp := head
	for tmp != nil {
		fmt.Println("Item : ", tmp.Val)
		tmp = tmp.Next
	}
}
func main() {
	head := &ListNode{1, nil}
	head.Next = &ListNode{1, nil}
	head.Next.Next = &ListNode{2, nil}
	// head.Next.Next.Next = &ListNode{3, nil}
	// head.Next.Next.Next.Next = &ListNode{4, nil}
	// head.Next.Next.Next.Next.Next = &ListNode{4, nil}
	// head.Next.Next.Next.Next.Next.Next = &ListNode{4, nil}
	print(head)
	head = deleteDuplicates(head)
	println("_________________________________________-")
	print(head)
}

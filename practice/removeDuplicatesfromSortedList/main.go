package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// func deleteDuplicates(head *ListNode) *ListNode {

// 	lastItem := head.Val
// 	previos := head
// 	for previos.Next.Next != nil {
// 		if previos.Next.Val != lastItem {
// 			lastItem = previos.Val
// 			previos = previos.Next

// 		} else {
// 			lastItem = previos.Val
// 			previos.Next = previos.Next.Next
// 		}
// 	}
// 	if previos.Next.Val == lastItem {
// 		previos.Next = nil
// 	}
// 	return head
// }

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	current := head

	for current != nil && current.Next != nil {
		if current.Next.Val == current.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
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
	head.Next.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{4, nil}
	print(head)
	head = deleteDuplicates(head)
	println("_________________________________________-")
	print(head)
}

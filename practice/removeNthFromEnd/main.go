package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var items []int
	counter := 0
	h := head
	for h != nil {
		counter++
		items = append(items, h.Val)
		h = h.Next
	}
	items = append(items[:counter-n], items[counter-n+1:]...)
	if len(items) == 0 {
		return nil
	}

	h2 := head
	for _, item := range items {
		h2.Val = item
		h2 = h2.Next
	}
	h2.Next = nil
	return head
}

func printList(head *ListNode) {
	tmp := head

	for tmp != nil {
		fmt.Println("val : ", tmp.Val)
		tmp = tmp.Next
	}
}

func main() {
	// l := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val: 3,
	// 			Next: &ListNode{
	// 				Val: 4,
	// 				Next: &ListNode{
	// 					Val:  5,
	// 					Next: nil,
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	l := &ListNode{1, nil}

	l = removeNthFromEnd(l, 1)
	printList(l)
}

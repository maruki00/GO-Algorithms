package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	counter := 0
	tmp := head
	nodes := []int{}
	for tmp != nil {
		counter++
		nodes = append(nodes, tmp.Val)
		tmp = tmp.Next
	}

	nodes = append(nodes[:counter-n], nodes[counter-n+1:]...)

	if len(nodes) == 0 {
		return nil
	}

	h := &ListNode{Val: nodes[0], Next: nil}
	head = h
	for i := 1; i < len(nodes); i++ {
		h.Next = &ListNode{Val: nodes[i], Next: nil}
		h = h.Next
	}
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

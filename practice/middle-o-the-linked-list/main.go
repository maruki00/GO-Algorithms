package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {

	low := head
	height := head

	lowCount := 0
	hieghtCount := 0

	for height != nil {
		maxCount := hieghtCount / 2
		if hieghtCount%2 == 0 {
			maxCount--
		}
		for lowCount <= maxCount {
			low = low.Next
			lowCount++
		}

		height = height.Next
		hieghtCount++

	}

	return low
}

func printList(l *ListNode) {
	head := l
	for head != nil {
		fmt.Println("val : ", head.Val)
		head = head.Next
	}
}

func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}

	l = middleNode(l)

	printList(l)

}

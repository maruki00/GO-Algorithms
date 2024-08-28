package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	lowCount := 0
	heightCount := 0

	lowPointer := head
	heightPointer := head

	for heightPointer != nil {
		for lowCount < heightCount-(n+1) {
			lowPointer = lowPointer.Next
			lowCount++
		}

		heightCount++
		heightPointer = heightPointer.Next
	}

	lowPointer = lowPointer.Next.Next

	return head

}

func printList(head *ListNode) {
	tmp := head

	for head != nil {
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

	l = removeNthFromEnd(l, 1)
	printList(l)
}
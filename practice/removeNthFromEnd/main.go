package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	tmp := head
	tmp2 := head
	l := -1
	for tmp != nil {
		l++
		tmp = tmp.Next
	}

	currIndex := 0

	for currIndex+n < l {
		fmt.Println("--- ", currIndex, n, l)
		tmp2 = tmp2.Next
		currIndex++

	}
	printList(tmp2)
	if tmp2.Next.Next == nil {
		tmp2 = nil
		return head
	}
	tmp2.Next = nil

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

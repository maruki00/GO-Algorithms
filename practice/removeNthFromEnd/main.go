package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	lw := head
	hg := head.Next

	lwCounter := 0
	hgCounter := 0

	for hg != nil {
		hg = hg.Next
		hgCounter++
		for lwCounter+n < hgCounter {
			lw = lw.Next
		}
	}
	lw = nil
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

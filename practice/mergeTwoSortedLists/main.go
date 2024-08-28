package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var list *ListNode
	var head *ListNode
	l1 := list1
	l2 := list2

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			next := l1.Next
			if list == nil {
				list = l1
				head = list
				list.Next = nil
			} else {
				list.Next = l1
				list = list.Next
			}

			l1 = next
		} else {
			next := l2
			if list == nil {
				list = l2
				head = list

				list.Next = nil
			} else {
				list.Next = l2
				list = list.Next
			}

			l2 = next
		}
	}
	return head
}

func printList(l *ListNode) {
	tmp := l
	for tmp != nil {
		fmt.Println("value : ", tmp.Val)
		tmp = tmp.Next
	}
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  5,
			Next: nil,
		},
	}

	l2 := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	l := mergeTwoLists(l1, l2)
	printList(l)
}

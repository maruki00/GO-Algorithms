package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListNodee(src *ListNode) {
	tmp := src
	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
}
func lenList(node *ListNode) int {
	counter := 0
	tmp := node
	for tmp != nil {
		counter++
		tmp = tmp.Next
	}
	return counter
}
func reorderList(head *ListNode) {
	head1 := head
	nodesArray := make([]int, lenList(head))
	indx := 0
	for head1 != nil {
		nodesArray[indx] = head1.Val
		indx++
		head1 = head1.Next
	}

	start, end := 0, len(nodesArray)-1
	finalList := &ListNode{Val: nodesArray[0]}
	visited := make([]bool, len(nodesArray))
	tmp := finalList
	visited[start] = true
	for start <= end {
		if !visited[start] {
			tmp.Next = &ListNode{Val: nodesArray[start], Next: nil}
			tmp = tmp.Next
			visited[start] = true

		}
		if end == start {
			break
		}
		if !visited[end] {
			tmp.Next = &ListNode{Val: nodesArray[end], Next: nil}
			tmp = tmp.Next
			visited[end] = true

		}
		end--
		start++
	}
	*head = *finalList
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	reorderList(head)
	printListNodee(head)
}

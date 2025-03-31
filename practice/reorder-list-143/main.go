package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseListNode(src *ListNode) []int {
	tmp := src
	result := []int{}
	for tmp != nil {
		result = append(result, tmp.Val)
		tmp = tmp.Next
	}
	return result
}
func printListNodee(src *ListNode) {
	tmp := src
	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
}
func reorderList(head *ListNode) {
	nodesArray := reverseListNode(head)
	start, end := 0, len(nodesArray)-1
	var finalList *ListNode = nil
	visited := make([]bool, len(nodesArray))
	var tmp *ListNode
	for start <= end {
		if !visited[start] {
			if finalList == nil {
				finalList = &ListNode{Val: nodesArray[start], Next: nil}
				tmp = finalList
			} else {
				tmp.Next = &ListNode{Val: nodesArray[start], Next: nil}
				tmp = tmp.Next
			}
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

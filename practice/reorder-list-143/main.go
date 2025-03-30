package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseListNode(src, dst *ListNode) (*ListNode, int) {
	tmp := src
	counter := 0
	for tmp != nil {
		dst = &ListNode{Val: tmp.Val, Next: dst}
		tmp = tmp.Next
		counter++
	}
	return dst, counter
}
func printListNodee(src *ListNode) {
	tmp := src
	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
}
func reorderList(head *ListNode) {
	var tmpNodes *ListNode
	tmpNodes, counter := reverseListNode(head, tmpNodes)
	visited := make(map[int]bool)
	start, end := 0, counter
	var finalList *ListNode
	for visited[start] && visited[end] {

	}
	printListNodee(head)
	printListNodee(tmpNodes)

}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	reorderList(head)
}

package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func NewNode(val int) *Node {
	return &Node{
		Val:  val,
		Next: nil,
	}
}

func ReversalList(head *Node) {
	if head == nil {
		return
	}
	var prev *Node = nil
	curr, next := head, head.Next
	curr.Next = prev
	for curr != nil {
		prev = curr
		curr = next
		if next != nil {
			next = next.Next
		}
	}
}

func Print(head *Node) {
	tmp := head

	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}

}
func main() {
	head := NewNode(1)
	head.Next = &Node{Val: 7}
	ReversalList(head)
	Print(head)
}

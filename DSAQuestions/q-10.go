package main

import "fmt"

/*
	10. Detect a cycle in a linked list.
*/

type Node struct {
	Val  int
	Next *Node
}

func detectCycle(n *Node) bool {
	if n == nil || n.Next == nil {
		return false
	}
	slow, fast := n, n.Next
	for fast != nil && fast.Next != nil {
		if slow.Val == fast.Val {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false

}

func main() {

	n4 := &Node{Val: 4, Next: nil}
	n3 := &Node{Val: 3, Next: n4}
	n2 := &Node{Val: 2, Next: n3}
	n1 := &Node{Val: 1, Next: n2}
	n4.Next = n2
	fmt.Println("result : ", detectCycle(n1))
}

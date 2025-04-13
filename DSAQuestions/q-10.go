package main

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
	slow, fast := n, n.Next.Next
	for fast != nil && fast.Next != nil {
		if slow.Val == fast.Val {
			return true
		}
		fast = fast.Next.Next
		slow = fast.Next
	}
	return false

}

func main() {
	cll := &Node{
		Val: 1,
		Next: &Node{
			Val: ,
			Next
		},
	}

}

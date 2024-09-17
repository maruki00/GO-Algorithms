package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	// if head == nil {
	// 	return nil
	// }

	// oldToNew := make(map[*Node]*Node)

	// curr := head
	// for curr != nil {
	// 	oldToNew[curr] = &Node{Val: curr.Val}
	// 	curr = curr.Next
	// }

	// curr = head
	// for curr != nil {
	// 	oldToNew[curr].Next = oldToNew[curr.Next]
	// 	oldToNew[curr].Random = oldToNew[curr.Random]
	// 	curr = curr.Next
	// }

	// return oldToNew[head]

	if head == nil {
		return nil
	}

	nodes := make(map[*Node]*Node)

	tmp := head

	for tmp != nil {
		nodes[tmp] = &Node{tmp.Val, nil, nil}
		tmp = tmp.Next
	}

	tmp = head

	for tmp != nil {
		nodes[tmp].Next = nodes[tmp.Next]
		nodes[tmp].Random = nodes[tmp.Random]
		tmp = tmp.Next
	}

	return nodes[head]
}

func printList(node *Node) {
	n := node
	for n != nil {
		fmt.Println("---> ", n.Val, n.Next, n.Random)
		n = n.Next
	}
}

func main() {

	node1 := &Node{3, nil, nil}
	node2 := &Node{3, nil, nil}
	node3 := &Node{3, nil, nil}

	node1.Next = node2
	node1.Random = nil

	node2.Next = node3
	node2.Random = node1

	printList(node1)

	copied := copyRandomList(node1)

	printList(copied)
}

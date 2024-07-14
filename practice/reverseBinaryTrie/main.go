package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func reverse(node *Node) {

	if node == nil {
		return
	}

	node.Left, node.Right = node.Right, node.Left

	reverse(node.Left)
	reverse(node.Right)

}

func main() {
	root := &Node{1, nil, nil}
	reverse(root)
	fmt.Println("Result : ", root)
}

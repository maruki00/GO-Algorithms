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

func print(root *Node) {
	if root == nil {
		return
	}
	fmt.Println("item : ", root.Val)
	print(root.Left)
	print(root.Right)
}

func main() {

	root := &Node{1, &Node{3, &Node{4, nil, nil}, &Node{8, nil, nil}}, &Node{2, &Node{6, nil, nil}, &Node{9, nil, nil}}}
	reverse(root)
	print(root)
	fmt.Println("Result : ", root)
}

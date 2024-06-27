package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var depth func(root *TreeNode, start int) int
	depth = func(node *TreeNode, start int) int {
		if node == nil {
			return 0
		}
		return max(depth(node.Left, start+1), depth(node.Right, start+1)) + 1
	}

	diff := depth(root.Left, 0) - depth(root.Right, 0)
	return diff >= -1 && diff <= 1 && (isBalanced(root.Left) && isBalanced(root.Right))
}

func main() {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println("Result: ", isBalanced(root))
}

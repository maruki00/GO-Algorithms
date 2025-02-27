package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		node.Left, node.Right = node.Right, node.Left
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	return root

}

func main() {
	root := &TreeNode{
		Val:   2,
		Right: &TreeNode{1, nil, nil},
		Left:  &TreeNode{3, nil, nil},
	}
	res := invertTree(root)
	fmt.Println("result : ", res, res.Left, res.Right)
}

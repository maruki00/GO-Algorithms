package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	res := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		fmt.Println("node : ", root.Val)
		if root.Left != nil && (root.Left.Left == nil && root.Left.Right == nil) {
			res += root.Left.Val
		}
		sumOfLeftLeaves(root.Left)
		sumOfLeftLeaves(root.Right)
	}
	return res
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	_ = sumOfLeftLeaves(root)
}

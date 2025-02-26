package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	res := 0
	var dfs func(root *TreeNode, tp string)
	dfs = func(root *TreeNode, tp string) {
		if root == nil {
			return
		}
		if root.Left == nil && tp == "left" {

			res += root.Val
		}
		dfs(root.Left, "left")
		dfs(root.Right, "right")
	}
	dfs(root, "root")
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
	fmt.Println(sumOfLeftLeaves(root))
}

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var dfs func(root *TreeNode, val int) bool
	dfs = func(root *TreeNode, val int) bool {
		if root == nil {
			return true
		}
		val += root.Val
		if root.Left != nil && root.Val <= root.Left.Val && val < root.Left.Val {
			return false
		}
		if root.Right != nil && root.Val >= root.Right.Val && val > root.Right.Val {
			return false
		}
		return dfs(root.Left, root.Val+val) && dfs(root.Right, root.Val+val)
	}
	return dfs(root, 0)
}

func main() {
	// root = [5,4,6,null,null,3,7]

	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println("result : ", isValidBST(root))
}

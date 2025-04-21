package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var dfs func(root *TreeNode, mxVal, mnVal int) bool
	dfs = func(root *TreeNode, mxVal, mnVal int) bool {
		if root == nil {
			return true
		}
		if root.Val < mnVal && root.Val > mxVal {
			return false
		}
		return dfs(root.Left, root.Val, mnVal) && dfs(root.Right, mxVal, root.Val)
	}
	return dfs(root, 0, 0)
}

func main() {
	// root = [5,4,6,null,null,3,7]

	// root := &TreeNode{
	// 	Val: 5,
	// 	Left: &TreeNode{
	// 		Val: 4,
	// 	},
	// 	Right: &TreeNode{
	// 		Val: 6,
	// 		Left: &TreeNode{
	// 			Val: 3,
	// 		},
	// 		Right: &TreeNode{
	// 			Val: 7,
	// 		},
	// 	},
	// }
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}}
	fmt.Println("result : ", isValidBST(root))
}

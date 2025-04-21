package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, mnVal int, mxVal int) bool
	dfs = func(node *TreeNode, mnVal int, mxVal int) bool {
		if node == nil {
			return true
		}
		if !(mnVal < node.Val && node.Val < mxVal) {
			return false
		}
		return dfs(node.Left, mnVal, node.Val) && dfs(node.Right, node.Val, mxVal)
	}
	return dfs(root, math.MinInt, math.MaxInt)
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

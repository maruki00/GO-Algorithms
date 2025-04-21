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

func getMinimumDifference(root *TreeNode) int {
	var res = math.MaxInt
	var dfs func(root *TreeNode, parent int)
	dfs = func(root *TreeNode, parent int) {
		if root == nil {
			return
		}
		dfs(root.Left, root.Val)
		if parent != math.MaxInt {
			res = min(res, int(math.Abs(float64(root.Val-parent))))
		}
		dfs(root.Right, root.Val)

	}
	dfs(root, math.MaxInt)
	return res
}

func main() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}

	fmt.Println("result : ", getMinimumDifference(root))

}

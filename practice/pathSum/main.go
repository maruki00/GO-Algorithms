package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	var dfs func(root *TreeNode, sum int, target int) bool
	dfs = func(root *TreeNode, sum int, target int) bool {
		if root == nil {
			return false
		}
		if sum == target {
			return true
		}
		return dfs(root.Left, sum, target+root.Val) || dfs(root.Right, sum, target+root.Val)
	}
	return dfs(root, sum, 0)
}

func main() {

	nums := &TreeNode{5, &TreeNode{4, &TreeNode{11, &TreeNode{7, nil, nil}, &TreeNode{2, nil, nil}}, nil}, nil}

	fmt.Println("Result : ", hasPathSum(nums))
}

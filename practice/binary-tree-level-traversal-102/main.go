package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if len(result) <= level {
			result = append(result, []int{root.Val})
		} else {
			result[level] = append(result[level], root.Val)
		}
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 0)
	return result
}

func main() {
	root := &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 5},
	}

	fmt.Println("result : ", levelOrder(root))
}

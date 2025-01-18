package main

import "strings"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0)

	var dfs := func(node *TreeTreeNode, val string)
	dfs = func(node *TreeTreeNode, val []string) {
		if node == nil {
			result = append(result, strings.Join(val, "->"))
			return
		}
		dfs(node.Left, []string{val})

	}
	dfs = func ()

	return []string{}
}

func main() {}

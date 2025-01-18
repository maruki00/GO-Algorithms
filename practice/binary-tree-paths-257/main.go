package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0)

	var dfs func(node *TreeNode, items []string)
	dfs = func(node *TreeNode, items []string) {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil {
			result = append(result, strings.Join(append(items, strconv.Itoa(node.Val)), "->"))
			return
		}
		dfs(node.Left, append(items, strconv.Itoa(node.Val)))
		dfs(node.Right, append(items, strconv.Itoa(node.Val)))
	}
	dfs(root, []string{})
	return result
}

func main() {}

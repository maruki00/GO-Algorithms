package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {

	var dfs func(root *TreeNode, dep int) int
	dfs = func(root *TreeNode, dep int) int {
		if root == nil {
			return dep
		}
		return max(dfs(root.Left, dep+1), dfs(root.Right, dep+1))
	}
	return dfs(root, 0)
}

func main() {

}

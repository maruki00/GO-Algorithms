package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 0
	var dfs func(root *TreeNode, dep int)
	dfs = func(root *TreeNode, dep int) {
		if root == nil {
			return
		}

		dfs(root.Left, dep+1)
		if dep > depth {
			depth = dep
		}
		dfs(root.Right, dep+1)
		if dep > depth {
			depth = dep
		}

	}
	dfs(root, 1)
	return depth
}
func main() {

}

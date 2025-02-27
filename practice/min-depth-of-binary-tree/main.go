package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {

	var dfs func(root *TreeNode, depth int) int
	dfs = func(node *TreeNode, depth int) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left, depth+1)
		r := dfs(node.Right, depth+1)

		if min(l, r) == 0 {
			return max(l, r) + 1
		}
		return min(l, r) + 1
	}
	return dfs(root, 0)
}

func main() {

}

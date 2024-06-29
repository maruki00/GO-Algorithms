package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var items []int = []int{}
	var dfs func(root *TreeNode) []int
	dfs = func(root *TreeNode) {
		items = append(items, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}

	return dfs(root)
}

func main() {

}

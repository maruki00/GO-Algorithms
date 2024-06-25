package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {

	if root == nil || (root.Left == nil && root.Right != nil) || (root.Left != nil && root.Right == nil) {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	var dfs func(node1 *TreeNode, node2 *TreeNode) bool
	dfs = func(node1 *TreeNode, node2 *TreeNode) bool {
		if node1 == nil && node2 == nil {
			return false
		}
		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}
		return dfs(node1.Left, node2.Right) && dfs(node1.Right, node2.Left)
	}
	return dfs(root.Left, root.Right)
}

func main() {

}

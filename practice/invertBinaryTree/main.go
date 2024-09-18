package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		node.Left, node.Right = node.Right, node.Left
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	return root

}

func main() {

}

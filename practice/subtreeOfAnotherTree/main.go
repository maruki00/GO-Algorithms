package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidSubTree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}

	if root != nil || subRoot != nil || (root.Val != subRoot.Val) {
		return false
	}
	return isValidSubTree(root.Left, subRoot.Left) && isValidSubTree(root.Right, subRoot.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	if root.Val == subRoot.Val {
		return isSubtree(root, subRoot)
	}

	return isSubtree(root.Left, subRoot.Left) && isSubtree(root.Right, subRoot.Right)
}
func main() {

}

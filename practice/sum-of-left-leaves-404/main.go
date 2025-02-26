package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	fmt.Println("node : ", root.Val)
	if root.Left != nil && (root.Left.Left == nil && root.Left.Right == nil) {
		return root.Left.Val
	}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

func main() {

}

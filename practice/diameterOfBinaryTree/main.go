package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	pathDistance := 0

	var dfs func(node *TreeNode, longest int) int
	dfs = func(node *TreeNode, longest int) int {
		if node == nil {
			return longest
		}
		left := dfs(node.Left, longest+1)
		right := dfs(node.Right, longest+1)
		fmt.Println("result : ", left, right)
		return max((left + right), pathDistance)
	}
	pathDistance = dfs(root, 0)
	return 0
}

func main() {

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	fmt.Println("result : ", diameterOfBinaryTree(root))

}

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	tmp := root

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Val == key {
			if root.Left != nil && root.Left.Val < key {
				root.Left.Val, root.Val = root.Val, root.Left.Val
				dfs(root.Left)
			} else {
				root.Right.Val, root.Val = root.Val, root.Right.Val
				dfs(root.Right)
			}
		}
		if root.Val < key {
			dfs(root.Left)
		} else {
			dfs(root.Right)
		}
	}
	dfs(tmp)
	return root
}

func buildTree() *TreeNode {
	return &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
}
func main() {
	root := buildTree()
	fmt.Println("result : ", root)
}

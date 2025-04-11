package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	tmp := root
	if tmp.Val == key && tmp.Left == tmp.Right && tmp.Right == nil {
		return nil
	}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		} else if root.Val == key {
			if root.Left == nil && root.Right == nil {
				root.Val = 0
				root.Left = nil
				root.Right = nil
				root = nil
				return
			}
			if root.Left != nil && root.Left.Val > key {
				root.Left.Val, root.Val = root.Val, root.Left.Val
				if root.Left.Left == nil && root.Left.Right == nil {
					root.Left = nil
				}
				dfs(root.Left)
			} else {
				root.Right.Val, root.Val = root.Val, root.Right.Val
				if root.Right.Left == nil && root.Right.Right == nil {
					root.Right = nil
				}
				dfs(root.Right)
			}
		} else if root.Val > key {
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

func print(root *TreeNode) {
	tmp := root
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		fmt.Println("val : ", root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(tmp)
}
func main() {
	root := buildTree()
	print(root)
	result := deleteNode(root, 5)

	println("-------------")
	print(result)
	//fmt.Println("result : ", root)
}

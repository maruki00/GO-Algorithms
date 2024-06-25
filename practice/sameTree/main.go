package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	isSameTree(p.Left, q.Right)
	isSameTree(p.Right, q.Left)
	return true
}

func main() {
	root1 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	root2 := &TreeNode{1, &TreeNode{3, nil, nil}, &TreeNode{2, nil, nil}}

	fmt.Println("result : ", isSameTree(root1, root2))
}

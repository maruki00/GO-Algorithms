package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isSameTree_(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	root1 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	root2 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}

	fmt.Println("result : ", isSameTree(root1, root2))
}

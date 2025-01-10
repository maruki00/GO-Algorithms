package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue []*TreeNode

func (q *Queue) Dequeue() *TreeNode {
	old := *q
	lst := old[len(*q)-1]
	old = old[:len(*q)-1]
	*q = old
	return lst
}

func (q *Queue) Inqueue(node *TreeNode) {
	*q = append(*q, node)
}

func (q Queue) Empty() bool {
	return len(q) == 0
}

func kthSmallest(root *TreeNode, k int) int {
	result := math.MaxInt
	i := 0
	nodes := &Queue{root}
	for i < k && !nodes.Empty() {
		nd := nodes.Dequeue()
		if nd == nil {
			continue
		}
		result = min(result, nd.Val)
		nodes.Inqueue(nd.Left)
		nodes.Inqueue(nd.Right)
		i++
	}
	return result
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 5,
		},
	}
	fmt.Println("result : ", kthSmallest(root, 2))
}

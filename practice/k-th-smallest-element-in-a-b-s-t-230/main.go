package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
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
	for i <= k && !nodes.Empty() {
		nd := nodes.Dequeue()
		if nd == nil {
			continue
		}
		result = min(result, nd.Val)
		if nd.Right != nil {
			nodes.Inqueue(nd.Right)
		}
		if nd.Left != nil {
			nodes.Inqueue(nd.Left)
		}
		i++
	}
	return result
}
*/

func kthSmallest(root *TreeNode, k int) int {
	result := -1
	bfs(root, &k, &result)
	return result
}

func bfs(root *TreeNode, counter, result *int) {
	if root == nil {
		return
	}
	fmt.Println("value : ", root.Val)
	bfs(root.Left, counter, result)

	*counter--
	if *counter == 0 {
		*result = root.Val
		return
	}
	bfs(root.Right, counter, result)
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

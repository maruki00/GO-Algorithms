package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue []*TreeNode

func (obj Queue) Push(node *TreeNode) {
	obj = append(obj, node)
}

func (obj Queue) Pop() *TreeNode {
	if len(obj) == 0 {
		return nil
	}
	node := obj[0]
	obj = obj[1:]
	return node
}

func (obj Queue) Empty() bool {
	return len(obj) == 0
}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}

	var bfs func(root *TreeNode)
	bfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		var queue Queue
		var nodes []int
		queue = append(queue, root.Left)
		queue = append(queue, root.Right)
		for !queue.Empty() {
			node := queue.Pop()
			nodes = append(nodes, node.Val)
			bfs(node)
		}
		result = append(result, nodes)
	}

	bfs(root)
	return result

}

func main() {

}

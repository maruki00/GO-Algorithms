package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	visited := make(map[int]bool)
	var dfs func(root1 *TreeNode)
	dfs = func(root *TreeNode) {
		if root.Left == nil && root.Right == nil {
			visited[root.Val] = !visited[root.Val]
		}
	}
	dfs(root1)
	dfs(root2)
	for _, k := range visited {
		if k == true {
			return false
		}
	}
	return true
}

func createBinaryTree(arr []interface{}) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	root := &TreeNode{Val: arr[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for i < len(arr) {
		node := queue[0]
		queue = queue[1:]
		if arr[i] != nil {
			node.Left = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(arr) && arr[i] != nil {
			node.Right = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func printTree(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.Val, " ")
	printTree(node.Left)
	printTree(node.Right)
}

func main() {
	root1Data := []interface{}{3, 5, 1, 6, 2, 9, 8, nil, nil, 7, 4}
	root2Data := []interface{}{3, 5, 1, 6, 7, 4, 2, nil, nil, nil, nil, nil, nil, 9, 8}
	root1 := createBinaryTree(root1Data)
	root2 := createBinaryTree(root2Data)
	fmt.Println("result : ", leafSimilar(root1, root2))
}

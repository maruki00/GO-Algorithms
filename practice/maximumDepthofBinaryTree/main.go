package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	depth := 0
	var godepth func(root *TreeNode)
	godepth = func(root *TreeNode) {
		depth = 10
	}
	godepth(root)
	return depth
}
func main() {

}
